// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"github.com/alimy/kr/internal/cmd"
)

func main() {
	cmd.Setup(
		"kr",                   // command name
		"develop help toolkit", // command short describe
		"develop help toolkit", // command long describe
	)
	// execute start application
	cmd.Execute()
}
