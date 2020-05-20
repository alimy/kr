// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package process

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

type ExecRun struct {
	Describe string
	Cmd      string
	Argv     []string
}

func (r *ExecRun) Run() error {
	logrus.Info(r.Describe)
	attr := &os.ProcAttr{}
	if homedir, err := os.UserHomeDir(); err == nil {
		attr.Dir = homedir
	}
	process, err := os.StartProcess(r.Cmd, r.Argv, attr)
	if err != nil {
		return err
	}
	ps, err := process.Wait()
	if err != nil {
		return err
	}
	if !ps.Success() {
		return errors.New(ps.String())
	}
	return nil
}
