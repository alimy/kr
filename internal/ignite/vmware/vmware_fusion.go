// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package vmware

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/alimy/kr/internal/ignite/process"
	"github.com/alimy/kr/internal/ignite/provision"
)

type vmwareFusion struct {
	execRun     string
	displayMode string
	stateMode   string
}

func (vm *vmwareFusion) init(config provision.ProviderConfig) {
	rootDir := config.RootDir()
	if fs, err := os.Stat(rootDir); err == nil && fs.IsDir() {
		dir, _ := filepath.EvalSymlinks(rootDir)
		vm.execRun = path.Join(dir, "Contents/Public/vmrun")
	}

	display := config.Feature("display")
	if display == "gui" || display == "nogui" {
		vm.displayMode = display
	}

	state := config.Feature("state")
	if state == "hard" || state == "soft" {
		vm.stateMode = state
	}
}

func (vm *vmwareFusion) Start(unit *provision.Unit) error {
	exec := &process.ExecRun{
		Describe: fmt.Sprintf("start tier: %s", unit.Description),
		Cmd:      vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"start",
			unit.Path,
			vm.displayMode,
		},
	}
	if err := exec.Run(); err != nil {
		return err
	}
	return nil
}

func (vm *vmwareFusion) Stop(unit *provision.Unit) error {
	exec := &process.ExecRun{
		Describe: fmt.Sprintf("stop tier: %s", unit.Description),
		Cmd:      vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"stop",
			unit.Path,
			vm.stateMode,
		},
	}
	if err := exec.Run(); err != nil {
		return err
	}
	return nil
}

func (vm *vmwareFusion) Reset(unit *provision.Unit) error {
	exec := &process.ExecRun{
		Describe: fmt.Sprintf("reset tier: %s", unit.Description),
		Cmd:      vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"reset",
			unit.Path,
			vm.stateMode,
		},
	}
	if err := exec.Run(); err != nil {
		return err
	}
	return nil
}

func (vm *vmwareFusion) Suspend(unit *provision.Unit) error {
	exec := &process.ExecRun{
		Describe: fmt.Sprintf("suspend tier: %s", unit.Description),
		Cmd:      vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"suspend",
			unit.Path,
			vm.stateMode,
		},
	}
	if err := exec.Run(); err != nil {
		return err
	}
	return nil
}

func newVMwareFusion() *vmwareFusion {
	return &vmwareFusion{
		execRun:     "/Applications/VMware Fusion.app/Contents/Public/vmrun",
		displayMode: "nogui",
		stateMode:   "hard",
	}
}
