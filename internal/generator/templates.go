// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

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

// tmplFiles map of generated file's assets info
var tmplFiles = make(map[string]tmplInfos, 2)
