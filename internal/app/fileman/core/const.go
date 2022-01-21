package core

import (
	"time"
)

type File struct {
	Kind    string                 `json:"kind"`
	Path    string                 `json:"path"`
	Name    string                 `json:"name"`
	User    string                 `json:"user"`
	Size    int64                  `json:"size"`
	Tags    map[string]interface{} `json:"tags"`
	Created time.Time              `json:"created"`
}

type Config struct {
	Proxy              string `json:"proxy"`
	ThumbnailCachePath string `json:"thumbnailCachePath"`
}

var DataDir = ""
var AppConfig Config
