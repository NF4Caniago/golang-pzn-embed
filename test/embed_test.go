package golang_pzn_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var file string

func TestString(t *testing.T) {
	fmt.Println(file)
}

//go:embed images.png
var images []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("images_new.png", images, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	b, _ := files.ReadFile("files/b.txt")
	c, _ := files.ReadFile("files/c.txt")

	fmt.Println(string(a))
	fmt.Println(string(b))
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
