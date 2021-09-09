package plugin

import (
	"io"
	"net/http"
	"os"

	"github.com/maldan/gam-app-fileman/internal/app/fileman/core"
)

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

func DownloadFromVK(args *core.Download, ch chan int) {
	ch <- 1
	defer func() {
		<-ch
	}()

	// Create file
	out, err := os.Create(args.Path + args.Name + ".jpg")
	if err != nil {
		return
	}

	// Get the data
	resp, err := http.Get(args.Url)
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
