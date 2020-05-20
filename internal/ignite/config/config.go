// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

var (
	igniteConfig *IgniteConfig
)

type Staging struct {
	Fallback        bool   `hcl:"fallback,optional"`
	DefaultProvider string `hcl:"default-provider,optional"`
}

type Features struct {
	Body hcl.Body `hcl:",remain"`
}

type Provider struct {
	Name        string    `hcl:"name,label"`
	Description string    `hcl:"description,optional"`
	RootDir     string    `hcl:"rootdir,attr"`
	Features    *Features `hcl:"features,block"`
}

type Unit struct {
	Name        string   `hcl:"name,label"`
	Description string   `hcl:"description,optional"`
	Path        string   `hcl:"path,attr"`
	Provider    string   `hcl:"provider,optional"`
	Hosts       []string `hcl:"hosts,optional"`
}

type Tier struct {
	Name         string   `hcl:"name,label"`
	Dependencies []string `hcl:"dependencies,optional"`
}

type Workspace struct {
	Name        string   `hcl:"name,label"`
	Description string   `hcl:"description,attr"`
	TierList    []string `hcl:"tiers,optional"`
	Tiers       []*Tier  `hcl:"tier,block"`
}

type IgniteConfig struct {
	Version    string       `hcl:"version,attr"`
	Staging    *Staging     `hcl:"staging,block"`
	Providers  []*Provider  `hcl:"provider,block"`
	Units      []*Unit      `hcl:"unit,block"`
	Workspaces []*Workspace `hcl:"workspace,block"`
}

func MyConfig() *IgniteConfig {
	return igniteConfig
}

func ParseFrom(filename string) (*IgniteConfig, error) {
	igniteConfig = &IgniteConfig{}
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, hcl.Diagnostics{
				{
					Severity: hcl.DiagError,
					Summary:  "Configuration file not found",
					Detail:   fmt.Sprintf("The configuration file %s does not exist.", filename),
				},
			}
		}
		return nil, hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  "Failed to read configuration",
				Detail:   fmt.Sprintf("Can't read %s: %s.", filename, err),
			},
		}
	}
	file, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{Line: 1, Column: 1})
	if diags != nil && diags.HasErrors() {
		return nil, diags
	}
	diags = gohcl.DecodeBody(file.Body, nil, igniteConfig)
	if diags != nil && diags.HasErrors() {
		return igniteConfig, diags
	}
	return igniteConfig, nil
}
