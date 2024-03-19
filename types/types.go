package types

type File struct {
	Name     string `json:"name"`
	IsDir    bool   `json:"isDir"`
	Path     string `json:"path"`
	IsImage  bool   `json:"isImage"`
	Children []File `json:"children"`
}

type DirPath struct {
	Dir string `json:"dir"`
}
