// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package provision

import (
	"sync"

	"github.com/alimy/kr/internal/ignite/config"
	"github.com/sirupsen/logrus"
)

var (
	providers       = make(map[string]providerInstance)
	providerConfigs = make(map[string]ProviderConfig)
)

type providerInstance struct {
	factory  ProviderFactory
	instance Provider
	once     sync.Once
}

type providerSpec struct {
	name        string
	rootdir     string
	description string
	features    map[string]string
}

func (p *providerInstance) Get() Provider {
	p.once.Do(func() {
		c := providerConfigs[p.factory.Name()]
		p.instance = p.factory.NewProvider(c)
	})
	return p.instance
}

func (p *providerSpec) Name() string {
	return p.name
}

func (p *providerSpec) Description() string {
	return p.description
}

func (p *providerSpec) RootDir() string {
	return p.rootdir
}

func (p *providerSpec) Feature(name string) string {
	return p.features[name]
}

func FindProviderByName(name string) Provider {
	if p, exist := providers[name]; exist {
		return p.Get()
	}
	return nil
}

func Register(factory ProviderFactory) {
	if factory != nil {
		providers[factory.Name()] = providerInstance{
			factory: factory,
		}
	}
}

func InitProviderWith(config *config.IgniteConfig) {
	for _, spec := range config.Providers {
		ps := &providerSpec{
			name:        spec.Name,
			description: spec.Description,
			rootdir:     spec.RootDir,
		}
		providerConfigs[spec.Name] = ps
		attrs, diag := spec.Features.Body.JustAttributes()
		if diag != nil && diag.HasErrors() {
			logrus.Warn(diag)
			continue
		}
		ps.features = make(map[string]string, len(attrs))
		for name, attr := range attrs {
			v, diag := attr.Expr.Value(nil)
			if diag != nil && diag.HasErrors() {
				logrus.Warn(diag)
				continue
			}
			ps.features[name] = v.AsString()
		}
	}
}
