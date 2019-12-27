package nf

import (
	"os"

	"github.com/jacoblister/noisefloor/pkg/vfs"
)

type FS struct{}

var fs FS

// Open implements filesystem for golang, just a passthrough to the OS method
func (fs FS) Open(name string) (vfs.File, error) {
	return os.Open(name)
}

// Create implements filesystem for golang, just a passthrough to the OS method
func (fs FS) Create(name string) (vfs.File, error) {
	return os.Create(name)
}
