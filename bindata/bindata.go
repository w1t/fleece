// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package bindata generated by go-bindata.// sources:
// docker/go-fuzz.dockerfile
// docker/go-fuzz.sh
// fuzzing/common.go
// fuzzing/utils.go
package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dockerGoFuzzDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\xcd\x6a\xeb\x30\x10\x85\xf7\x7a\x8a\x03\x77\x2d\x1b\xc3\x0d\x85\x40\x37\x4d\xda\x12\xda\xc4\xc5\x10\x4a\x97\xb2\xf5\x8b\x65\x8d\x88\xa4\x94\xe6\xe9\x8b\x4d\x03\xe9\xae\x2b\x89\x99\x39\x1f\xdf\x79\xea\xda\x3d\x0c\x79\x11\xcc\xba\xa9\x9a\xff\x5c\xf8\xe8\x82\x62\xec\x1f\x76\x21\x65\xe1\x3d\xa4\x8a\x09\x22\x48\x94\xa4\x74\xf1\xc8\xd6\x05\x93\x58\x77\x3c\x40\xc4\x11\x42\x4a\xf4\xc5\x79\xc9\x7b\x91\x14\xe2\x68\x06\x0a\xda\x19\xf4\x22\x59\x9c\xdd\x04\xe3\xf2\x2d\xcf\x10\xd7\xe5\x72\x59\x00\xcf\x6d\xd3\x34\xfb\x76\x7b\x7c\x7d\xbc\x27\xad\x61\x08\x46\x65\xf0\x32\x87\x6c\xe9\xab\x81\xa6\x5a\x9e\xbf\xca\x48\xe7\xfa\x27\x78\x7d\xff\x70\xc2\x17\x31\xf6\xde\x76\x2f\xdb\x5d\x87\x3a\x4f\xb1\x9e\xe7\x2e\x98\xd9\xe8\xf0\xb0\x86\xa6\xd3\x5c\x0c\x9f\x2e\x5b\x48\x1a\x46\x75\xe2\x03\x4d\x91\x92\x62\x9b\xf6\xed\x03\xd5\x15\x56\x25\x8b\x9b\xff\xe2\x3f\xd8\x89\x24\xee\x56\xab\x5f\x1b\xb6\xd9\x6f\x97\xfa\xec\x3b\x00\x00\xff\xff\x7b\x8f\x6d\x78\x5f\x01\x00\x00")

func dockerGoFuzzDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_dockerGoFuzzDockerfile,
		"docker/go-fuzz.dockerfile",
	)
}

