package main

import (
	"encoding/json"
	"fmt"
	"github.com/Brandon689/jp-primereact-go-mpv/handlers"
	"net/http"
)

func main() {

	v := handlers.ListFilesAll("C:\\2024\\6\\golang-gpt-todo-echo-gorm")
	jsonData, _ := json.Marshal(v)
	fmt.Println(string(jsonData))
	fmt.Println(v)

	http.HandleFunc("/files", handlers.HandleGetFilesAndDirectories)

	http.Handle("/", http.FileServer(http.Dir("./frontend/primereact-ui/dist")))
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
