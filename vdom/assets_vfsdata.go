// Code generated by vfsgen; DO NOT EDIT.

package vdom

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// assets statically implements the virtual filesystem provided to vfsgen.
var assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 4, 5, 1, 3, 41, 228339641, time.UTC),
		},
		"/index.html": &vfsgen۰CompressedFileInfo{
			name:             "index.html",
			modTime:          time.Date(2019, 4, 5, 1, 2, 37, 437648339, time.UTC),
			uncompressedSize: 2550,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\xcd\x8e\xdb\x36\x10\xbe\xef\x53\x4c\x75\x08\x24\xc4\x90\xb6\x69\x0b\x04\xb6\x14\xa0\xd8\x04\x48\xdb\x24\x2d\xba\x5b\xe4\x10\xec\x81\x4b\x8e\x2d\x76\x69\x52\x20\xc7\x72\x8c\xc2\xef\x5e\x90\x92\x6c\x4a\xf2\x36\xcd\xa1\x40\x7d\xb0\xa4\xf9\x9f\x6f\x7e\xc8\xb2\xa6\xad\x7a\x75\x05\x50\xd6\xc8\x84\x7f\x01\x28\x1d\xb7\xb2\x21\xa0\x43\x83\x55\x42\xf8\x99\x8a\x3f\x59\xcb\x3a\x6a\xd2\xc9\x00\x70\xa3\x1d\xc1\x8d\x92\xfc\x11\x2a\xb8\x5e\xf4\x64\x80\xf7\x66\xe7\xf0\xb5\xd9\x6b\xa8\xe0\xdb\x09\xf9\x8f\x06\x2a\x78\x71\x26\xfe\x82\x87\x5e\xf2\xbb\x11\x31\xc8\x7d\x7f\x26\xdd\xd4\x4c\x6f\x10\x2a\xf8\x61\x75\xd5\x13\x5b\x66\xc1\x19\xfe\x88\x74\x22\xad\x77\x9a\x93\x34\x1a\x2c\x6e\xa4\x23\xb4\x6f\x5a\xd4\xf4\x96\x69\xa1\xd0\xa6\xda\x08\x5c\x84\xa4\x16\x20\x45\x06\x7f\x9d\xac\xbb\xbd\x24\x5e\x43\xea\x79\x31\x1d\x80\x33\x87\x5d\x8e\xcb\x88\x0a\xe0\x6d\xe5\x4c\x88\xe0\xe0\x9d\xf7\xa5\xd1\xa6\x09\xf7\x92\xc9\xe2\x14\x48\x8a\x2d\x8d\x0d\xfa\x9f\x42\x02\x61\xb6\x41\x17\xaa\x19\x1b\xe0\xee\xd0\xe0\xb2\x8b\x74\xc6\x7b\xa3\x70\x8b\x9a\x7e\x12\x4b\x90\x62\xce\x7e\xcd\x88\x2d\xa1\x0b\x04\x45\x32\xe1\x1f\x57\x13\x42\x07\x60\xee\x50\x8b\xf4\xe7\xdb\x5f\x3f\xe4\x8e\xac\xd4\x1b\xb9\x3e\xa4\x43\x84\x59\x36\x56\x3a\x4e\xbe\x1f\x2c\xb2\xc7\xd5\x0c\xb3\x50\xb0\x7f\x07\x5a\x10\xfd\x1f\xa0\x86\x2d\xe5\xc4\xec\x06\x29\x6f\x99\xda\xe1\x97\xc0\xf3\x13\x60\x14\xe6\xca\x6c\xce\x68\xfd\xf7\x08\x1f\xaf\x86\xe7\xb4\xed\xb9\x45\x46\xd8\xe7\xfa\x3b\xf2\x9d\x75\xb2\xc5\x14\x3b\x42\x0c\xa9\x07\xd3\x97\x03\x2a\x10\x86\xef\x3c\x3b\x1f\x69\x0f\x4a\xf9\x07\xb6\xc5\xec\x34\x61\x00\x6b\x63\x21\xf5\xb3\xa7\xd9\x16\x41\x6a\x18\x04\x7f\x24\xb2\x6e\x5c\xb6\x50\x70\x87\xe4\x59\xf2\x61\x47\x98\x7a\xa5\xc5\x58\xe5\x93\xa7\xdd\x67\x71\x7e\x73\x67\xd2\xaf\x98\x15\x48\x28\x4f\xca\x37\xb5\x54\xc2\xa2\xce\x15\xea\x0d\xd5\x2b\x90\xcf\x9f\x8f\xdd\xcb\x35\xa4\x53\xe9\x4f\xf2\x3e\xf7\x8d\x02\x55\x05\x2f\xa6\x4d\x16\xe2\x95\x5a\xa3\xbd\xc3\xcf\xbe\xc9\x2e\x69\x77\x41\x27\x5e\x22\xb9\x8f\x0b\x75\x04\x54\x0e\x27\x26\xb9\xd7\x84\xea\x0b\xa5\x89\x1d\x4c\x8a\xdf\x0d\x4d\xd3\xa0\x16\x41\x28\x0d\x16\x47\x42\xc7\xaf\x85\x2e\x5e\x8b\xee\x69\xfc\x2e\x6e\xd1\x59\x70\xe3\x69\xba\xe8\x62\x00\xfd\xb2\x68\x0f\xa8\x14\xc9\x7d\xc4\xbf\xdc\x10\x16\x69\x67\x75\xf0\xbb\x7a\x72\x0c\x58\xd3\xa8\xc3\x6f\x8c\x78\x9d\x36\xfe\x3f\xce\x6a\xa8\xc7\xa9\xeb\x1f\x8c\x38\xe4\x6b\x69\x1d\xf5\xc5\x09\x28\x9f\x9d\xfb\x26\xea\x94\x9e\x3d\x83\x6f\xfa\x57\xa9\x1d\x31\xcd\xd1\xac\xe1\xed\xdd\xfb\x77\xb7\xe1\x64\xec\xf5\xb3\x31\x88\x63\x47\x16\xb7\xa6\xc5\xcb\x85\x8c\xd2\xec\x47\xf3\x89\xa6\x09\x49\xe5\x83\xbb\xb3\x81\xb1\xab\xb8\x67\xbc\xbd\x6c\x06\xd8\xde\x41\x05\x52\x4b\xfa\x78\x9b\x66\xf3\x43\x74\xe0\xc4\x07\x65\x58\x66\x50\x81\xc6\x3d\x7c\xc4\x87\xdb\xf0\x9d\x26\x7b\xb7\x2c\x0a\x65\x38\x53\xb5\x71\xb4\x7c\x79\xfd\xf2\xba\xe0\x4a\xa2\xa6\x24\x8a\xaf\x5f\x85\x46\x9b\x06\xfd\x91\x7f\xda\xf7\x93\x33\x37\x5a\xab\x49\xef\xd1\x6b\xc4\xa6\x8e\x17\xac\x6e\xd1\x39\x16\x2e\x09\xe7\x83\x64\x6c\x39\x00\x07\x15\x84\x3d\xdc\x30\xeb\x30\xc5\x5c\x30\x62\xa3\x71\x8a\xfd\x77\xfd\x13\x73\x67\xcd\xf5\xcf\x41\x71\x65\x1c\x7e\x5d\xae\x41\xe5\x89\x64\xfb\xfe\x1f\xee\x3d\x43\x41\xfd\x7f\x59\x74\x17\xb4\x70\x99\x2b\x86\xdb\x5c\xe9\x7b\xe1\x55\x59\x84\xc7\x55\x59\x74\xd7\xbd\xbf\x03\x00\x00\xff\xff\xd5\x70\xc3\xae\xf6\x09\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/index.html"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}