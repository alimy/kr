// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"github.com/alimy/kr/cmd/ignite"
)

func init() {
	register(ignite.Cmd())
}
