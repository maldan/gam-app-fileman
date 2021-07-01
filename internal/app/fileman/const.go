package fileman

import "time"

type File struct {
	Kind    string    `json:"kind"`
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	Created time.Time `json:"created"`
}

type Path struct {
	Path string `json:"path"`
}
