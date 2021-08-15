package fileman

import "time"

type File struct {
	Kind    string    `json:"kind"`
	Name    string    `json:"name"`
	User    string    `json:"user"`
	Size    int64     `json:"size"`
	Created time.Time `json:"created"`
}

type ArgsFileInfo struct {
	Path string `json:"path"`
	Data string `json:"data"`
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
