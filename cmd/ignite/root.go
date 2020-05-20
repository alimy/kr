// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package ignite

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ignite",
		Short: "vm help toolkit",
		Long:  `vm help toolkit`,
	}
)

// register add sub-command
func register(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func Cmd() *cobra.Command {
	return rootCmd
}
