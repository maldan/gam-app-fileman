module github.com/maldan/gam-app-fileman

go 1.16

replace github.com/maldan/go-restserver => ../../../go_lib/restserver

replace github.com/maldan/go-cmhp => ../../../go_lib/cmhp

require (
	github.com/maldan/go-cmhp v0.0.3
	github.com/maldan/go-restserver v1.2.1
	github.com/zserge/lorca v0.1.10
)
