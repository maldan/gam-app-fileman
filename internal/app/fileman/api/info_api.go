package api

import (
	"github.com/maldan/gam-app-fileman/internal/app/fileman/core"
)

type InfoApi struct {
}

func (r InfoApi) GetIndex(args ArgsNone) string {
	return core.README
}
