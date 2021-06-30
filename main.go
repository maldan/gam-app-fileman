package main

import (
	"embed"

	helloworld "github.com/maldan/gam-app-fileman/internal/app/fileman"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
