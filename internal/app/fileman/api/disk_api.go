package api

import (
	"regexp"
	"strings"

	"github.com/maldan/gam-app-fileman/internal/app/fileman/core"
	"github.com/maldan/go-cmhp/cmhp_convert"
	"github.com/maldan/go-cmhp/cmhp_process"
)

type DiskApi struct {
}

// Download list
func (r DiskApi) GetUsage() []core.Usage {
	x := cmhp_process.Exec("df", "--output=used,size,source,target")
	lines := strings.Split(x, "\n")
	space := regexp.MustCompile(`\s+`)
	usageList := make([]core.Usage, 0)

	for i, line := range lines {
		if i == 0 {
			continue
		}
		s := strings.Trim(space.ReplaceAllString(line, " "), " ")
		tt := strings.Split(s, " ")
		if len(tt) != 4 {
			continue
		}

		usageList = append(usageList, core.Usage{
			Used:  int64(cmhp_convert.StrToInt(tt[0])),
			Total: int64(cmhp_convert.StrToInt(tt[1])),
			Fs:    tt[2],
			Mount: tt[3],
		})
	}

	return usageList
}
