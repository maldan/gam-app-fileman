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

type Usage struct {
	Used  int64  `json:"used"`
	Total int64  `json:"total"`
	Fs    string `json:"fs"`
	Mount string `json:"mount"`
}

type CreateFile struct {
	Path  string `json:"path"`
	Files [][]byte
}

type Path struct {
	Path string `json:"path"`
}

type Config struct {
	Proxy string `json:"proxy"`
}

var DataDir = ""
var Host = ""
var Folder = ""
var AppConfig Config
