// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package ignite

import (
	"github.com/alimy/kr/internal/ignite/config"
	"github.com/alimy/kr/internal/ignite/provision"

	_ "github.com/alimy/kr/internal/ignite/vmware"
)

func Initialize(config *config.IgniteConfig) {
	provision.InitProviderWith(config)
}
