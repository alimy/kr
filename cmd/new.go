// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/alimy/kr/internal/generator"
	"github.com/spf13/cobra"
)

var (
	dstPath string
	pkgName string
	style   []string
)

func init() {
	newCmd := &cobra.Command{
		Use:   "new",
		Short: "create template project",
		Long:  "create template project",
		Run:   newRun,
	}

	// parse flags for agentCmd
	newCmd.Flags().StringVarP(&dstPath, "dst", "d", ".", "genereted destination target directory")
	newCmd.Flags().StringVarP(&pkgName, "pkg", "p", "github.com/alimy/examples", "project's package name")
	newCmd.Flags().StringSliceVarP(&style, "style", "s", []string{"grpc"}, "generated engine type style:[{grpc[,zim]}|zim|{[mir,]gin|chi|echo|iris|macaron|mux|httprouter}|{tars[[,mir],gin|,simple]}], default is grpc")

	// register agentCmd as sub-command
	register(newCmd)
}

// newRun run new command
func newRun(_cmd *cobra.Command, _args []string) {
	path, err := filepath.EvalSymlinks(dstPath)
	if err != nil {
		if os.IsNotExist(err) {
			if !filepath.IsAbs(dstPath) {
				cwd, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				path = filepath.Join(cwd, dstPath)
			} else {
				path = dstPath
			}
		} else {
			log.Fatal(err)
		}
	}
	if err = generator.Generate(path, style, pkgName); err != nil {
		log.Fatal(err)
	}
}
