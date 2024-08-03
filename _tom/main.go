package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/images", handleImages)
	http.ListenAndServe(":8080", nil)

	fmt.Println("Listening on port 8080!")
}

func handleImages(w http.ResponseWriter, r *http.Request) {
	imageDir := "images/" // Replace with the actual path to your image folder

	// Get a list of all image files in the directory
	imageFiles, err := filepath.Glob(filepath.Join(imageDir, "*.png"))
	if err != nil {
		http.Error(w, "Failed to read image directory", http.StatusInternalServerError)
		return
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Choose a random image file
	//fmt.Println(imageFiles)
	randomImage := imageFiles[rand.Intn(len(imageFiles))]

	// Open the image file
	file, err := os.Open(randomImage)
	if err != nil {
		http.Error(w, "Failed to open image file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Read the image file
	imageData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read image file", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate that the response is an image
	w.Header().Set("Content-Type", "image/png")

	// Write the image data to the response
	w.Write(imageData)
}
