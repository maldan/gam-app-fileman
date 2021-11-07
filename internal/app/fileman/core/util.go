package core

import (
	"os"
	"strings"

	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
	"github.com/maldan/go-restserver"
)

func GetFileHash(path string) string {
	// Get path hash
	pathHash := cmhp_crypto.Sha1(path)
	if cmhp_file.Exists("/tmp/file_hash/" + pathHash) {
		hash, _ := cmhp_file.ReadText("/tmp/file_hash/" + pathHash)
		return hash
	}

	// Get file hash
	h, _ := cmhp_process.Exec("sha1sum", path)
	fileHash := strings.ReplaceAll(h, "\n", "")
	if fileHash == "" {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "path", "Can't get file hash")
	}
	finalHash := strings.Trim(strings.Split(fileHash, " ")[0], " ")
	os.MkdirAll("/tmp/file_hash", 0777)
	cmhp_file.WriteText("/tmp/file_hash/"+pathHash, finalHash)
	return finalHash
}
