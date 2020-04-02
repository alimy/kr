// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

import (
	"sort"
	"strings"
)

//go:generate go-bindata -nomemcopy -pkg=${GOPACKAGE} -ignore=README.md -prefix=templates -debug=false -o=templates_gen.go templates/...

// tmplCtx template context for generate project
type tmplCtx struct {
	PkgName string
}

// tmplInfo template data info
type tmplInfo struct {
	name   string
	isTmpl bool
	isExec bool
}

type tmplInfos = map[string]tmplInfo

var (
	// tmplFiles map of generated file's assets info
	tmplFiles = make(map[string]tmplInfos, 2)

	// styles map of style slice to style
	styles = make(map[string]string)
)

func init() {
	for _, s := range []struct {
		styles []string
		target string
	}{
		{[]string{"zim"}, "grpc"},
		{[]string{"grpc"}, "grpc"},
		{[]string{"grpc", "zim"}, "grpc"},
		{[]string{"gin"}, "gin"},
		{[]string{"mir", "gin"}, "gin"},
		{[]string{"chi"}, "chi"},
		{[]string{"mir", "chi"}, "chi"},
		{[]string{"echo"}, "echo"},
		{[]string{"mir", "echo"}, "echo"},
		{[]string{"macaron"}, "macaron"},
		{[]string{"mir", "macaron"}, "macaron"},
		{[]string{"mux"}, "mux"},
		{[]string{"mir", "mux"}, "mux"},
		{[]string{"iris"}, "iris"},
		{[]string{"mir", "iris"}, "iris"},
		{[]string{"fiber"}, "fiber"},
		{[]string{"mir", "fiber"}, "fiber"},
		{[]string{"httprouter"}, "httprouter"},
		{[]string{"mir", "httprouter"}, "httprouter"},
	} {
		sort.Strings(s.styles)
		styles[strings.Join(s.styles, ":")] = s.target
	}
}

func tmplInfosFrom(s []string) (tmplInfos, bool) {
	sort.Strings(s)
	style := styles[strings.Join(s, ":")]
	res, exist := tmplFiles[style]
	return res, exist
}
