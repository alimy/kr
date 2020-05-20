// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package ignite

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "start a workspace",
		Long:  "start a workspace",
		Run:   startRun,
	}

	// flags inflate
	startCmd.Flags().StringVarP(&confPath, "file", "f", "Ignitefile", "config file path")

	// register startCmd as sub-command
	register(startCmd)
}

func startRun(cmd *cobra.Command, _args []string) {
	w, t := workspaceTier(cmd)
	staging := prepareStaging()
	if err := staging.Start(w, t); err != nil {
		logrus.Fatal()
	}
}
