package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed images.png
var images []byte

//go:embed files/*.txt
var files embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("images_new.png", images, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := files.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := files.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
