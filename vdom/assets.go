// +build ignore

package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var cwd, _ = os.Getwd()
	fs := http.Dir(filepath.Join(cwd, "assets"))
	err := vfsgen.Generate(fs, vfsgen.Options{PackageName: "vdom"})
	if err != nil {
		log.Fatalln(err)
	}
}
