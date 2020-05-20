// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package vmware

import (
	"github.com/alimy/kr/internal/ignite/provision"
)

func init() {
	provision.Register(providerFactory{
		name: "vmware-fusion",
	})
}

type providerFactory struct {
	name string
}

func (p providerFactory) Name() string {
	return p.name
}

func (p providerFactory) NewProvider(config provision.ProviderConfig) provision.Provider {
	vm := newVMwareFusion()
	if config != nil && config.Name() == "vmware-fusion" {
		vm.init(config)
	}
	return vm
}
