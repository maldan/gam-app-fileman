package api

import (
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/maldan/gam-app-fileman/internal/app/fileman/core"
	"github.com/maldan/gam-app-fileman/internal/app/fileman/plugin"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_net"
	"github.com/maldan/go-restserver"
)

type DownloadApi struct {
}

type WriteCounter struct {
	Total    uint64
	Download *core.Download
}

var semaphoreChan = make(chan int, 3)

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.Download.Progress = float64(wc.Total) / float64(wc.Download.TotalSize)
	return n, nil
}

var DownloadList = make([]*core.Download, 0)

// Download list
func (r DownloadApi) GetList() []*core.Download {
	dl := DownloadList
	sort.SliceStable(dl, func(i, j int) bool {
		return dl[i].Created.Unix() > dl[j].Created.Unix()
	})
	return dl
}

// Download
func (r DownloadApi) PostIndex(args core.Download) {
	args.Created = time.Now()
	args.Name = cmhp_crypto.UID(16)
	DownloadList = append(DownloadList, &args)
	os.MkdirAll(args.Path, 0777)

	if strings.Contains(args.Url, "xvideos.com") {
		go DownloadFromXvideos(&args, semaphoreChan)
		return
	}

	if strings.Contains(args.Url, "simply-hentai.com") {
		go plugin.DownloadFromSimplyHentai(&args, semaphoreChan)
		return
	}

	if strings.Contains(args.Url, "userapi.com") {
		go plugin.DownloadFromVK(&args, semaphoreChan)
		return
	}

	if strings.Contains(args.Url, "rule34.xxx") {
		go plugin.DownloadFromRule34(&args, semaphoreChan)
		return
	}

	if strings.Contains(args.Url, "discordapp.net") {
		go plugin.DownloadImage(&args, semaphoreChan)
		return
	}

	if strings.Contains(args.Url, ".m3u8") {
		go plugin.DownloadM3U8(&args, semaphoreChan)
		return
	}

	restserver.Fatal(500, restserver.ErrorType.Unknown, "url", "Url not supported")
}

func DownloadFromXvideos(args *core.Download, ch chan int) {
	ch <- 1
	defer func() {
		<-ch
	}()

	x := cmhp_net.Request(cmhp_net.HttpArgs{
		Url:    args.Url,
		Method: "GET",
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
		},
		Proxy: core.AppConfig.Proxy,
	})

	// Find url
	r, _ := regexp.Compile(`html5player\.setVideoUrlHigh\('(.*?)'\);`)
	xx := r.FindStringSubmatch(string(x.Body))

	// Create file
	out, err := os.Create(args.Path + args.Name + ".mp4")
	if err != nil {
		return
	}

	// Get the data
	resp, err := http.Get(xx[1])
	if err != nil {
		out.Close()
		return
	}
	defer resp.Body.Close()

	// Create our progress reporter and pass it to be used alongside our writer
	args.TotalSize = resp.ContentLength
	counter := &WriteCounter{
		Download: args,
	}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return
	}

	out.Close()
}
