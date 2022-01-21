package plugin

import (
	"github.com/maldan/gam-app-fileman/internal/app/fileman/core"
	"github.com/maldan/go-cmhp/cmhp_convert"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_net"
	"github.com/maldan/go-cmhp/cmhp_string"
	"os"
	"strings"
)

type Stream struct {
	Url       string
	Bandwidth int
}

func getBestStream(url string) Stream {
	// Load m3u8 file
	response := cmhp_net.Request(cmhp_net.HttpArgs{
		Url:    url,
		Method: "GET",
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
		},
	})
	lines := strings.Split(string(response.Body), "\n")
	if response.StatusCode != 200 {
		panic("Can't get m3u8")
	}

	var bestStream Stream

	for index, line := range lines {
		if strings.Contains(line, "#EXT-X-STREAM-INF") {
			info := cmhp_string.ParseParameterList(strings.Replace(line, "#EXT-X-STREAM-INF:", "", 1), ",", "=")

			if cmhp_convert.StrToInt(info["BANDWIDTH"].(string)) > bestStream.Bandwidth {
				bestStream = Stream{
					Url:       lines[index+1],
					Bandwidth: cmhp_convert.StrToInt(info["BANDWIDTH"].(string)),
				}
			}
		}
	}

	return bestStream
}

func DownloadM3U8(args *core.Download, ch chan int) {
	ch <- 1
	defer func() {
		<-ch
	}()

	ss := strings.Split(args.Url, "/")
	rootUrl := strings.Join(ss[0:len(ss)-1], "/")

	// Load master
	bestStream := getBestStream(args.Url)

	// Load stream segment list
	indexList := cmhp_net.Request(cmhp_net.HttpArgs{
		Url:    rootUrl + "/" + bestStream.Url,
		Method: "GET",
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
		},
	})
	ssList := strings.Split(string(indexList.Body), "\n")
	segmentList := make([]string, 0)
	totalTime := 0.0
	for i, segment := range ssList {
		if strings.Contains(segment, "#EXTINF:") {
			time := cmhp_convert.StrToFloat(strings.ReplaceAll(strings.ReplaceAll(segment, "#EXTINF:", ""), ",", ""))
			totalTime += time
			segmentList = append(segmentList, ssList[i+1])
		}
	}

	args.TotalSize = int64(totalTime)
	stepProgress := 1.0 / float64(len(segmentList))

	// Prepare path
	os.MkdirAll(args.Path, 0777)

	for _, segment := range segmentList {
		videoPart := cmhp_net.Request(cmhp_net.HttpArgs{
			Url:    rootUrl + "/" + segment,
			Method: "GET",
			Headers: map[string]string{
				"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
			},
		})
		cmhp_file.Append(args.Path+args.Name+".mp4", videoPart.Body)
		args.Progress += stepProgress
	}
}
