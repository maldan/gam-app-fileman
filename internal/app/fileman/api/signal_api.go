package api

import (
	"github.com/maldan/gam-app-fileman/internal/app/fileman/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-restserver"
)

type SignalApi struct {
}

type ArgsApplicationData struct {
	Path string
}

type SignalApplicationData struct {
	Type string
	Path string
	Url  string
}

// Download list
func (r SignalApi) PostApplicationData(args ArgsApplicationData) {
	x := SignalApplicationData{}
	err := cmhp_file.ReadJSON(args.Path, &x)
	if err != nil {
		restserver.Fatal(500, "", "", err.Error())
	}
	if x.Type == "download_photo" {
		r := DownloadApi{}
		r.PostIndex(core.Download{
			Url:  x.Url,
			Path: x.Path,
		})
		cmhp_file.Delete(args.Path)
	}
}
