package api

type ArgsFileInfo struct {
	Path string `json:"path"`
	Data string `json:"data"`
}
type ArgsRename struct {
	From string `json:"from"`
	To   string `json:"to"`
}
type ArgsVideoPreview struct {
	Path string `json:"path"`
	Time string `json:"time"`
}
type ArgsTags struct {
	Path string                 `json:"path"`
	Tags map[string]interface{} `json:"tags"`
}
