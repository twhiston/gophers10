package main

import (
	"net/http"
	"github.com/gobuffalo/packr"
	"log"
)

func getAssetBox() packr.Box {
	return packr.NewBox("./assets")
}


func main() {
	box := getAssetBox()
	http.Handle("/", http.FileServer(box))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}