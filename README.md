## Kr
develop help toolkit

### Usage
```bash
% go get github.com/alimy/kr@latest
% mkdir -p $GOPATH/src/github.com/alimy
% cd $GOPATH/src/github.com/alimy
% kr new -d examples -p github.com/alimy/examples
% tree -L 2 examples
examples
├── Makefile
├── README.md
├── assets
│   ├── README.md
│   ├── certs
│   └── config
├── cmd
│   ├── README.md
│   ├── core
│   ├── zimctl
│   ├── zimctl.go
│   ├── zimlet
│   └── zimlet.go
├── core
│   ├── core.go
│   ├── models
│   ├── proto
│   └── servant
├── docs
│   └── README.md
├── go.mod
├── hack
│   ├── README.md
│   ├── goprotoc.sh
│   └── systemd
├── internal
│   ├── assets
│   ├── config
│   ├── insecure
│   ├── logus
│   └── rpcx
├── servants
│   ├── agent
│   ├── business
│   ├── dataware
│   ├── servants.go
│   └── share
└── version
    └── version.go

% cd examples
% make build
% ls
Makefile  assets    core      go.mod    hack      servants  zimctl
README.md cmd       docs      go.sum    internal  version   zimlet
```
