package api

type ArgsFileInfo struct {
	Path string `json:"path"`
	Data string `json:"data"`
}
type ArgsRename struct {
	From string `json:"from"`
	To   string `json:"to"`
}
