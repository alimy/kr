// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

func init() {
	tmplFiles["gin"] = tmplInfos{
		"Makefile":               {"mir/makefile.tmpl", false, false},
		"README.md":              {"mir/readme.tmpl", false, false},
		"go.mod":                 {"mir/gin_go_mod.tmpl", true, false},
		"main.go":                {"mir/gin_main.tmpl", false, false},
		"mirc/main.go":           {"mir/gin_mirc_main.tmpl", true, false},
		"mirc/routes/site.go":    {"mir/gin_mirc_routes_site.tmpl", false, false},
		"mirc/routes/v1/site.go": {"mir/gin_mirc_routes_site_v1.tmpl", false, false},
		"mirc/routes/v2/site.go": {"mir/gin_mirc_routes_site_v2.tmpl", false, false},
	}
}
