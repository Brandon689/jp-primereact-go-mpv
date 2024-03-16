package main

import (
	"fmt"
	"github.com/Brandon689/jp-primereact-go-mpv/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/files", handlers.HandleFiles)
	http.HandleFunc("/files2", handlers.HandleFiles2)
	http.HandleFunc("/send-data", handlers.HandleData)
	http.HandleFunc("/send-path", handlers.HandlePath)
	http.HandleFunc("/files3", handlers.HandleFiles3)

	http.Handle("/", http.FileServer(http.Dir("./frontend/primereact-ui")))
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
