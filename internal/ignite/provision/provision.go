// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package provision

const (
	TierStateUnknown TierState = iota
	TierStateDone
	TierStateFailed
)

type TierState = int8

type Workspace struct {
	Description string
	Name        string
	Tiers       map[string]*Tier
}

type Tier struct {
	*Unit
	Parents  map[string]TierState
	Children map[string]TierState
}

func (t *Tier) SetParentState(name string, state TierState) {
	t.Parents[name] = state
}

func (t *Tier) SetChildState(name string, state TierState) {
	t.Children[name] = state
}

func (t *Tier) IsParentsDone() bool {
	for _, p := range t.Parents {
		if p != TierStateDone {
			return false
		}
	}
	return true
}

func (t *Tier) IsChildrenDone() bool {
	for _, c := range t.Children {
		if c != TierStateDone {
			return false
		}
	}
	return true
}

type Unit struct {
	Description string
	Name        string
	Provider    string
	Path        string
	Hosts       []Host
}

type Host struct {
	Name string
}

type ProviderConfig interface {
	Name() string
	Description() string
	RootDir() string
	Feature(string) string
}

type Provider interface {
	Start(*Unit) error
	Stop(*Unit) error
	Reset(*Unit) error
	Suspend(*Unit) error
}

type ProviderFactory interface {
	Name() string
	NewProvider(ProviderConfig) Provider
}
