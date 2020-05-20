// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package ignite

import (
	"github.com/spf13/cobra"
)

func init() {
	resumeCmd := &cobra.Command{
		Use:   "resume",
		Short: "resume a workspace",
		Long:  "resume a workspace",
		Run:   startRun,
	}

	// flags inflate
	resumeCmd.Flags().StringVarP(&confPath, "file", "f", "Ignitefile", "config file path")

	// register resumeCmd as sub-command
	register(resumeCmd)
}
