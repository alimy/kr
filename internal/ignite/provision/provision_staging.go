// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package provision

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alimy/kr/internal/ignite/config"
	"github.com/sirupsen/logrus"
)

var (
	errNotExistWorkspace = errors.New("not exist workspace")
	errNotExistTier      = errors.New("not exist tier")

	actionStart   = "start"
	actionStop    = "stop"
	actionReset   = "reset"
	actionSuspend = "suspend"
)

type StateMsg struct {
	State TierState
	Tier  *Tier
	Error error
}

type Staging struct {
	Fallback        bool
	DefaultProvider string
	Workspaces      map[string]*Workspace
}

type handler struct {
	Name string
	Func func(*Unit) error
}

func handleFun(name string, h func(*Unit) error) *handler {
	return &handler{
		Name: name,
		Func: h,
	}
}

func (s *Staging) Start(workspace string, tier string) error {
	return s.run(actionStart, workspace, tier)
}

func (s *Staging) Stop(workspace string, tier string) error {
	return s.run(actionStop, workspace, tier)
}

func (s *Staging) Reset(workspace string, tier string) error {
	return s.run(actionReset, workspace, tier)
}

func (s *Staging) Suspend(workspace string, tier string) error {
	return s.run(actionSuspend, workspace, tier)
}

func (s *Staging) run(action string, workspace string, tier string) error {
	var h *handler
	switch action {
	case actionStart:
		h = handleFun("start", func(unit *Unit) error {
			return unit.Start()
		})
	case actionStop:
		h = handleFun("stop", func(unit *Unit) error {
			return unit.Stop()
		})
	case actionReset:
		h = handleFun("reset", func(unit *Unit) error {
			return unit.Reset()
		})
	case actionSuspend:
		h = handleFun("suspend", func(unit *Unit) error {
			return unit.Suspend()
		})
	}
	if workspace == "" {
		return s.handleAllWorkspace(h)
	}
	return s.handleWorkspace(workspace, tier, h)
}

func (s *Staging) tiersBy(workspace string, tier string) (map[string]*Tier, error) {
	ws, exist := s.Workspaces[workspace]
	if !exist {
		return nil, errNotExistWorkspace
	}
	if tier != "" {
		tier, exist := ws.Tiers[tier]
		if !exist {
			return nil, errNotExistTier
		}
		return map[string]*Tier{tier.Name: tier}, nil
	}
	return ws.Tiers, nil
}

func (s *Staging) handleTiers(tiers map[string]*Tier, h *handler) (err error) {
	switch h.Name {
	case actionStart:
		err = s.handleTiersAsc(tiers, h)
	case actionStop, actionReset, actionSuspend:
		err = s.handleTiersDesc(tiers, h)
	}
	return
}

// handleTiersAsc 正序处理tiers
func (s *Staging) handleTiersAsc(tiers map[string]*Tier, h *handler) error {
	stateChan := make(chan *StateMsg, 8)
	remainTiersCount := len(tiers)
	for _, tier := range tiers {
		if len(tier.Parents) == 0 {
			go handleTier(stateChan, tier, h.Func)
		}
	}
	for sc := range stateChan {
		if sc.State == TierStateFailed {
			return fmt.Errorf("%s tier of %s failed: %w", h.Name, sc.Tier.Name, sc.Error)
		}
		if remainTiersCount--; remainTiersCount == 0 { // handler finish
			break
		}
		children := sc.Tier.Children
		for name := range children {
			tier := tiers[name]
			tier.SetParentState(sc.Tier.Name, TierStateDone)
			if tier.IsParentsDone() {
				go handleTier(stateChan, tier, h.Func)
			}
		}
	}
	return nil
}

// handleTiersAsc 反序处理tiers
func (s *Staging) handleTiersDesc(tiers map[string]*Tier, h *handler) error {
	stateChan := make(chan *StateMsg, 8)
	remainTiersCount := len(tiers)
	for _, tier := range tiers {
		if len(tier.Children) == 0 {
			go handleTier(stateChan, tier, h.Func)
		}
	}
	for sc := range stateChan {
		if sc.State == TierStateFailed {
			return fmt.Errorf("%s tier of %s failed: %w", h.Name, sc.Tier.Name, sc.Error)
		}
		if remainTiersCount--; remainTiersCount == 0 { // handler finish
			break
		}
		parents := sc.Tier.Parents
		for name := range parents {
			tier := tiers[name]
			tier.SetChildState(sc.Tier.Name, TierStateDone)
			if tier.IsChildrenDone() {
				go handleTier(stateChan, tier, h.Func)
			}
		}
	}
	return nil
}