func dockerGoFuzzDockerfile() (*asset, error) {
	bytes, err := dockerGoFuzzDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docker/go-fuzz.dockerfile", size: 351, mode: os.FileMode(420), modTime: time.Unix(1600679318, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dockerGoFuzzSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x52\xc1\x6e\xdb\x3a\x10\xbc\xef\x57\x4c\x6c\x21\x88\x1f\x1e\xe5\x26\xc7\xb8\x0a\xda\xa2\xe8\x35\x97\xde\x0c\x23\xa6\xa4\x95\x44\x98\x21\x05\x92\xaa\x91\x38\xfe\xf7\x82\x94\x6c\x23\xee\x49\xd4\x60\x77\x67\x76\x76\xe6\x37\x58\x0e\xde\x2d\x4b\x65\x96\x6c\xfe\xa0\x94\xbe\x23\xcf\x01\x82\x89\xe6\xf8\xfd\xfc\xf3\xf9\x11\x8e\x7b\x2d\x2b\x46\xe8\x94\x87\xef\x58\x6b\xf8\xca\xa9\x3e\xd0\x1c\xd9\x3d\x04\x9a\xe1\xfd\x9d\x1d\x7a\x59\xed\x64\xcb\x11\x7d\x80\x80\x91\xaf\x0c\xdb\x20\x74\x7c\xaa\x68\x06\x53\x05\x65\x0d\xee\xf6\x4a\x6b\x94\x8c\xc1\x73\x8d\xc6\xba\x54\xb5\xdd\x5b\xb7\xab\x95\xdb\x42\xba\x76\x78\x65\x13\x10\x2c\xb6\xad\x15\xb1\x7f\xbb\xc8\x69\x8e\xad\x10\xe5\xa0\x74\xbd\x85\xc0\x9d\xed\xe3\x34\xa9\x17\x48\x98\xbf\x52\x82\x92\x1b\xeb\xe2\x47\x99\x16\x6e\x30\x71\xc0\xf7\xba\x56\x63\x57\x64\xf1\x68\xac\xd6\x76\x1f\x0b\xb6\x42\x44\x66\x46\x2f\x7d\x94\x55\x2b\xc7\x55\xd0\x6f\x9f\x44\xe4\x44\xaa\xc1\x7a\x8d\x6c\x0e\xa1\x03\x1e\xb0\xd9\xac\xa2\x7a\x43\x00\x57\x9d\xc5\x6c\xf0\xb2\xe5\x47\xb0\x09\xee\xad\xb7\xca\x84\xdc\x77\xf8\x7a\x25\xad\x97\xa1\x7b\x3a\xa3\xd1\x98\x64\xd8\x13\xd6\xa2\xfc\x98\x76\xdc\x60\x2d\x04\xd6\x13\x75\x94\xbb\xfe\x1f\x79\x9e\x6f\x36\x9b\x19\x35\x8a\xa8\xdf\xb5\x45\x76\x4f\xbe\x53\x4d\xa0\xd8\x7e\xf9\x3b\x1d\x6f\x30\xb5\x1d\x2f\xd7\xc9\x6a\x97\xd3\x64\x71\x91\x2f\x1b\xcd\x5c\xf1\x72\x02\xfc\x32\x3b\xc4\x09\x47\x2a\x95\x29\xb2\x43\xbf\x6b\x1f\x1f\x8e\x89\x38\x7f\x57\x7d\x44\x5f\xa2\xe6\x22\x3b\x4c\x1d\xc7\x65\x76\x28\x95\x39\x12\x75\xd2\xbf\x44\xc1\xa1\x68\xa4\xf6\x4c\x49\xfc\xdd\x02\x07\x02\x5e\x63\x29\x44\x8f\x4b\x1f\x01\xd3\x4a\xe3\x9a\x10\x16\x69\x54\x22\x38\x22\x91\x1f\xe9\x48\xb4\xef\x94\xe6\x93\xd9\x6d\xc0\x97\x64\x76\x6d\x09\xa8\xa4\xe7\x98\x3e\x15\x7d\x17\x25\x3e\x30\x99\xb6\x20\xe0\x74\x89\x1f\x11\x88\xa7\x4d\xfe\x95\xca\x48\xf7\x96\xe7\xf9\x2c\x95\xa4\xea\xf4\xba\xe8\x0f\x6e\xe0\x04\x8d\x26\xc6\xd7\x6a\x15\x09\xc4\xe2\x0a\x76\xec\xc3\x4b\x0c\x50\x91\x7d\x1b\xc7\x39\x96\xbb\x4b\xc7\x7f\xd7\x0d\x09\x65\x2f\x2b\xaa\xad\xe1\x73\x8a\xce\xdc\x28\x0a\xcc\x92\x7f\x33\xdc\xde\xe2\x06\x82\x91\x9d\x4c\xf9\x37\x64\xbf\x2e\x1b\xc1\xd8\x80\xc6\x0e\xa6\x5e\x8d\x4b\x29\xd3\x4e\x5b\x8e\x3b\xc6\xa8\x7c\xd2\x4b\xa7\x48\x89\xf1\xd8\x17\xef\x45\x4c\x62\x31\x45\x01\xe2\x94\x96\xcb\xf1\x90\x9d\x27\xd1\xdf\x00\x00\x00\xff\xff\x9d\x50\xca\x5a\x3e\x04\x00\x00")

func dockerGoFuzzShBytes() ([]byte, error) {
	return bindataRead(
		_dockerGoFuzzSh,
		"docker/go-fuzz.sh",
	)
}

