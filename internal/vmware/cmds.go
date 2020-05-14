// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package vmware

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	optConfPath string
	optCmd      string
	optGui      bool
	optSoft     bool
	optWs       bool

	errInvalideClusterName = errors.New("invalide cluster name")
)

func NewCmds() []*cobra.Command {
	cmds := []*cobra.Command{
		{
			Use:   "start",
			Short: "start a vmware node",
			Long:  "start a vmware node",
			Run:   startRun,
		},
		{
			Use:   "stop",
			Short: "stop a vmware node",
			Long:  "stop a vmware node",
			Run:   stopRun,
		},
		{
			Use:   "reset",
			Short: "reset a vmware node",
			Long:  "reset a vmware node",
			Run:   resetRun,
		},
		{
			Use:   "suspend",
			Short: "suspend a vmware node",
			Long:  "suspend a vmware node",
			Run:   suspendRun,
		},
		{
			Use:   "resume",
			Short: "resume a vmware node",
			Long:  "resume a vmware node",
			Run:   startRun,
		},
	}
	for i, cmd := range cmds {
		if i == 0 { // start cmd
			cmd.Flags().BoolVar(&optGui, "gui", false, "whether in gui mode")
		} else {
			cmd.Flags().BoolVar(&optSoft, "soft", false, "whether in soft mode")
		}
		cmd.Flags().StringVar(&optCmd, "cmd", "vmrun", "cmd full path")
		cmd.Flags().BoolVar(&optWs, "ws", false, "whether in ws")
		cmd.Flags().StringVarP(&optConfPath, "config", "c", "cluster.yml", "config file path")
	}
	return cmds
}

func startRun(cmd *cobra.Command, _args []string) {
	clusters := clusterInfos(cmd)
	argGui, _, argFusion := argsFixed()
	ci := &cmdInfo{
		cmd: optCmd,
		argv: []string{
			optCmd,
			"-T",
			argFusion,
			"start",
			"",
			argGui,
		},
	}
	for _, cluster := range clusters {
		logrus.Infof("start cluster %s\n", cluster.name)
		for _, node := range cluster.nodes {
			ci.describe = fmt.Sprintf("start node %s", node.name)
			ci.argv[4] = node.path
			if err := runCmd(ci); err != nil {
				logrus.Fatal(err)
			}
		}
	}
}

func stopRun(cmd *cobra.Command, _args []string) {
	clusters := clusterInfos(cmd)
	argGui, _, argFusion := argsFixed()
	ci := &cmdInfo{
		cmd: optCmd,
		argv: []string{
			optCmd,
			"-T",
			argFusion,
			"stop",
			"",
			argGui,
		},
	}
	for _, cluster := range clusters {
		logrus.Infof("stop cluster %s\n", cluster.name)
		for _, node := range cluster.nodes {
			ci.describe = fmt.Sprintf("stop node %s", node.name)
			ci.argv[4] = node.path
			if err := runCmd(ci); err != nil {
				logrus.Fatal(err)
			}
		}
	}
}

func resetRun(cmd *cobra.Command, _args []string) {
	clusters := clusterInfos(cmd)
	_, argHard, argFusion := argsFixed()
	ci := &cmdInfo{
		cmd: optCmd,
		argv: []string{
			optCmd,
			"-T",
			argFusion,
			"reset",
			"",
			argHard,
		},
	}
	for _, cluster := range clusters {
		logrus.Infof("reset cluster %s\n", cluster.name)
		for _, node := range cluster.nodes {
			ci.describe = fmt.Sprintf("reset node %s", node.name)
			ci.argv[4] = node.path
			if err := runCmd(ci); err != nil {
				logrus.Fatal(err)
			}
		}
	}
}

func suspendRun(cmd *cobra.Command, _args []string) {
	clusters := clusterInfos(cmd)
	_, argHard, argFusion := argsFixed()
	ci := &cmdInfo{
		cmd: optCmd,
		argv: []string{
			optCmd,
			"-T",
			argFusion,
			"suspend",
			"",
			argHard,
		},
	}
	for _, cluster := range clusters {
		logrus.Infof("suspend cluster %s\n", cluster.name)
		for _, node := range cluster.nodes {
			ci.describe = fmt.Sprintf("suspend node %s", node.name)
			ci.argv[4] = node.path
			if err := runCmd(ci); err != nil {
				logrus.Fatal(err)
			}
		}
	}
}

func argsFixed() (string, string, string) {
	argGui := "nogui"
	argHard := "hard"
	argFusion := "fusion"
	if optGui {
		argGui = "gui"
	}
	if optSoft {
		argHard = "sort"
	}
	if optWs {
		argFusion = "ws"
	}
	return argGui, argHard, argFusion
}

func clusterInfos(cmd *cobra.Command) []*cluster {
	clusters, err := parseFrom(optConfPath)
	if err != nil {
		logrus.Fatal(err)
	}

	// set target cluster name
	if cmd.Flags().NArg() > 0 {
		name := cmd.Flags().Arg(0)
		c, exist := clusters[name]
		if !exist {
			logrus.Fatal(errInvalideClusterName)
		}
		return []*cluster{
			c,
		}
	}

	// process all cluster
	cs := make([]*cluster, 0, len(clusters))
	for _, c := range clusters {
		cs = append(cs, c)
	}
	return cs
}
