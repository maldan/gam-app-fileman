package api

import (
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/maldan/gam-app-fileman/internal/app/fileman/core"
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

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.Download.Progress = float64(wc.Total) / float64(wc.Download.TotalSize)
	return n, nil
}

var DownloadList = make([]*core.Download, 0)

// Download list
func (r DownloadApi) GetList() []*core.Download {
	return DownloadList
}

// Download
func (r DownloadApi) PostIndex(args core.Download) {
	if !strings.Contains(args.Url, "xvideos.com") {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "url", "Url not supported")
	}

	args.Created = time.Now()
	args.Name = cmhp_crypto.UID(16)
	DownloadList = append(DownloadList, &args)
	go DownloadFromXvideos(&args)
}

func DownloadFromXvideos(args *core.Download) {
	x := cmhp_net.Request(cmhp_net.HttpArgs{
		Url:    args.Url,
		Method: "GET",
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
		},
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