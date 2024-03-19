package handlers

import (
	"encoding/json"
	"github.com/Brandon689/jp-primereact-go-mpv/types"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ListFilesAll(directory string) types.File {
	var result types.File
	result.Name = filepath.Base(directory)
	result.IsDir = true

	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip adding the root directory itself
		if path == directory {
			return nil
		}

		// Extract the relative path
		relPath, err := filepath.Rel(directory, path)
		if err != nil {
			return err
		}

		// Split the path into individual components
		parts := strings.Split(relPath, string(filepath.Separator))

		// Traverse the result to add directories
		currentDir := &result
		for _, part := range parts[:len(parts)-1] {
			found := false
			// Check if the current part already exists in the children
			for i := range currentDir.Children {
				if currentDir.Children[i].Name == part {
					currentDir = &currentDir.Children[i]
					found = true
					break
				}
			}
			// If not found, create a new directory and append it
			if !found {
				newDir := types.File{Name: part, IsDir: true}
				currentDir.Children = append(currentDir.Children, newDir)
				currentDir = &currentDir.Children[len(currentDir.Children)-1]
			}
		}

		// Add the file to the current directory's children
		currentDir.Children = append(currentDir.Children, types.File{Name: parts[len(parts)-1], IsDir: info.IsDir(), Path: path})

		return nil
	})

	return result
}

//func ListFilesAll(directory string) []types.File {
//	var result types.File
//	result.Name = filepath.Base(directory)
//	result.IsDir = true
//	var garden []types.File
//
//	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
//		if err != nil {
//			return err
//		}
//		// Skip adding the root directory itself
//		if path == directory {
//			return nil
//		}
//		// Extract the relative path
//		relPath, err := filepath.Rel(directory, path)
//		if err != nil {
//			return err
//		}
//
//		// Split the path into individual components
//		parts := strings.Split(relPath, string(filepath.Separator))
//
//		// Traverse the result to add directories
//		currentDir := &result
//		for _, part := range parts[:len(parts)-1] {
//			found := false
//			// Check if the current part already exists in the children
//			for i := range currentDir.Children {
//				if currentDir.Children[i].Name == part {
//					currentDir = &currentDir.Children[i]
//					found = true
//					break
//				}
//			}
//			// If not found, create a new directory and append it
//			if !found {
//				newDir := types.File{Name: part, IsDir: true}
//				currentDir.Children = append(currentDir.Children, newDir)
//				currentDir = &currentDir.Children[len(currentDir.Children)-1]
//
//				garden = append(garden, types.File{Name: part, IsDir: true})
//			}
//		}
//		// Add the file to the current directory's children
//		currentDir.Children = append(currentDir.Children, types.File{Name: parts[len(parts)-1], IsDir: info.IsDir(), Path: path})
//		garden = append(garden, types.File{Name: parts[len(parts)-1], IsDir: info.IsDir(), Path: path})
//
//		return nil
//	})
//	return garden
//}

func listFiles(directory string) []types.File {
	var files []types.File
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		var isImage = isImage(path)
		if isImage {
			copyImage(path, "./file-browser/dist/thumbnails/"+info.Name())
		}
		files = append(files, types.File{Name: info.Name(), IsDir: info.IsDir(), Path: path, IsImage: isImage})
		return nil
	})
	return files
}

func isImage(filename string) bool {
	// Convert filename to lowercase for case-insensitive comparison
	filename = strings.ToLower(filename)

	// Check if the filename ends with a known image extension
	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".avif"}
	for _, ext := range imageExtensions {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}

func copyImage(sourcePath, destPath string) error {
	// Open the source image file
	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()

	// Check if the destination file already exists
	_, err = os.Stat(destPath)
	if err == nil {
		return os.ErrExist // Return an error indicating that the destination file already exists
	} else if !os.IsNotExist(err) {
		return err // Return any other error encountered while checking the destination file
	}

	// Create the destination image file
	destination, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destination.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	return nil
}

func respondWithJSON(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
