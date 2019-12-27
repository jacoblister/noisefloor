package vfs

import (
	"io"
	"net/http"
	"os"
)

// A FileSystem implements access to a collection of named files.
// The interface is compatible with net/http/fs
type FileSystem interface {
	Open(name string) (File, error)
	Create(name string) (File, error)
}

// A File is returned by a FileSystem's Open method
// The methods should behave the same as those on an *os.File.
type File interface {
	io.Closer
	io.Reader
	io.Writer
	io.Seeker
	Readdir(count int) ([]os.FileInfo, error)
	Stat() (os.FileInfo, error)
}

var defaultFS FileSystem

// DefaultFS returns the application wide default file system
func DefaultFS() FileSystem {
	return defaultFS
}

// SetDefaultFS sets the application wide default file system
func SetDefaultFS(filesystem FileSystem) {
	defaultFS = filesystem
}

// HTTPFileSystem is an interface compaitible with http.FileSystem
type HTTPFileSystem struct {
	filesystem FileSystem
}

// Open returns a file compatble with http.FileSystem
func (fs HTTPFileSystem) Open(name string) (http.File, error) {
	cwd, _ := os.Getwd()
	return fs.filesystem.Open(cwd + name)
}

// MakeHTTPFileSystem returns an http.FileSystem from a filesystem
func MakeHTTPFileSystem(filesystem FileSystem) http.FileSystem {
	return HTTPFileSystem{filesystem: filesystem}
}
