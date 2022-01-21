module github.com/maldan/gam-app-fileman

go 1.18

replace github.com/maldan/go-rapi => ../../go_lib/rapi

replace github.com/maldan/go-cmhp => ../../go_lib/cmhp

require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/h2non/filetype v1.1.3
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/maldan/go-cmhp v0.0.20
	github.com/maldan/go-rapi v0.0.6
	github.com/maldan/go-restserver v1.2.10
)

require (
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	golang.org/x/sys v0.0.0-20210423082822-04245dca01da // indirect
)
