package golang_pzn_embed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var file string

func TestString(t *testing.T) {
	fmt.Println(file)
}
