// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"github.com/alimy/kr/internal/vmware"
	"github.com/spf13/cobra"
)

func init() {
	vmCmd := &cobra.Command{
		Use:   "vm",
		Short: "vmware help toolkit",
		Long:  "vmware help toolkit",
	}
	vmCmd.AddCommand(vmware.NewCmds()...)
	register(vmCmd)
}
