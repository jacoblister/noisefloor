package main

import (
	"os"
	"strings"
	"syscall/js"

	"github.com/jacoblister/noisefloor/pkg/vfs"
)

// FS is the Javascript filesystem (pass calls the fetch() api)
type FS struct{}

var fs FS

// Open implements filesystem for golang, just a passthrough to the OS method
func (fs FS) Open(name string) (vfs.File, error) {
	content := js.Global().Call("FetchFile", name)

	file := File{name: name, reader: strings.NewReader(content.String())}

	return file, nil
}

// Create implements filesystem for golang, just a passthrough to the OS method
func (fs FS) Create(name string) (vfs.File, error) {
	return nil, nil
}

// File is the Javascript filesystem file (stored as a string)
type File struct {
	name   string
	reader *strings.Reader
}

// Close closes an in memory file
func (f File) Close() error {
	return nil
}

// Read reads from an in memory file
func (f File) Read(p []byte) (n int, err error) {
	return f.reader.Read(p)
}

// Write writes to  an in memory file
func (f File) Write(p []byte) (n int, err error) {
	return 0, nil
}

// Seek seeks to a location in the memory file
func (f File) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

// Readdir reads the directory contents
func (f File) Readdir(count int) ([]os.FileInfo, error) {
	files := []os.FileInfo{}

	content := js.Global().Call("FetchFile", f.name+"/$")
	lines := strings.Split(content.String(), "\n")
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) > 0 {
			files = append(files, FileInfo{name: lines[i]})
		}
	}

	return files, nil
}

// Stat gets the status of an in memory file
func (f File) Stat() (os.FileInfo, error) {
	panic("not implemented")
}
