// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package vmware

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

type cmdInfo struct {
	describe string
	cmd      string
	argv     []string
}

func runCmd(ci *cmdInfo) error {
	logrus.Info(ci.describe)
	process, err := os.StartProcess(ci.cmd, ci.argv, nil)
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
