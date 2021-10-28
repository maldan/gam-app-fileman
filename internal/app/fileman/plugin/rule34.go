package plugin

import (
	"os"

	"github.com/maldan/gam-app-fileman/internal/app/fileman/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_net"
)

func DownloadFromRule34(args *core.Download, ch chan int) {
	ch <- 1
	defer func() {
		<-ch
	}()

	// Find url
	image := cmhp_net.Request(cmhp_net.HttpArgs{
		Url:    args.Url,
		Method: "GET",
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
		},
		Proxy: core.AppConfig.Proxy,
	})

	os.MkdirAll(args.Path, 0777)
	cmhp_file.WriteBin(args.Path+args.Name+".jpg", image.Body)
}
