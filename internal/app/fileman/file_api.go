package fileman

import (
	"fmt"

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
