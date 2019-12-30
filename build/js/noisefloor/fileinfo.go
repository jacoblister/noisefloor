package main

import (
	"os"
	"time"
)

//FileInfo is the Javascript file info abstraction
type FileInfo struct {
	name    string
	size    int64       // length in bytes for regular files; system-dependent for others
	mode    os.FileMode // file mode bits
	modTime time.Time   // modification time
}

//Name is the base name of the file
func (f FileInfo) Name() string {
	return f.name
}

//Size is the length in bytes for regular files; system-dependent for others
func (f FileInfo) Size() int64 {
	return f.size
}

//Mode is the file mode bits
func (f FileInfo) Mode() os.FileMode {
	return f.mode
}

//ModTime is the modification time
func (f FileInfo) ModTime() time.Time {
	return f.modTime
}

//IsDir is an abbreviation for Mode().IsDir()
func (f FileInfo) IsDir() bool {
	return f.mode.IsDir()
}

//Sys is the underlying data source (can return nil)
func (f FileInfo) Sys() interface{} {
	return nil
}
