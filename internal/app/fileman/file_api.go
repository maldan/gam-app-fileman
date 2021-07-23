package fileman

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
	"github.com/maldan/go-cmhp/cmhp_slice"
	"github.com/maldan/go-restserver"
)

type FileApi struct {
}

// Get file content
func (f FileApi) GetFile(args Path) *os.File {
	file, err := os.OpenFile(args.Path, os.O_RDONLY, 0777)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "path", err.Error())
	}
	return file
}

func (f FileApi) GetList(args Path) []File {
	files, _ := cmhp_file.List(args.Path)
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
	if cmhp_slice.Includes([]interface{}{
		".jpeg", ".png", ".gif", ".jpg", ".webp",
	}, filepath.Ext(args.Path)) {
		appName = "dev/gallery"
	}

	// If video
	if cmhp_slice.Includes([]interface{}{
		".mp4",
	}, filepath.Ext(args.Path)) {
		appName = "dev/player"
	}

	if appName == "" {
		return
	}

	cmhp_process.Exec(
		"gam",
		"run",
		appName,
		fmt.Sprintf("--host=%v", Host),
		fmt.Sprintf("--file=%v", args.Path),
		fmt.Sprintf("--folder=%v", filepath.Dir(args.Path)),
	)
}

// Add new file sx
func (r FileApi) PostFile(args CreateFile) {
	if len(args.Files) > 0 {
		cmhp_file.WriteBin(args.Path, args.Files[0])
	} else {
		cmhp_file.WriteBin(args.Path, []byte{})
	}
}

// Rename
func (r FileApi) PostRename(args ArgsRename) {
	err := os.Rename(args.From, args.To)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "from", err.Error())
	}
}

// Add new folder
func (r FileApi) PostFolder(args CreateFile) {
	os.MkdirAll(args.Path, 0777)
}

// Add new folder
func (r FileApi) DeleteFile(args CreateFile) {
	cmhp_file.Delete(args.Path)
}

// Delete folder
func (r FileApi) DeleteFolder(args CreateFile) {
	if cmhp_slice.Includes([]interface{}{
		"/", "/home", "/root", "/var", "/www", "/etc", "/bin", "/boot", "/mnt", "/media",
		"/var", "/usr", "/tmp",
	}, args.Path) {
		return
	}
	cmhp_file.DeleteDir(args.Path)
}
