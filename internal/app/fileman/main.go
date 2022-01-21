package fileman

import (
	"embed"
	"flag"
	"fmt"
	"github.com/maldan/go-cmhp/cmhp_file"
	"os"

	"github.com/maldan/gam-app-fileman/internal/app/fileman/api"
	"github.com/maldan/gam-app-fileman/internal/app/fileman/core"
	"github.com/maldan/go-rapi"
	"github.com/maldan/go-rapi/rapi_core"
	"github.com/maldan/go-rapi/rapi_rest"
	"github.com/maldan/go-rapi/rapi_vfs"
)

func Start(frontFs embed.FS, readme string) {
	// Server
	var host = flag.String("host", "127.0.0.1", "Server Hostname")
	var port = flag.Int("port", 16000, "Server Port")
	_ = flag.Int("clientPort", 8080, "Client Port")

	// Data
	var dataDir = flag.String("dataDir", "db", "Data Directory")
	_ = flag.String("appId", "id", "App id")
	flag.Parse()

	// Set
	core.DataDir = *dataDir
	core.Host = *host
	core.README = readme

	// Read config
	_ = cmhp_file.ReadJSON(core.DataDir+"/config.json", &core.AppConfig)
	if core.AppConfig.ThumbnailCachePath == "" {
		core.AppConfig.ThumbnailCachePath = os.TempDir()
	}

	rapi.Start(rapi.Config{
		Host: fmt.Sprintf("%s:%d", *host, *port),
		Router: map[string]rapi_core.Handler{
			"/": rapi_vfs.VFSHandler{
				Root: "frontend/build",
				Fs:   frontFs,
			},
			"/api": rapi_rest.ApiHandler{
				Controller: map[string]interface{}{
					"file":     api.FileApi{},
					"download": api.DownloadApi{},
					"disk":     api.DiskApi{},
				},
			},
			"/system": rapi_rest.ApiHandler{
				Controller: map[string]interface{}{
					"signal": api.SignalApi{},
					"info":   api.InfoApi{},
				},
			},
		},
		DbPath: core.DataDir,
	})
}
