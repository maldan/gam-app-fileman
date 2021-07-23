package fileman

import "time"

type File struct {
	Kind    string    `json:"kind"`
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	Created time.Time `json:"created"`
}

type CreateFile struct {
	Path  string `json:"path"`
	Files [][]byte
}

type Path struct {
	Path string `json:"path"`
}

type ArgsRename struct {
	From string `json:"from"`
	To   string `json:"to"`
}
