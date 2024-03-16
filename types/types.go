package types

type File struct {
	Name     string `json:"name"`
	IsDir    bool   `json:"isDir"`
	Path     string `json:"path"`
	IsImage  bool   `json:"isImage"`
	Children []File `json:"children"`
}
type FileFlat struct {
	Name    string `json:"name"`
	IsDir   bool   `json:"isDir"`
	Path    string `json:"path"`
	IsImage bool   `json:"isImage"`
}
type DirPath struct {
	Dir string `json:"dir"`
}
