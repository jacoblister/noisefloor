// +build ignore

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/shurcooL/vfsgen"
)

func main() {
	var cwd, _ = os.Getwd()
	fs := http.Dir(filepath.Join(cwd, "files"))
	err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "assets",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}

	// add !js build tag
	filename := "assets_vfsdata.go"
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
		return
	}

	input = append([]byte("// +build !js\n"), input...)
	err = ioutil.WriteFile(filename, input, 0644)
	if err != nil {
		log.Fatalln("Error writing output:", filename)
		log.Fatalln(err)
		return
	}
}
