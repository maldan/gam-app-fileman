package core

import "time"

type File struct {
	Kind    string    `json:"kind"`
	Name    string    `json:"name"`
	User    string    `json:"user"`
	Size    int64     `json:"size"`
	Created time.Time `json:"created"`
}

type Download struct {
	Url       string    `json:"url"`
	Path      string    `json:"path"`
	Name      string    `json:"name"`
	Progress  float64   `json:"progress"`
	TotalSize int64     `json:"totalSize"`
	Created   time.Time `json:"created"`
}

type CreateFile struct {
	Path  string `json:"path"`
	Files [][]byte
}

type Path struct {
	Path string `json:"path"`
}

var DataDir = ""
var Host = ""
var Folder = ""