func (s *Staging) handleWorkspace(workspace string, tier string, h *handler) error {
	tiers, err := s.tiersBy(workspace, tier)
	if err != nil {
		return err
	}
	if tier != "" {
		logrus.Infof("%s workspace: %s tier: %s\n", h.Name, workspace, tier)
	} else {
		logrus.Infof("%s workspace: %s\n", h.Name, workspace)
	}
	return s.handleTiers(tiers, h)
}

func (s *Staging) handleAllWorkspace(h *handler) error {
	for name := range s.Workspaces {
		logrus.Infof("%s workspace: %s\n", h.Name)
		tiers, _ := s.tiersBy(name, "")
		if err := s.handleTiers(tiers, h); err != nil {
			return err
		}
	}
	return nil
}

func handleTier(stateChan chan<- *StateMsg, tier *Tier, action func(*Unit) error) {
	if err := action(tier.Unit); err != nil {
		stateChan <- &StateMsg{
			State: TierStateFailed,
			Tier:  tier,
			Error: err,
		}
	} else {
		stateChan <- &StateMsg{
			State: TierStateDone,
			Tier:  tier,
		}
	}
}

func DefaultStaging() *Staging {
	return &Staging{
		DefaultProvider: "vmware-fusion",
		Workspaces:      make(map[string]*Workspace),
	}
}

func StagingFrom(config *config.IgniteConfig) *Staging {
	staging := DefaultStaging()
	if config == nil {
		return staging
	}
	staging.Fallback = config.Staging.Fallback
	if config.Staging.DefaultProvider != "" {
		staging.DefaultProvider = config.Staging.DefaultProvider
	}
	units := unitsFrom(config, staging.DefaultProvider)
	for _, ws := range config.Workspaces {
		workspace := &Workspace{
			Name:        ws.Name,
			Description: ws.Description,
		}
		tiers := make(map[string]*Tier, len(ws.TierList)+len(ws.Tiers))
		for _, name := range ws.TierList {
			if _, exist := tiers[name]; exist {
				continue
			}
			unit, exist := units[name]
			if !exist {
				logrus.Warnf("not exist unit named %s\n", name)
				continue
			}
			tiers[name] = &Tier{
				Unit: unit,
			}
		}
		for _, ts := range ws.Tiers {
			unit, exist := units[ts.Name]
			if !exist {
				logrus.Warnf("not exist unit named %s\n", ts.Name)
				continue
			}
			tmpTiers := make(map[string]*Tier, len(ts.Dependencies))
			for _, dn := range ts.Dependencies {
				unit, exist := units[dn]
				if !exist {
					logrus.Warnf("not exist unit named %s\n", dn)
					continue
				}
				tier, exist := tiers[dn]
				if !exist {
					tmpTiers[dn] = &Tier{
						Unit: unit,
					}
				}
				if tier.Children == nil {
					tier.Children = make(map[string]TierState)
				}
				tier.Children[ts.Name] = TierStateUnknown
			}
			for _, tmpTier := range tmpTiers {
				tiers[tmpTier.Name] = tmpTier
			}
			tier, exist := tiers[ts.Name]
			if !exist {
				tier = &Tier{
					Unit: unit,
				}
				tiers[ts.Name] = tier
			}
			if tier.Parents == nil {
				tier.Parents = make(map[string]TierState, len(ts.Dependencies))
			}
			for _, name := range ts.Dependencies {
				tier.Parents[name] = TierStateUnknown
			}
		}
		workspace.Tiers = tiers
		staging.Workspaces[ws.Name] = workspace
	}
	return staging
}

func unitsFrom(config *config.IgniteConfig, defaultProvider string) map[string]*Unit {
	units := make(map[string]*Unit, len(config.Units))
	for _, spec := range config.Units {
		unit := &Unit{
			Name:        spec.Name,
			Description: spec.Description,
			Path:        fixedPath(spec.Path),
		}
		if spec.Provider != "" {
			unit.Provider = spec.Provider
		} else {
			unit.Provider = defaultProvider
		}
		unit.Hosts = make([]Host, 0, len(spec.Hosts))
		for _, name := range spec.Hosts {
			unit.Hosts = append(unit.Hosts, Host{
				Name: name,
			})
		}
		units[spec.Name] = unit
	}
	return units
}

func fixedPath(path string) string {
	if path[0] == '~' {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		return filepath.Join(homedir, path[1:])
	}
	if absPath, err := filepath.Abs(path); err == nil {
		return absPath
	}
	return path
}
