module github.com/maldan/gam-app-fileman

go 1.16

replace github.com/maldan/go-rapi => ../../go_lib/rapi

replace github.com/maldan/go-cmhp => ../../go_lib/cmhp

require (
	github.com/Knetic/govaluate v3.0.0+incompatible
	github.com/h2non/filetype v1.1.3
	github.com/maldan/go-cmhp v0.0.20
	github.com/maldan/go-rapi v0.0.6
)
