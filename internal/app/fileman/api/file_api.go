package api

import (
	"fmt"
	"github.com/maldan/gam-app-fileman/internal/app/fileman/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
	"github.com/maldan/go-cmhp/cmhp_slice"
	"github.com/maldan/go-rapi/rapi_core"
	"github.com/maldan/go-restserver"
	"os"
	"os/user"
	"path"
	"strconv"
	"strings"
	"syscall"
)

type FileApi struct {
}

var CurrentPath = "/"

func (f FileApi) GetPath() string {
	return CurrentPath
}

func (f FileApi) PostPath(args core.Path) {
	CurrentPath = args.Path
}

// GetVideoThumbnail Get video thumbnail
func (f FileApi) GetVideoThumbnail(ctx *rapi_core.Context, args ArgsVideoPreview) string {
	ctx.IsServeFile = true

	// Temp file
	_ = os.MkdirAll(core.AppConfig.ThumbnailCachePath+"/video_preview/", 0777)
	hash := cmhp_crypto.Sha1([]byte(args.Path + "" + args.Time))
	f1 := hash[0:2]
	f2 := hash[2:4]
	_ = os.MkdirAll(core.AppConfig.ThumbnailCachePath+"/video_preview/"+f1+"/"+f2, 0777)
	tmpFile := core.AppConfig.ThumbnailCachePath + "/video_preview/" + f1 + "/" + f2 + "/" + hash

	// If file exists
	if cmhp_file.Exists(tmpFile) {
		return tmpFile
	}

	// Generate preview
	_, err := cmhp_process.Exec("ffmpeg",
		"-ss", args.Time,
		"-i", args.Path,
		"-frames:v", "1",
		"-f", "webp",
		"-vf", "scale=320:200:force_original_aspect_ratio=decrease",
		tmpFile, "-y")
	rapi_core.FatalIfError(err)

	return tmpFile
}

// GetImageThumbnail Get image thumbnail
func (f FileApi) GetImageThumbnail(ctx *rapi_core.Context, args core.Path) string {
	ctx.IsServeFile = true

	// Temp file
	_ = os.MkdirAll(core.AppConfig.ThumbnailCachePath+"/image_preview/", 0777)
	hash := cmhp_crypto.Sha1([]byte(args.Path))
	f1 := hash[0:2]
	f2 := hash[2:4]
	_ = os.MkdirAll(core.AppConfig.ThumbnailCachePath+"/image_preview/"+f1+"/"+f2, 0777)
	tmpFile := core.AppConfig.ThumbnailCachePath + "/image_preview/" + f1 + "/" + f2 + "/" + hash

	// If file exists
	if cmhp_file.Exists(tmpFile) {
		return tmpFile
	}

	// Generate preview
	_, err := cmhp_process.Exec(
		"magick", args.Path,
		"-quality", "70",
		// "-define", "g:format=png",
		"-thumbnail", "192x192^",
		"-gravity", "center",
		"-extent", "192x192",
		tmpFile)

	rapi_core.FatalIfError(err)

	return tmpFile
}

// PostSetVideoPreview Set video thumbnail
func (f FileApi) PostSetVideoPreview(args ArgsVideoPreview) {
	// Get tags
	tags := f.GetTags(ArgsTags{Path: args.Path})
	tags["videoPreviewTime"] = args.Time
	f.PostTags(ArgsTags{Path: args.Path, Tags: tags})
}

// Get file content
func (f FileApi) GetFile(ctx *rapi_core.Context, args core.Path) string {
	ctx.IsServeFile = true
	return args.Path
}

// Get file content
func (f FileApi) GetDirSize(args core.Path) int64 {
	out, _ := cmhp_process.Exec("du", "-d", "0", "-b", args.Path)
	i, _ := strconv.ParseInt(strings.Split(out, "\t")[0], 10, 64)
	return i
}

// GetList Get list
func (f FileApi) GetList(args core.Path) []core.File {
	out := make([]core.File, 0)
	files, err := cmhp_file.List(args.Path)
	rapi_core.FatalIfError(err)

	for _, file := range files {
		kind := "file"
		if file.IsDir() {
			kind = "dir"
		}
		if fmt.Sprintf("%v", file.Mode())[0] == 'L' {
			kind = "dir"
		}

		// Create file
		filePath := strings.ReplaceAll(args.Path+"/"+file.Name(), "//", "/")
		outFile := core.File{
			Path:    filePath,
			Name:    file.Name(),
			User:    "",
			Kind:    kind,
			Size:    file.Size(),
			Created: file.ModTime(),
			Tags:    f.GetTags(ArgsTags{Path: filePath}),
		}

		// Get user
		info, err := os.Stat(args.Path + "/" + file.Name())
		if err == nil {
			if stat, ok := info.Sys().(*syscall.Stat_t); ok {
				u := strconv.FormatUint(uint64(stat.Uid), 10)
				uu, err1 := user.LookupId(u)
				if err1 == nil {
					outFile.User = uu.Username
				}
			}
		}

		// Append
		out = append(out, outFile)
	}
	return out
}