func dockerGoFuzzSh() (*asset, error) {
	bytes, err := dockerGoFuzzShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docker/go-fuzz.sh", size: 1086, mode: os.FileMode(493), modTime: time.Unix(1600695326, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fuzzingCommonGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x57\xdd\x6f\xdc\xc6\x11\x7f\xe6\xfe\x15\x13\x02\x91\xc9\xf8\xc4\x6b\xfa\x78\xed\xbd\xd4\x8e\x02\x15\xb5\x6c\xa4\x06\xf2\x60\x1b\xcd\x1e\x39\xbc\x5b\x88\xb7\x7b\xd9\x1d\x4a\x91\x1c\xfd\xef\xc5\xec\x07\xbf\x4e\x6a\x8a\xf6\x00\x09\xdc\xe5\xec\x7c\xfc\x76\xe6\x37\xc3\xf5\xfa\xf5\xae\x57\x5d\x03\x7b\xd3\xf6\x8f\x8f\x42\x9c\x64\x7d\x2b\xf7\x08\xbc\x52\x7a\x2f\x84\x3a\x9e\x8c\x25\x28\x44\x96\xb7\x47\xca\x45\x96\x2b\xb3\x56\xa6\x27\xd5\xf1\xc2\x38\xfe\x7f\x92\x74\x58\xb7\xaa\x43\x7e\xe0\x0d\x8b\x6d\x87\xb5\x17\x27\x74\xa4\xf4\x3e\x17\x22\xcb\xf7\x8a\x0e\xfd\xae\xaa\xcd\x71\xed\xc8\x22\xd5\x07\xbb\xf6\xef\xdb\x87\xb5\xc5\x5f\x7b\x65\x31\x17\xa5\x10\xeb\x35\x5c\xf5\x8f\x8f\x50\x1b\xed\x48\x6a\x72\xd0\x1a\x0b\x7b\x73\xc9\x6e\x01\x19\xe8\x1d\xc2\xfd\x01\x35\x58\xa4\xde\x6a\xa5\xf7\xd0\x5a\x73\x04\x3a\x60\x38\xd9\xf6\xba\x16\x77\xd2\xb2\xe3\xbc\x71\xad\x09\x6d\xf0\x04\xb6\xf0\x7d\xd8\xbc\x31\xf6\x28\x3b\xf0\xbf\x2d\xfc\x29\x6c\xbe\x55\xae\x96\xb6\x89\x9b\x97\xdf\x47\x7f\xde\x58\xe9\x0e\x68\xc1\xe2\xc9\xa2\x43\xf6\x49\x0e\x1e\xe5\x75\x78\x9b\x43\x21\x35\x28\x7d\xea\x09\xe8\x20\x09\xc2\x7e\xe3\xfd\xb2\xe8\x4e\x58\x93\xba\x43\xb1\x5e\x67\x6d\xf2\x92\x94\xd1\xe5\x0a\x14\x39\xa8\x8d\x65\x21\xa3\x1b\x76\xd3\xf4\xc4\x7a\x8a\x93\xd4\xaa\x86\x23\x3a\x27\xf7\x58\xae\x40\xea\x06\xb4\x3c\x22\x14\xc1\xd0\x41\xba\x43\x59\x09\x7a\x38\xe1\xe0\xa5\x23\xdb\xd7\x04\x5f\x45\x76\xc3\x92\xfc\x73\x64\xf9\x42\xb3\x6b\x7f\x08\x00\x3e\x7d\xd9\x3d\x10\x8a\xec\x7d\xb0\x33\x4a\x30\x0a\x57\xbd\xae\x21\x3d\x88\xa7\x29\x04\xd7\x84\x56\x92\xb1\xa0\x1c\x70\xb4\x69\x39\xbd\xa2\x04\x88\xcb\xa1\x33\xb5\x24\x6c\x40\x69\x46\x81\x43\x1f\x81\x80\x19\x0a\xaf\x1c\xdc\x1b\x7b\xcb\xb1\x37\xca\x62\x4d\xc6\x3e\xcc\xe3\x1a\x4c\x8f\xf1\x29\x18\x7e\x4a\x93\xc8\x94\x6e\x8d\x0b\xeb\x4f\x5f\x8c\xab\xae\x54\x87\xd7\xba\x35\xc2\x23\xee\xc3\x82\x31\xb0\x2c\xfa\xf9\x56\xd9\x14\xfe\x93\x08\x26\x07\x14\xd8\xbd\x22\x80\x55\x7a\x1b\xfe\xf5\x4f\x58\x9b\x3b\xb4\x6f\x64\xd7\xed\x64\x7d\x1b\xa4\xfc\x55\xbd\x73\xfb\xa8\x2b\x64\xce\x0d\xde\x2f\xdc\xbf\x32\x36\x26\xee\x39\x82\x09\xb8\x90\x3f\x9d\x7c\x54\xdd\x03\x74\x46\x36\x2e\xa3\x03\x2a\x1b\xb2\xcb\xf9\x2c\x08\x19\xe2\x2a\xc1\xc6\x9f\xb1\x53\xb4\xcb\x9b\x2c\xa1\xf8\x6e\x21\xb4\x02\xb4\xd6\xd8\x92\xc1\xf4\x69\xb5\xd9\xc2\x8f\x48\x2c\xcd\xb9\x33\xe8\x28\x45\xc6\xb7\xd3\x28\x1b\x25\x7e\x0e\xab\x82\x0f\x95\x33\x24\xc3\xfb\x37\xc3\xc6\x42\x84\xaf\xc3\x79\xb3\x2c\x19\xb8\xa4\xfa\x09\x65\xc3\xa2\xa3\x9a\x52\x64\xaa\xf5\x52\xdf\x6c\x41\xab\x8e\x1d\xcc\x02\x6c\xbc\xf4\x0a\x44\xf6\x24\xd2\xde\xc5\x22\x30\x16\xf7\xd9\xb0\x09\xe9\x30\xb3\x2e\xb2\x89\xc3\x1b\x48\xdc\x55\xfd\xdd\x28\x5d\xc4\x38\x57\x93\x34\x2e\xf9\x44\x82\x82\x15\xa6\xe7\x95\xc8\x9e\x56\xec\x4f\x2c\x92\x77\xbd\xa3\xf3\x9b\xf8\xe3\xeb\xde\xf5\x04\x3e\x7d\x1c\xa8\x96\xc5\xfc\xad\x80\xa9\xeb\xde\xa6\x1b\x7e\x5e\xf7\x73\xb7\xbc\xbc\x64\x5f\x29\x84\x76\x40\xfd\x3f\xa8\x79\x16\x77\xef\x5a\x81\x96\x6f\x65\x84\x9c\x55\xc6\xb8\x63\x3d\x30\x29\x28\x4d\xa8\x1b\x66\x3d\x03\x3b\x84\x06\x5b\xb4\x16\x9b\x0a\xae\x09\x6a\xd9\x75\x2e\xd2\x61\x38\x50\xa7\x0a\xba\x57\x74\x48\x14\x11\xea\x67\x24\x5b\xc9\xfc\x00\xa6\x9d\x9e\xc4\x06\xee\x64\xd7\x63\x64\x16\xc0\x3b\xd4\xc4\x32\x32\x00\x19\x41\x4b\x19\x35\x60\x52\x26\x5f\x8b\xa8\xe8\xcd\x6e\x59\xcd\xbe\x18\x54\x0b\x1e\xaa\x28\x55\x94\x7f\x81\x45\x2a\xc6\xd3\x45\x7b\xa4\xea\x9f\x27\xab\x34\x15\xb6\xf4\x00\x05\x4c\x3e\xa2\x9b\x46\x3c\xdc\xf5\x8c\xf5\x3c\xf3\x07\x26\xe7\x9a\x8e\x5a\x5d\xe8\x67\x21\x23\x18\x91\x84\x0e\x9c\xa4\x73\xd8\x9c\xc1\xf7\x72\xb4\xec\xc5\x1f\x84\xea\xaf\x28\xf9\x57\x9d\xc1\x33\x96\x6e\x95\x52\x2c\xd9\xa9\x7c\x3b\x29\x63\xc0\x37\xf8\x1b\xc1\x1e\x29\xc4\xab\x79\x95\xfc\xe1\xd8\x94\xae\x2d\x1e\x7d\xef\xe4\xf7\xa9\x10\x92\xeb\xbc\x3e\xcb\xdc\xd2\x2b\x2d\x4a\x28\xbc\xba\xf4\x7a\x05\x8d\xd1\x08\x3b\x63\x02\x11\x4c\x48\x8c\xfd\xf8\xdb\xc6\x6b\xaf\x42\x37\xa8\x8d\x26\xa9\x74\x84\x79\x15\x79\x33\x74\xd2\x5f\x7b\x13\xba\x13\x5f\x00\xd3\x80\xab\x58\x45\xf6\x33\xc2\xbd\xd4\xc4\x39\xec\x6e\xd5\x09\xb4\xd1\x97\x73\x21\x2e\xdf\x6f\xbc\x17\x89\x69\x3c\x9d\x0d\x66\x3f\x85\xc7\x2f\x22\x1b\x78\x95\xf7\x2b\xcf\xaa\x25\x1f\x69\xc3\xc6\xb5\x63\xe2\x2b\xe1\xf7\xdf\x47\x1e\xfa\xe1\x37\x0a\xac\xc9\x19\x97\xe7\xde\x44\x16\x14\xbe\x7e\xcd\xcf\xde\x70\x32\x07\xdb\x2d\x74\xa8\x8b\xd1\x78\x79\xf9\x67\x96\xe2\xc8\x95\xee\x51\x64\xd9\x93\xf0\x5e\xfa\xc0\xcf\xa9\x97\xdb\x64\x31\x67\x41\xaf\x6c\x24\xc9\x95\x1f\x3a\xca\xe8\xf8\x82\x1e\x66\xbc\xcc\xae\x45\x76\x0e\x56\x13\xde\xff\x87\x59\x78\x0d\x79\x15\xf4\xe4\xff\x8b\x0b\x3e\x79\xb6\x43\x8f\xf0\x07\xf8\x22\x42\x6f\x60\x13\x4c\xf0\x61\x38\xda\x84\x61\x82\x5d\xe6\xbd\x30\x1f\x6d\x86\x01\xa9\x08\x6e\xf8\x8e\x30\xcc\x4a\x31\xdf\x26\x5d\x21\x02\x3e\x5e\xd9\x7f\x75\x63\x3b\x8b\xf2\x76\x4a\xb1\xec\x78\x8a\x67\xec\x32\x5c\xd3\x57\x52\x75\x4a\xef\xff\xa1\x8e\x8a\x80\xa7\x68\x07\x28\xeb\x43\xaa\xb8\x57\x03\xab\xec\x39\xf7\xc9\xf3\xcc\x62\xf4\x0a\x83\x68\xe4\xa1\x5e\x93\xea\x40\x11\x38\x44\x07\xbf\x74\xac\xf7\x17\x68\x83\x95\x38\x76\x4c\x2b\xf5\xac\x50\x97\x4e\x15\x04\xdf\xc5\xe9\xbf\xfa\xb8\x02\xaf\x90\x9b\x43\x09\xc5\xbf\x26\x85\x9c\xba\x9e\xa6\x15\x90\x21\xd9\x05\x99\xaf\xe3\xbc\x30\x9d\x54\x5e\x68\x5e\x33\xf4\x4b\x46\xcf\x7f\x4d\x54\x37\xe6\x07\x26\x86\x22\x64\x5f\x29\x3c\x3d\x7c\x7c\xff\xf6\xfd\x06\x4e\xd2\xca\xae\xc3\x4e\x3d\xa2\xc8\xf8\x53\x21\x80\xdc\xa8\xe6\x83\x9f\xb8\x99\x5b\xc2\x8b\x56\x59\x47\x83\xbf\x4b\x86\x9d\x32\xc1\xc5\x45\x0a\xe7\xaf\x31\xde\xaf\xe3\xa4\x31\x49\x4b\xd8\xc2\x22\xb8\x2a\xd0\x9c\xef\x2b\x2f\xb9\x9e\x0d\xbe\x6d\xa1\x95\x9d\xc3\x51\x77\xe5\x59\xfe\xd9\x21\x34\x54\xc7\xe4\x28\x59\x4f\x08\x5c\x43\xd3\xc8\x38\x25\xb9\x98\x2e\x2e\x86\x10\xa7\xe5\x95\xcd\x65\x93\x0c\xbf\x7a\xe2\x6c\x8f\x55\x39\xd8\xf1\x87\x02\x18\x3e\xfd\x59\xc8\x5f\x2f\xaf\xb8\x36\xb8\x6b\x7e\xe0\xa6\xd9\x16\xf9\xf0\xdd\xd2\x1f\x8f\xd2\x3e\x6c\x3e\xeb\xed\xfc\xf7\x59\x73\xd9\x4f\x8f\x5c\xfa\x4e\xa8\xf4\x7e\x03\xdf\x36\x9f\x75\x1e\x93\xe7\x32\x98\x3c\x13\x8e\x79\x3c\x08\xbf\x20\xe6\x75\xcc\x35\x32\xf2\x4b\xa8\x26\xc0\x0c\xe7\x3b\x5d\xe4\x9e\x9b\xa6\x1a\x7d\x47\xf4\xb6\xfc\x77\x4d\x04\x6d\x03\xdf\xba\xcf\x9a\xff\xf2\xd5\x4c\xb1\x6f\x0e\x8b\xad\xc0\x3f\xb3\xb9\x6b\x9e\x90\x21\x94\xe8\x2d\x13\x84\x2f\xd2\xc5\x84\x3e\x64\x43\x1c\xad\xbe\x8a\xec\x74\xbb\xff\x20\xe9\x10\x66\x1c\xff\xbd\x5e\x7d\x7c\x38\xe1\xfb\xb6\x48\x2c\xf9\x54\x56\x1f\x82\x10\x67\xe6\xd1\x34\xfc\x38\x94\xe1\x8f\x48\xef\xc2\x56\x11\x55\x3d\x3b\x3a\xc6\x7e\x7c\xfd\xea\x08\x27\x8b\x44\x0f\xe0\x7a\x8b\x40\x07\xe5\xc0\x1d\x4c\xdf\x35\xfa\x15\xf1\xa0\x78\x32\xce\xa9\x5d\x87\x67\xe3\xe6\x24\xee\x69\xa7\x18\xfc\xc9\xdb\x0e\xb1\xc6\x7c\x05\x79\x1c\xdd\x5d\x1e\xfb\xc6\x14\x8e\xc5\x07\xc9\x33\x88\x3c\x6b\x66\xf9\xa9\x33\xfb\x2c\x10\x4f\xe2\xdf\x01\x00\x00\xff\xff\x83\xd6\xab\x95\x44\x11\x00\x00")

