package plugin

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/maldan/gam-app-fileman/internal/app/fileman/core"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_net"
)

type SimplyHentaiPageSize struct {
	Full string `json:"full"`
}

type SimplyHentaiPage struct {
	Id    int                  `json:"id"`
	Sizes SimplyHentaiPageSize `json:"sizes"`
}

type SimplyHentaiData struct {
	Pages []SimplyHentaiPage `json:"pages"`
}

type SimplyHentai struct {
	Data SimplyHentaiData `json:"data"`
}

func DownloadFromSimplyHentai(args *core.Download, ch chan int) {
	ch <- 1
	defer func() {
		<-ch
	}()

	// Find url
	r, _ := regexp.Compile(`simply-hentai\.com\/(.*?)\/page`)
	xx := r.FindStringSubmatch(string(args.Url))
	xxx := strings.Split(xx[1], "/")[1]
	outto := SimplyHentai{}
	args.Name = xxx

	cmhp_net.Request(cmhp_net.HttpArgs{
		Url:    "https://api.simply-hentai.com/v3/album/" + xxx + "/pages",
		Method: "GET",
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
		},
		OutputJSON: &outto,
	})

	os.MkdirAll(args.Path+args.Name, 0777)

	for i := 0; i < len(outto.Data.Pages); i++ {
		image := cmhp_net.Request(cmhp_net.HttpArgs{
			Url:    outto.Data.Pages[i].Sizes.Full,
			Method: "GET",
			Headers: map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
			},
			Proxy: core.AppConfig.Proxy,
		})
		args.Progress += 1.0 / float64(len(outto.Data.Pages))

		if image.StatusCode == 200 {
			cmhp_file.WriteBin(fmt.Sprintf("%s/p_%03d.jpg", args.Path+args.Name, i), image.Body)
		}
	}
}
