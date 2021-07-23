package fileman

import (
	"fmt"
	"path/filepath"

	"github.com/maldan/go-cmhp"
)

type FileApi struct {
}

func (f FileApi) GetList(args Path) []File {
	files, _ := cmhp.FileList(args.Path)
	out := make([]File, 0)
	for _, file := range files {
		kind := "file"
		if file.IsDir() {
			kind = "dir"
		}
		if fmt.Sprintf("%v", file.Mode())[0] == 'L' {
			kind = "dir"
		}
		out = append(out, File{
			Name:    file.Name(),
			Kind:    kind,
			Size:    file.Size(),
			Created: file.ModTime(),
		})
	}
	return out
}

func (f FileApi) PostOpen(args Path) {
	appName := ""

	// If image
	if cmhp.SliceIncludes([]interface{}{
		".jpeg", ".png", ".gif", ".jpg", ".webp",
	}, filepath.Ext(args.Path)) {
		appName = "dev/gallery"
	}

	// If video
	if cmhp.SliceIncludes([]interface{}{
		".mp4",
	}, filepath.Ext(args.Path)) {
		appName = "dev/player"
	}

	if appName == "" {
		return
	}

	cmhp.ProcessExec(
		"gam",
		"run",
		appName,
		fmt.Sprintf("--host=%v", Host),
		fmt.Sprintf("--file=%v", args.Path),
		fmt.Sprintf("--folder=%v", filepath.Dir(args.Path)),
	)
}
