package main

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
)

func Screenshot() (*image.RGBA) {
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err!= nil {
		panic(err)
	}
	fileName := fmt.Sprintf("%d_%dx%d.png", 0, bounds.Dx(), bounds.Dy())
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)
	return img
}
func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, "<html><head></head><body><img  src=\"0_1920x1080.png\"></body></html>")
		return
	}
}
func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/static/", func(writer http.ResponseWriter, request *http.Request) {
		Screenshot()
		http.ServeFile(writer, request, "0_1920x1080.png")
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Listen And Server", err.Error())
	}
}