func fuzzingCommonGoBytes() ([]byte, error) {
	return bindataRead(
		_fuzzingCommonGo,
		"fuzzing/common.go",
	)
}

func fuzzingCommonGo() (*asset, error) {
	bytes, err := fuzzingCommonGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fuzzing/common.go", size: 4420, mode: os.FileMode(420), modTime: time.Unix(1600694133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fuzzingUtilsGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x94\xcf\x6e\xdb\x38\x10\xc6\xcf\xe2\x53\xcc\x12\xc8\x42\xc4\x3a\x72\x72\x0d\xe0\x3d\x34\x4d\x7a\x69\x52\x03\x05\x7a\x31\x7c\x60\xc4\x21\x45\x98\x22\x55\x8a\xca\x3f\xc7\xef\x5e\x0c\x25\x39\x09\x90\x00\x3d\x89\x22\x67\xe6\xfb\x7d\x23\x8e\x96\xcb\xff\xee\x06\xeb\x14\x98\xa0\x87\xe7\x67\xc6\x3a\x59\xef\xa4\x41\xa0\x37\xeb\x0d\x63\xb6\xed\x42\x4c\x50\xb2\x82\x63\x8c\x21\xf6\x9c\x15\x5c\xb7\x89\x1e\x9d\x4c\xcd\x52\x5b\x87\xb4\xa0\x8d\x88\xda\x61\x9d\xcf\xe2\xe0\x93\x6d\x91\x33\x56\x70\x13\x9c\xf4\xa6\x0a\xd1\x2c\x1f\x97\x29\x04\xd7\x2f\x4d\x58\x4e\x52\x3d\x67\x82\xb1\xf4\xd4\x8d\xa2\xd7\x83\xaf\x41\x0f\xbe\x2e\x37\xdb\xbb\xa7\x84\x02\xac\x4f\x8c\xd1\x0e\x7c\xc3\x74\x13\xd4\x5a\xa6\xa6\x1c\xb1\x68\x09\x7d\x8a\xd6\x1b\x01\xe5\xb8\x58\x40\xe6\x14\xb0\x67\x45\x1b\x14\xc2\xc5\x0a\x66\xa9\xea\x16\x51\xdd\xca\x16\xe1\xe5\xfd\xde\x4d\x50\x83\x43\x56\x74\x3b\x73\x19\xbc\xb6\x86\xb2\xfe\x3d\x86\x8c\x7b\xfb\x9b\xa0\xf0\x02\xa8\xe8\xe1\x4d\x68\xf5\x85\x3a\x78\xed\xa4\xe9\x61\x05\x9b\xed\x88\xb1\xe7\xa7\xa7\x49\x9a\x9e\x2f\x80\x8f\xcd\xe5\x63\x52\x9f\x01\xdf\x51\x7d\x0f\x52\x95\xc7\x7a\x0b\x78\x35\x27\x58\x61\x75\x8e\xff\x67\x05\xde\x3a\x32\x55\x44\x4c\x43\xf4\xc0\x79\xae\xc4\x8a\x43\x0e\x72\xe8\xa9\x46\x2f\xe0\x7f\x38\xff\x20\x2e\x44\xb2\xfa\x50\xea\x36\x55\x3f\xbb\x68\x7d\xd2\x25\x2b\x8a\x82\xb7\x21\x22\xa4\x46\x7a\x08\x1e\x67\x28\x68\x65\xaa\x1b\x54\xd0\xc9\x94\x30\x7a\x38\xf9\xcd\x17\x14\xfe\xca\x46\xaf\x42\x90\x7c\xf6\x95\x1d\xed\x4c\xbf\x39\xdb\x66\x9e\x6e\x67\xaa\xab\x2c\x3b\xb3\xbf\xbc\xcc\x90\xd3\x81\xa0\x93\xb3\xcc\x8a\x31\x5e\x86\xc1\xa7\x77\x7d\x59\x13\xe5\x18\x3a\x5a\xfb\x7b\x53\x27\x6a\x3a\x06\xf4\x35\x15\xc6\x88\x0a\x1e\x1a\xeb\x10\x5c\x90\xca\x7a\x73\xb4\xfa\x60\x53\x33\xf5\x9c\xec\x36\x47\xaf\x33\xd4\xdb\x2f\xf2\xd6\x75\x1b\xd4\xe4\xba\x9a\x2f\x90\xd5\x74\x3f\x60\xf5\xd9\xd7\xfa\x14\x78\xf0\xf2\xce\x21\xa4\x90\xf1\xa8\xc8\xe0\x10\x74\x88\x47\x4c\x99\x3e\x84\xfc\x18\x6d\x52\x6d\x83\xaa\xbe\xda\xb8\x20\x1a\x76\x78\x1d\x23\x9a\x32\x9a\x83\x52\xd3\x78\x61\xd4\xb2\xc6\xfd\x41\x4c\xb3\x44\xdc\xf7\xd2\x91\xb7\x69\xa2\xab\x5f\xd2\x0d\xf8\x43\x97\x5a\xb0\x42\x2a\x95\xef\xef\xbd\x74\xd5\x3a\xe4\xfc\x52\x1c\x25\xe7\xff\x41\x75\xf5\x98\xca\xe9\x2f\x50\x91\xde\x75\x88\xeb\xcb\x92\x92\x45\x95\xb5\x85\xd8\x9c\x5f\x6c\xd9\x81\xfd\x09\x00\x00\xff\xff\x25\x2f\xb6\xbd\x86\x04\x00\x00")

func fuzzingUtilsGoBytes() ([]byte, error) {
	return bindataRead(
		_fuzzingUtilsGo,
		"fuzzing/utils.go",
	)
}

func fuzzingUtilsGo() (*asset, error) {
	bytes, err := fuzzingUtilsGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fuzzing/utils.go", size: 1158, mode: os.FileMode(420), modTime: time.Unix(1600696615, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"docker/go-fuzz.dockerfile": dockerGoFuzzDockerfile,
	"docker/go-fuzz.sh":         dockerGoFuzzSh,
	"fuzzing/common.go":         fuzzingCommonGo,
	"fuzzing/utils.go":          fuzzingUtilsGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"docker": &bintree{nil, map[string]*bintree{
		"go-fuzz.dockerfile": &bintree{dockerGoFuzzDockerfile, map[string]*bintree{}},
		"go-fuzz.sh":         &bintree{dockerGoFuzzSh, map[string]*bintree{}},
	}},
	"fuzzing": &bintree{nil, map[string]*bintree{
		"common.go": &bintree{fuzzingCommonGo, map[string]*bintree{}},
		"utils.go":  &bintree{fuzzingUtilsGo, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}