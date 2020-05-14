// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package vmware

import "errors"

var (
	errConfNotExist = errors.New("config file not exist")
)

type node struct {
	name string
	path string
}

type cluster struct {
	name  string
	nodes []node
}

type clusterMap map[string]*cluster

type config struct {
	Version string
}

func parseFrom(path string) (clusterMap, error) {
	// TODO
	return nil, errConfNotExist
}
