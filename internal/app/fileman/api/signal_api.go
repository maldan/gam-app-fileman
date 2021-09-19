package api

import (
	"encoding/base64"

	"github.com/maldan/gam-app-fileman/internal/app/fileman/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
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
	//Path string
	//Url  string
}

type SignalDownloadPhoto struct {
	Path string
	Url  string
}

type SignalDownloadBlobPhoto struct {
	Path string
	Data string
}

// Download list
func (r SignalApi) PostApplicationData(args ArgsApplicationData) {
	// Read base info
	x := SignalApplicationData{}
	err := cmhp_file.ReadJSON(args.Path, &x)
	if err != nil {
		restserver.Fatal(500, "", "", err.Error())
	}

	// Download photo
	if x.Type == "download_photo" {
		x := SignalDownloadPhoto{}
		cmhp_file.ReadJSON(args.Path, &x)

		r := DownloadApi{}
		r.PostIndex(core.Download{
			Url:  x.Url,
			Path: x.Path,
		})
		cmhp_file.Delete(args.Path)
		return
	}

	// Download blob photo
	if x.Type == "download_blob_photo" {
		x := SignalDownloadBlobPhoto{}
		cmhp_file.ReadJSON(args.Path, &x)

		// Decode and store photo
		photo, _ := base64.StdEncoding.DecodeString(x.Data)
		cmhp_file.WriteBin(x.Path+"/"+cmhp_crypto.UID(10)+".jpeg", photo)

		cmhp_file.Delete(args.Path)
		return
	}
}