// PostOpen Open file
/*func (f FileApi) PostOpen(args core.Path) {
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
}*/

// Add new file
func (f FileApi) PostFile(args core.CreateFile) {
	cmhp_file.WriteBin(args.Path, args.File.Data)
}

// PostRename Rename
func (f FileApi) PostRename(args ArgsRename) {
	err := os.Rename(args.From, args.To)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "from", err.Error())
	}
}

// Add new dir
func (f FileApi) PostDir(args core.CreateFile) {
	os.MkdirAll(args.Path, 0777)
}

// Delete file
func (f FileApi) DeleteFile(args core.CreateFile) {
	cmhp_file.Delete(args.Path)
}

// Delete dir
func (f FileApi) DeleteDir(args core.CreateFile) {
	if cmhp_slice.Includes([]interface{}{
		"/", "/home", "/root", "/var", "/www", "/etc", "/bin", "/boot", "/mnt", "/media",
		"/var", "/usr", "/tmp",
	}, args.Path) {
		return
	}
	cmhp_file.DeleteDir(args.Path)
}

// Delete any
func (f FileApi) DeleteAny(args core.CreateFile) {
	s, _ := os.Stat(args.Path)
	if s.IsDir() {
		f.DeleteDir(args)
	} else {
		f.DeleteFile(args)
	}
}

// Copy file
func (f FileApi) PostCopy(args ArgsRename) {
	cmhp_process.Exec("cp", "-r", args.From, args.To)
}

// Copy file
func (f FileApi) PostMove(args ArgsRename) {
	cmhp_process.Exec("mv", args.From, args.To)
}

// Save file info
/*func (f FileApi) PostInfo(args ArgsFileInfo) {
	fileHash := core.GetFileHash(args.Path)
	os.MkdirAll(core.DataDir+"/file_info", 0777)
	cmhp_file.WriteText(core.DataDir+"/file_info/"+fileHash+".json", args.Data)
}*/

// Get file info
/*func (f FileApi) GetInfo(args core.Path) string {
	fileHash := core.GetFileHash(args.Path)
	data, err := cmhp_file.ReadText(core.DataDir + "/file_info/" + fileHash + ".json")
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "path", "Can't get file hash")
	}
	return data
}*/

// Get file info
func (f FileApi) GetFullInfo(args core.Path) core.FileFullInfo {
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

// PostTags Get file hash and set name as it
func (f FileApi) PostTags(args ArgsTags) {
	fileHash := strings.Split(path.Base(args.Path), ".")[0]
	f1 := fileHash[0:2]
	f2 := fileHash[2:4]
	_ = os.MkdirAll(core.AppConfig.ThumbnailCachePath+"/tags/"+f1+"/"+f2, 0777)
	tagFile := core.AppConfig.ThumbnailCachePath + "/tags/" + f1 + "/" + f2 + "/" + fileHash

	if args.Tags == nil {
		args.Tags = make(map[string]interface{}, 0)
	}

	// Set default tags
	args.Tags["path"] = args.Path

	err := cmhp_file.WriteJSON(tagFile, &args.Tags)
	rapi_core.FatalIfError(err)
}

// GetTags Get file tags
func (f FileApi) GetTags(args ArgsTags) map[string]interface{} {
	tags := make(map[string]interface{}, 0)

	fileHash := strings.Split(path.Base(args.Path), ".")[0]
	if len(fileHash) < 4 {
		return tags
	}

	f1 := fileHash[0:2]
	f2 := fileHash[2:4]
	tagFile := core.AppConfig.ThumbnailCachePath + "/tags/" + f1 + "/" + f2 + "/" + fileHash

	cmhp_file.ReadJSON(tagFile, &tags)
	return tags
}

// PostSetHashName Get file hash and set name as it
func (f FileApi) PostSetHashName(args core.Path) {
	hashName := core.Sha1Sum(args.Path)
	// hashName, err := cmhp_file.HashSha1(args.Path)
	mimeType := core.GetFileMimeType(args.Path)
	// rapi_core.FatalIfError(err)
	finalName := ""

	// Check mime type
	if mimeType == "image/png" {
		finalName = hashName + ".png"
	} else if mimeType == "image/jpeg" {
		finalName = hashName + ".jpeg"
	} else if mimeType == "image/gif" {
		finalName = hashName + ".gif"
	} else if mimeType == "image/webp" {
		finalName = hashName + ".webp"
	} else if mimeType == "video/mp4" || mimeType == "video/quicktime" {
		finalName = hashName + ".mp4"
	} else {
		rapi_core.Fatal(rapi_core.Error{Code: 500, Description: "Unsupported file format"})
	}

	// Check if file already exists
	newPath := path.Dir(args.Path) + "/" + finalName
	if cmhp_file.Exists(newPath) {
		rapi_core.Fatal(rapi_core.Error{Code: 500, Description: "File already exists"})
	}

	f.PostRename(ArgsRename{From: args.Path, To: newPath})
}
