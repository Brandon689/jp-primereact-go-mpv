package handlers

import (
	"encoding/json"
	"github.com/Brandon689/jp-primereact-go-mpv/types"
	"net/http"
)

func HandleGetFilesAndDirectories(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var localPath types.DirPath
	err := json.NewDecoder(r.Body).Decode(&localPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	files := ListFilesAll(localPath.Dir)
	respondWithJSON(w, files)
}
