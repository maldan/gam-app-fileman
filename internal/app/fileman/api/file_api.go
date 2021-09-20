package api

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/maldan/gam-app-fileman/internal/app/fileman/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
	"github.com/maldan/go-cmhp/cmhp_slice"
	"github.com/maldan/go-restserver"
)

type FileApi struct {
}

var CurrentPath = "/"

func (r FileApi) GetPath() string {
	return CurrentPath
}

func (r FileApi) PostPath(args core.Path) {
	CurrentPath = args.Path
}

// Get video thumbnail
func (f FileApi) GetVideoThumbnail(args core.Path) *os.File {
	// Temp file
	os.MkdirAll(os.TempDir()+"/video_preview/", 0777)
	tmpFile := os.TempDir() + "/video_preview/" + cmhp_crypto.Sha1(args.Path) + ".webp"

	// If file exists
	if cmhp_file.Exists(tmpFile) {
		f, err := os.OpenFile(tmpFile, os.O_RDONLY, 0777)
		if err != nil {
			restserver.Fatal(500, restserver.ErrorType.Unknown, "path", err.Error())
		}
		return f
	}

	// Generate preview
	cmhp_process.Exec("ffmpeg",
		"-i", args.Path, "-vf",
		"select=eq(n\\,320)", "-frames:v", "1",
		"-s", "256x256", tmpFile, "-y")

	f1, err1 := os.OpenFile(tmpFile, os.O_RDONLY, 0777)
	if err1 != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "path", err1.Error())
	}
	return f1
}

// Get image thumbnail
func (f FileApi) GetImageThumbnail(args core.Path) *os.File {
	// Temp file
	os.MkdirAll(os.TempDir()+"/image_preview/", 0777)
	tmpFile := os.TempDir() + "/image_preview/" + cmhp_crypto.Sha1(args.Path) + ".webp"

	// If file exists
	if cmhp_file.Exists(tmpFile) {
		f, err := os.OpenFile(tmpFile, os.O_RDONLY, 0777)
		if err != nil {
			restserver.Fatal(500, restserver.ErrorType.Unknown, "path", err.Error())
		}
		return f
	}

	// Generate preview
	cmhp_process.Exec("magick", args.Path, "-quality", "80", "-define", "webp:lossless=false", "-thumbnail", "256x256^", "-gravity", "center", "-extent", "256x256", tmpFile)

	f1, err1 := os.OpenFile(tmpFile, os.O_RDONLY, 0777)
	if err1 != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "path", err1.Error())
	}
	return f1
}

// Get file content
func (f FileApi) GetFile(args core.Path) *os.File {
	file, err := os.OpenFile(args.Path, os.O_RDONLY, 0777)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "path", err.Error())
	}
	return file
}

// Get file content
func (f FileApi) GetDirSize(args core.Path) int64 {
	out := cmhp_process.Exec("du", "-d", "0", "-b", args.Path)
	i, _ := strconv.ParseInt(strings.Split(out, "\t")[0], 10, 64)
	return i
}

// Get list
func (r FileApi) GetList(args core.Path) []core.File {
	files, _ := cmhp_file.List(args.Path)
	out := make([]core.File, 0)
	for _, file := range files {
		kind := "file"
		if file.IsDir() {
			kind = "dir"
		}
		if fmt.Sprintf("%v", file.Mode())[0] == 'L' {
			kind = "dir"
		}

		// Create file
		outFile := core.File{
			Name:    file.Name(),
			User:    "",
			Kind:    kind,
			Size:    file.Size(),
			Created: file.ModTime(),
		}

		// Get user
		info, err := os.Stat(args.Path + "/" + file.Name())
		if err == nil {
			if stat, ok := info.Sys().(*syscall.Stat_t); ok {
				u := strconv.FormatUint(uint64(stat.Uid), 10)
				uu, _ := user.LookupId(u)
				outFile.User = uu.Username
			}
		}

		// Append
		out = append(out, outFile)
	}
	return out
}

// Open file
func (f FileApi) PostOpen(args core.Path) {
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
		fmt.Sprintf("--host=%v", core.Host),
		fmt.Sprintf("--file=%v", args.Path),
		fmt.Sprintf("--folder=%v", filepath.Dir(args.Path)),
	)
}

// Add new file
func (r FileApi) PostFile(args core.CreateFile) {
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

// Add new dir
func (r FileApi) PostDir(args core.CreateFile) {
	os.MkdirAll(args.Path, 0777)
}

// Delete file
func (r FileApi) DeleteFile(args core.CreateFile) {
	cmhp_file.Delete(args.Path)
}

// Delete dir
func (r FileApi) DeleteDir(args core.CreateFile) {
	if cmhp_slice.Includes([]interface{}{
		"/", "/home", "/root", "/var", "/www", "/etc", "/bin", "/boot", "/mnt", "/media",
		"/var", "/usr", "/tmp",
	}, args.Path) {
		return
	}
	cmhp_file.DeleteDir(args.Path)
}

// Delete any
func (r FileApi) DeleteAny(args core.CreateFile) {
	s, _ := os.Stat(args.Path)
	if s.IsDir() {
		r.DeleteDir(args)
	} else {
		r.DeleteFile(args)
	}
}

// Copy file
func (r FileApi) PostCopy(args ArgsRename) {
	cmhp_process.Exec("cp", "-r", args.From, args.To)
}

// Copy file
func (r FileApi) PostMove(args ArgsRename) {
	cmhp_process.Exec("mv", args.From, args.To)
}

// Save file info
func (r FileApi) PostInfo(args ArgsFileInfo) {
	fileHash := core.GetFileHash(args.Path)
	os.MkdirAll(core.DataDir+"/file_info", 0777)
	cmhp_file.WriteText(core.DataDir+"/file_info/"+fileHash+".json", args.Data)
}

// Get file info
func (r FileApi) GetInfo(args core.Path) string {
	fileHash := core.GetFileHash(args.Path)
	data, err := cmhp_file.ReadText(core.DataDir + "/file_info/" + fileHash + ".json")
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "path", "Can't get file hash")
	}
	return data
}

// Get file info
func (r FileApi) GetFullInfo(args core.Path) core.FileFullInfo {
	info, err := os.Stat(args.Path)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "path", err.Error())
	}

	// Kind
	kind := "file"
	if info.IsDir() {
		kind = "dir"
	}
	if fmt.Sprintf("%v", info.Mode())[0] == 'L' {
		kind = "dir"
	}

	// Final info
	fullInfo := core.FileFullInfo{
		Name:    info.Name(),
		Kind:    kind,
		Size:    info.Size(),
		Created: info.ModTime(),
	}

	// Get user
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		u := strconv.FormatUint(uint64(stat.Uid), 10)
		uu, _ := user.LookupId(u)
		fullInfo.User = uu.Username
	}

	return fullInfo
}
