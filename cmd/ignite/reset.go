// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package ignite

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	resetCmd := &cobra.Command{
		Use:   "reset",
		Short: "reset a workspace",
		Long:  "reset a workspace",
		Run:   resetRun,
	}

	// flags inflate
	resetCmd.Flags().StringVarP(&confPath, "file", "f", "Ignitefile", "config file path")

	// register resetCmd as sub-command
	register(resetCmd)
}

func resetRun(cmd *cobra.Command, _args []string) {
	w, t := workspaceTier(cmd)
	staging := prepareStaging()
	if err := staging.Reset(w, t); err != nil {
		logrus.Fatal()
	}
}
