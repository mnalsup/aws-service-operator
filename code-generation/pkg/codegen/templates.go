// Code generated by go-bindata.
// sources:
// assets/cft.go.templ
// assets/controller.go.templ
// assets/template_functions.go.templ
// assets/types.go.templ
// DO NOT EDIT!

package codegen

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _cftGoTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6d\x6f\x9b\xc8\x13\x7f\x0d\x9f\x62\xfe\x56\xd5\x3f\x44\x0e\xbe\xbe\xb5\x2e\x27\xf5\x9a\x3e\x44\xad\x92\x2a\x76\xdb\x17\x51\xd4\xae\x61\x8c\xb7\xc1\x2c\xc7\x2e\xf1\xf9\x2c\xbe\xfb\x69\x1f\x30\x0b\x36\x76\xdc\xf6\xda\xeb\xc9\xa8\x95\x61\x99\xf9\xcd\xe3\xce\xec\x84\xc1\x00\x7e\xd3\x17\x9c\x5f\xc1\xe5\xd5\x18\x9e\x9f\x5f\x8c\x61\xfc\xea\x62\x04\x2f\x2e\xde\x3c\x87\x5f\xd7\x97\x3b\x18\xc0\x78\x46\x39\x4c\x69\x82\x40\x39\x90\x42\xb0\x18\x53\xcc\x89\xc0\x08\xee\x29\x81\x4f\x64\xc1\x4f\x59\x26\x57\x58\x7e\x1a\xb2\x08\x63\x4c\x21\xcb\x59\x88\x9c\x7f\x92\x00\x17\x53\x58\xb2\xe2\xff\x11\x24\xf4\x0e\x41\xcc\x10\xc2\x19\x49\x63\x04\x92\x2e\xc5\x8c\xa6\x31\x90\x09\x2b\x04\x88\xb5\xa0\x39\xb9\x43\xc0\x88\x0a\x0e\x82\x29\x8e\x40\xe0\x3c\x4b\x24\x9a\xd6\x24\x55\xab\xd9\x5d\x3c\x30\x12\x07\x84\x73\x14\x1c\x22\x9a\x63\x28\x58\xbe\x0c\x5c\x37\x23\xe1\x1d\x89\x11\x56\xab\x60\x94\x61\x18\x5c\x23\x67\x45\x1e\x62\x70\x49\xe6\x58\x96\xae\x4b\xe7\x19\xcb\x05\x78\xae\xd3\xc3\x3c\x67\x39\xef\xb9\x4e\x2f\xa6\x62\x56\x4c\x82\x90\xcd\x07\x64\xc1\xe5\xff\x53\x1e\xdd\x9d\xc6\x4c\xde\xee\x26\xe0\x98\xdf\xd3\x10\x07\x61\xc2\x8a\x68\xca\xf2\x39\x11\x94\xa5\x3d\xd7\x21\x0b\xfe\xfe\x09\x49\xb2\x19\x79\x02\x2d\xfe\x84\x4c\x0c\x86\x66\x5e\xfb\x72\x20\xad\x23\x19\xe5\x83\xf6\x9b\x40\x8a\xbd\x37\x78\x9b\x1a\xed\x46\x0c\x59\x3a\xa5\xf1\xa1\x5c\x33\x4c\x32\x94\xfe\xf1\x5d\x19\x84\x4b\x5c\x40\x95\x06\x1c\x08\xa4\xb8\x00\x36\xf9\x8c\xa1\x70\xa7\x45\x1a\xca\xf7\x9e\x16\x04\x27\xfa\x37\x78\xa6\x7e\xfa\x5d\xc1\x80\x13\xcb\x47\x41\x45\xf4\x9a\xa6\x51\x59\xf6\x41\xb0\x8c\x86\x4f\xaf\x2f\x81\x8b\x9c\xa6\xb1\x0f\x27\xcf\x1a\x2e\x86\x95\xeb\xe4\x28\x8a\x3c\x85\xc7\xcd\x37\x2b\xd7\x71\x9a\x68\xc3\x2e\x15\xfa\xae\xe3\x68\x65\x87\x8e\xbc\xf4\x7d\xdf\x05\x80\xb5\x02\x43\xd0\x57\xf5\xdc\x77\x9d\xd2\x2d\x95\x4f\x5a\x1a\x45\x38\xa5\x29\x72\x95\xa7\x5d\x36\x87\x53\xc1\x5d\xb1\xcc\xb0\xcd\xcc\x45\x5e\x84\x42\x5a\x65\xbc\x58\x5d\x4d\x6f\xba\x2d\xd3\x76\x39\xd1\xad\x95\x36\x58\xda\x97\x46\xfb\x91\x20\xe1\x9d\x54\x0b\xb4\x1f\xb5\xe2\xa9\x5c\x60\x53\x75\xcf\x25\x05\x4c\x08\xc7\x08\x98\xde\x7f\xad\xad\xaf\x54\x52\xf1\xf7\x78\x3b\x42\x7e\x2d\xc1\xf3\x8d\x68\x2b\x6a\x26\xbf\x82\x9a\x88\x07\x95\xa9\x49\xc1\x05\xe6\x72\xb1\x0f\xbd\x0e\x57\xf6\xfa\xc0\x5b\x06\x07\x9a\x63\xeb\x32\xcf\x48\x88\xbe\x31\xfd\x25\x8a\xab\x42\x64\x85\xe0\xc6\x76\xcb\x5c\x66\x5e\x4c\x73\x36\x57\xcb\xe7\xc8\xc3\x9c\x4e\x50\x29\xca\x21\x24\x49\xd2\x6d\x72\x8d\xec\xf9\xe0\xcd\x49\x76\xa3\x0d\xbf\xd5\x3f\x7d\x50\x55\xc7\x97\x7e\xa8\x04\x0d\xcf\x60\x83\x6e\x55\xba\x0e\x47\xae\x5e\xae\xdd\xf2\xf4\xc3\x68\x84\x9c\x53\x96\xba\x0e\xbf\x0f\xe5\xbb\x66\xd9\x09\xe4\x26\x94\x6c\xbe\xeb\x3a\xca\x9a\x8b\xb4\x12\xd1\xa2\x6c\x1a\xa5\xc8\xe4\xbe\x59\x07\x43\x66\x3d\x59\xc8\xe8\x48\x75\x3c\x3b\x4c\xbe\xaf\xb6\x40\x65\x80\x32\x49\xe9\x79\x1f\xb6\x60\xbd\xc7\x96\x12\xbe\xeb\xd0\xa9\xa2\xfd\xdf\x19\xa4\x34\x91\x2e\xa8\x72\x21\xa5\x89\x82\x91\xb8\x8e\x2c\x36\x4c\x00\x2f\x72\x04\x3a\xd5\x3d\x82\x72\xc0\x7b\xd9\x64\x18\xe7\x74\x92\xa0\xc2\x4a\x30\xf5\xb4\x0e\x5a\x3b\xee\x4b\xe8\x27\xdb\x80\x59\xce\x95\x73\x7a\x29\xd3\x61\xae\x02\x8f\x11\x2c\xa8\x98\x81\x98\x11\x61\x12\x40\xee\x80\x9e\xaf\x4d\x9c\xb2\x1c\x3e\xf6\x65\x4e\x48\x0b\x73\xd5\xc3\x1a\x22\x6f\x7e\xb9\x0d\xaa\x54\x92\x72\x4d\x50\x6f\x4e\x58\x21\xcc\x8b\xd7\xb8\xbc\x85\x33\xb0\x56\xde\x93\xa4\x40\x2d\xc0\xe8\x69\xd8\xb4\x13\x4c\x71\xc9\x91\x08\xed\x48\x58\xd0\x24\x81\x50\x2d\x58\x99\x6a\x14\x47\xe0\x45\x96\x25\x14\x23\xc8\x48\x4e\xe6\xbc\x3b\x3b\x2d\x4c\x99\x9e\x5a\x2a\x9c\xb4\x72\xc3\xa2\xba\xb2\x42\x5c\x67\xee\x57\x66\x66\x38\x55\xad\x5d\x1a\x33\x3c\x5b\xd7\x81\x97\x28\x94\xba\x2f\x2a\x9e\xb1\x21\x5a\xd7\x85\xc3\x6a\x41\xf5\x38\x5e\x66\x58\x96\xc1\x76\xec\x8e\x82\xf1\x60\x5e\x53\x55\xf6\xed\x36\xcb\xa3\x87\x6f\x35\xa7\x12\xf8\xee\xfa\xcd\xd0\xa6\xac\xfd\xa8\xc8\x2e\x99\xa0\x53\x1a\x2a\x89\x4f\xaf\x2f\xf9\x10\x6e\x6e\x4f\x4c\x39\x71\x1d\xc7\x69\xc8\xa8\x1a\x83\xe2\x2c\xfb\x55\x32\x6a\xaf\xaa\xa6\x60\x85\x46\xab\xff\x56\x26\x97\xd7\xbb\xb6\x88\xba\x6a\xb0\x5f\x63\xbd\xc7\x5c\xe6\xc5\x5e\x38\x43\xb7\x0d\xb1\x45\xe2\xbb\x4e\x5a\xb9\xbe\x13\x76\x1d\x9c\x2e\x15\x4d\xe4\x9c\xb0\xee\x35\x9d\x60\x56\x3f\x52\x70\x9b\x6d\xca\x77\x57\xab\x53\x53\x1f\x1e\x09\x96\x49\x28\x2d\xf2\x77\x16\x2d\x83\x51\x38\xc3\x39\x09\xde\xe6\xb2\x75\x0a\x8a\xbc\x2c\x65\x23\x3f\x95\xf5\x6d\xcb\xaa\xc1\x49\x91\xcb\x23\xb7\x84\x6a\xd0\xd4\xac\x26\x33\xe8\x5f\x18\x69\x5e\x29\x3b\x78\x8d\xcb\xb2\x5c\xad\x0c\xbf\x7e\x94\x94\x12\xa9\x57\xaf\x57\x69\x55\x96\x3d\x8d\x88\x09\xc7\x87\xc3\xc8\x7f\x65\x19\x5c\x4d\x3e\x07\xab\xd5\xa3\xc6\x7e\x31\x00\x23\x75\xa4\x51\x7c\x2a\x02\xf5\xa3\x62\xed\xd5\x82\xd3\x7d\xea\xab\x8a\xb9\xee\x34\x55\x90\x6a\xf3\xbd\x3d\x3a\xf7\xd7\x3c\xe7\x44\x90\xd5\xd5\xe4\xf3\x50\xa5\xc5\xa3\xe6\xa1\x53\x9f\xb1\x86\x50\x57\x9c\x57\x9a\x6d\xb8\xe6\x97\x85\xcc\x2f\x77\x35\x33\xab\x2d\xca\x6d\x05\xb0\x43\xb7\xce\x94\xdb\x0c\x93\xa2\xef\xf5\xad\x73\x93\xdc\xca\x74\xba\xdc\x65\xbb\xf2\x9b\xef\xb7\xdc\x6c\x87\x5a\x25\x92\x87\x7f\xa8\xb4\x55\x01\x84\x9e\x5d\x07\x7a\x7e\x3b\x34\xbb\x74\x56\x18\x7b\x15\xde\xf0\xbc\x2e\x1a\x7e\x43\xb7\x07\x65\xb9\xc9\x47\xa7\x2d\x7b\x67\x4e\x5b\x4c\x07\x25\xf1\xfe\xac\x3d\x28\x4d\x7f\x70\x5e\x7e\xe3\x98\x6e\xb8\x61\x5b\xd6\xa9\xdb\xfa\xce\x75\xd4\x81\x05\x05\xe6\xaa\x6b\xde\xdc\xb6\x8f\x22\x6f\xab\xf7\xf2\x40\x6c\x11\x9f\x01\xc9\x32\x4c\x23\xaf\x5e\xeb\x83\x9d\xb6\xfe\xc3\xc9\xeb\xb6\xb2\x8f\x23\xb5\xfa\xc6\x3e\xda\xf0\x07\x35\x87\xbd\x8a\xed\xa8\x16\xdd\x75\xe2\x10\x50\xbf\x2b\xd8\xd6\x19\x29\x18\xa1\x58\xc7\x96\x5b\x50\x7e\xeb\x18\x32\x26\xf1\x66\x6e\x8e\x49\xfc\xa5\xe7\x90\xfd\x78\x5f\x76\x10\xe9\xc6\xfd\x92\x93\x48\x37\xda\x43\x8e\x22\xae\x23\x48\xdc\xb1\xa1\xc6\x44\xcd\x96\x8a\x60\x1d\x49\xf9\xd4\xdc\x3e\x63\x12\xfb\xbb\xa9\x6a\x7f\x76\x10\xda\xae\xe9\x20\x69\x5a\xec\x6f\xa6\xc8\x98\xc4\x5c\xd1\xfa\xae\x0b\x76\x01\x03\x3d\x6e\xda\xf3\x4c\x6b\xd6\xd4\x45\xcf\xcc\x52\xef\xb2\xa8\x39\x4b\x15\x6a\x41\x0d\x4f\xf8\x27\xe5\x82\xa6\xb1\x1e\xaa\xba\x67\x27\x0b\xc3\xd3\xec\xd1\xae\x3f\xbe\x74\x8f\x57\x16\xd0\xcf\x33\x5e\x19\x8b\x0f\x99\xaa\x0e\x66\x79\xe0\x30\x65\xf9\xef\x38\x4c\x1d\x87\xa9\xe3\x30\xf5\x6f\x18\xa6\xcc\x66\xff\x0e\x47\xd5\xe3\x04\xf5\x4f\x4c\x50\xf0\x5f\x9a\xa0\xbe\x5f\x32\x1e\xe7\xa6\xe3\xdc\x74\x9c\x9b\x8e\x73\xd3\x4f\x36\x37\xd9\xb3\xcc\xae\xb9\xe9\x1c\x13\x6c\xcc\x4d\x91\x5a\xa8\xbf\x41\x75\x8f\x4b\x16\xab\xe7\x83\xf7\x4d\x27\x9c\x7d\x9f\x36\xd7\x92\xf5\x7c\x50\x6e\xf8\xc7\xfe\xe6\x6c\x0f\x08\xd2\x5b\x1f\x6d\x47\xd9\x56\xec\x72\xd4\x07\x42\xc5\xbb\x54\xd0\x44\x91\x6a\xae\xe8\x50\x97\x6d\x05\xf9\xee\xce\xfb\xea\xef\xc2\x60\xf9\x6f\x9b\x49\xcf\xd8\x3c\x93\xbf\x9d\x0e\xfd\x3b\x00\x00\xff\xff\xdd\xfa\xa4\x1b\xae\x24\x00\x00")

func cftGoTemplBytes() ([]byte, error) {
	return bindataRead(
		_cftGoTempl,
		"cft.go.templ",
	)
}

func cftGoTempl() (*asset, error) {
	bytes, err := cftGoTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cft.go.templ", size: 9390, mode: os.FileMode(436), modTime: time.Unix(1538638943, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _controllerGoTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x6d\x73\xdb\x36\xf2\x7f\x4d\x7e\x0a\x94\xff\x36\x25\x53\x85\xfa\xe7\xdd\x8d\xe7\x7c\x1d\x57\x76\x53\x4f\x5d\xdb\x67\x3b\xcd\x8b\x4c\xc6\x85\xc9\x95\xc4\x98\x02\x58\x00\xb2\xe2\xaa\xfa\xee\x37\x8b\x05\x28\x50\x96\x14\xa9\xc9\xf4\x6e\xee\x9a\xc9\x8c\x49\x60\xb1\xd8\x87\xdf\x3e\x00\x54\xbf\xcf\xfe\x41\xff\xd8\xf1\x05\x3b\xbf\xb8\x61\x27\xc7\xa7\x37\xec\xe6\x87\xd3\x6b\xf6\xfd\xe9\xd9\x09\xfb\x7b\xfb\x2f\xee\xf7\xd9\xcd\xb8\xd2\x6c\x58\xd5\xc0\x2a\xcd\xf8\xd4\xc8\x11\x08\x50\xdc\x40\xc9\x1e\x2a\xce\x7e\xe1\x33\xfd\x42\x36\x38\x22\x15\xf3\x73\xbf\xe0\xca\xd3\x21\x7b\x94\xd3\xaf\x4b\x56\x57\xf7\xc0\xcc\x18\x58\x31\xe6\x62\x04\x8c\x8b\x47\x33\xae\xc4\x88\xf1\x3b\x39\x35\xcc\xb4\x3b\x4c\xf8\x3d\x30\x28\x2b\xa3\x99\x91\x76\x45\x6e\x60\xd2\xd4\xc8\x8d\x44\x10\x76\xb4\xb9\x1f\xf5\x0b\x59\xc2\x08\x44\x9f\x6b\x0d\x46\xb3\xb2\x52\x50\x18\xa9\x1e\xf3\x38\x6e\x78\x71\xcf\x47\xc0\xe6\xf3\xfc\xba\x81\x22\xbf\x02\x2d\xa7\xaa\x80\xfc\x9c\x4f\x60\xb1\x88\xe3\x6a\xd2\x48\x65\x58\x1a\x47\xf3\xf9\x0b\x56\x0d\x19\xd1\x9d\xea\xc1\x54\x1b\x39\xa9\x7e\x83\x72\xb1\x88\x23\x36\x01\xc3\x1f\x5e\xb2\xe4\xfe\x6f\x3a\xaf\x64\x9f\x37\xd5\x84\x17\xe3\x4a\x80\x7a\xec\xa3\x0c\xbc\xa9\x74\x1f\x89\xfa\x0f\x2f\x93\x38\x4a\x46\x95\x19\x4f\xef\xf2\x42\x4e\xfa\x7c\xa6\x6b\x7e\xa7\xf1\xef\x0b\x0d\xea\xa1\x2a\xa0\x35\x93\x5d\x3b\x86\xba\x01\xa5\x13\x92\x01\x84\xdd\x31\x51\x30\xac\xa1\x30\x49\xbc\x27\xb7\x42\x8a\x61\x35\x4a\x56\x14\xfa\xe7\x14\xa6\xa8\x30\x63\x8c\xed\xc5\xee\x57\x5c\x98\xd8\x75\x85\x54\xd0\xb5\x41\x1f\x87\xac\xc6\xab\x7c\x2b\x2e\x0a\x59\xc3\x84\x8b\xbe\x36\xaa\xe0\xda\xf1\x48\xb4\x51\x95\x18\xad\x28\x2b\x9b\xfb\xca\x74\xd6\x17\x63\x55\x69\x23\x9b\x31\xa8\x31\x54\xa2\xef\x65\x7a\x71\x5f\x99\x24\x8e\x78\x53\xc1\x07\x03\x42\x57\x52\xe8\x87\x97\x77\x60\x78\x47\xb0\xe5\xe4\x0b\xf4\x0c\xa8\x07\x50\x4b\x3f\x75\x08\xfa\x6e\x39\x7a\xcd\x31\x28\xea\x0a\x84\x79\x31\x92\x7d\x23\x65\xad\xfb\x05\x2f\xc6\x80\x8e\xe0\x33\xcd\x9b\x6a\x3f\x03\xda\x1d\x57\x67\x72\x3e\x43\x13\xf0\x99\xfe\xf9\x25\xaf\x9b\x31\x4a\xff\x59\x98\xf6\x1f\x1c\x3f\xe2\x4e\x9a\xec\xc7\x9b\xd6\xb8\x3f\x1a\x4c\xff\x01\x14\x5a\x0a\xca\xbe\x79\x6c\xa0\xfc\xe8\xbe\x1d\xe0\xf9\x38\xe2\x06\x8d\x9d\x5f\x52\x40\x5a\x24\x26\x3e\x2a\x37\xd1\x74\x40\x92\xc5\x18\xfb\x3e\x80\x31\x05\x61\xfc\xcb\xbb\xf7\x50\x18\xa6\x8d\x54\xc0\x4a\x18\x56\xa2\x42\x26\xf1\x03\x57\x4b\xda\x43\x66\x01\xe6\xf6\xf1\xc3\xf3\x38\xc2\x2c\x70\xc0\xe8\x5f\xb2\x21\x47\x24\xbd\x38\xba\xac\xa7\x8a\xd7\x07\x1b\xe8\x68\x96\x28\x5f\x29\x39\x6d\x1c\x4b\x42\x4b\x6e\x87\x90\x57\x2f\x8e\x7e\x26\x53\x1e\x04\xd3\x6e\xa8\x17\x47\xd7\x85\x6c\xbc\x38\xeb\x00\x6e\x05\xd2\x0d\x2f\xc0\x52\x96\xbd\x38\xfa\xb1\x12\xa5\x57\xc0\x25\x8c\xfc\xe6\xb1\x81\x8b\x61\x1a\x60\x2b\xf7\x22\x23\xf9\x62\x31\x5f\x64\x96\x55\x9a\xe1\xa6\x63\xa9\x8c\x65\x7c\xc0\xde\xbe\xa3\xe8\x9c\xc7\x1d\x17\xb6\x8a\x5a\x5a\x81\xb4\xec\x05\x86\x2d\x52\x29\x9b\xc7\x37\x13\x22\x5d\x84\x36\x5b\x5a\x73\xe9\xd3\x20\x05\x2c\x7a\xf1\xc2\x7a\x78\x20\x85\x51\xb2\xae\x41\x31\x05\x8d\x02\x8d\x28\x64\x9c\x15\xcb\x71\xe7\xf4\xa1\x54\x5d\xff\x17\xd6\xbf\x4c\x39\x31\x74\x8c\x78\x0d\x19\x6a\xa3\xa6\x85\x61\xf3\x38\xa2\x2c\xe9\x2c\xf7\x9c\xde\xf2\x81\xfd\x63\x27\x0d\x7c\x30\x6e\xd2\x61\x87\xc6\x82\xa8\xd2\x60\x58\xfb\x92\x5f\x53\x50\xf8\x98\xf0\xa6\x3f\x15\x06\xd4\x90\x17\x10\x47\x46\x36\x55\x71\x74\x75\x6e\xb9\x92\xa1\x9d\xca\xe7\x30\x0b\x84\x2c\x14\x70\x03\xa1\xbe\xa8\xe8\x8c\x9b\xc2\x56\xca\xad\x1a\xbb\xc5\x65\x3c\x9c\x8a\xa2\xcb\x37\x75\x2a\x77\x95\xed\x31\xaf\x6c\x57\xcf\x1e\xfb\x83\x7a\x66\xec\x79\xa0\xcb\x3c\x8e\x14\x98\xa9\x12\xec\xd9\x72\x74\x1e\x47\xce\xfe\x1e\xba\xf4\xd6\xa3\x71\xdc\xfe\xa0\x1d\xb7\xc2\xc4\x51\xc7\xec\x07\x1d\xe1\x7a\x71\xb4\x70\x86\xbc\x36\x5c\x99\x37\x68\x29\xb2\x17\x68\x6b\xbb\x4a\x68\xc3\x05\x9a\x47\x0e\xd9\x05\xd9\xef\x7a\xbd\xfd\xb8\x28\x19\x2f\x8c\x66\xd2\xf6\x17\x13\x32\x64\x5a\x84\x5a\x65\xc1\x3e\xa9\xf0\x31\xe9\x5c\xda\x43\xcf\x34\x83\xb1\xed\x70\x1c\xe0\xe6\x8b\x8c\x81\x52\xd2\xd9\x83\xf6\xfa\x81\x8b\xb2\x06\xa5\xd9\xc1\x21\xb3\x05\xa6\x0d\x9f\x93\x07\x10\xc6\x4d\x7f\x3f\x15\x85\x46\x83\x1d\x95\x25\x3e\x5b\xc3\x14\xb9\x14\x47\x25\xc6\x7f\xf4\xba\x29\xb9\x01\x9a\xc1\x61\x7a\xc7\x99\x63\xa8\x21\x9c\xa1\x77\x6b\xac\xf5\xad\x41\x64\x4b\x7d\x61\x54\x8d\x22\xd9\x97\xfc\x1c\x66\x69\x91\x17\x1e\x2a\xb9\x77\x08\x2b\xf2\x8e\x0b\xd8\xcb\x2c\x8e\x8a\xdc\x43\xbc\xc7\x6e\xe9\x3f\x73\x8c\x90\x6b\x7e\x05\xa3\x4a\x1b\x50\xe9\xe6\x54\xcb\x9e\x6d\x4d\x5b\x71\x34\x92\x01\xc3\xc0\x0d\x24\x6e\x60\xb4\xd4\xaa\x45\xe6\x50\x99\xf7\x4a\x16\x66\x1f\xeb\x0b\x33\xa0\x12\x89\x5e\xe8\xe8\x94\x5f\x9d\x5c\xdf\xd0\x64\x9a\xc5\x11\xe1\x49\x21\x1d\x05\xca\x39\xcc\xde\xd0\x58\xea\xd5\xe8\xb1\x16\x0d\x3d\xb6\xea\x67\x3b\xe2\x76\x23\x45\x1c\xcb\x9c\x34\xd8\xaa\xf9\x52\xfe\x36\xa0\x44\x55\x23\xea\xd7\xba\xb2\xdf\x67\xa1\xfa\x6c\x56\xd5\x35\x33\xdc\xf5\xde\x13\xd0\x9a\x8f\x30\x36\x94\x9c\xd8\x11\x6b\x3d\x8b\xfd\x46\xc9\x02\xb4\x0e\xb0\x1f\x32\xda\x94\x43\x26\x7a\xc4\x9e\x93\x0b\x7e\x22\xe6\xdf\xc9\xf2\x31\x00\x7d\x2d\x47\x23\x32\x9e\x5b\x79\x66\x07\xe2\x08\xeb\x34\x1a\x2d\x30\x9d\xcf\x8d\x51\x35\x44\xc6\xb9\xdd\x9b\xdf\xd5\x80\x8c\x22\x24\x63\x87\x76\xc2\x9b\x1d\xc1\xe3\x66\x88\x01\x4d\xb7\xd5\x32\x8e\x16\x0c\x6a\x4d\xeb\xc9\xbf\xd7\x88\xd9\x5b\x94\x67\x99\xd9\xce\x61\xf6\xbd\x54\xa4\x92\x53\x94\x30\x60\x1f\xb3\x38\x6a\x43\x57\xf7\x50\x33\xab\x8d\xe7\xd6\x7a\x8c\x9a\x01\xc2\x73\x9a\x24\x59\x7e\x56\x69\x93\xd2\xf1\xc1\x3e\x5f\x34\xb6\xcb\xb1\x68\x46\x15\x91\xd3\x17\x87\xe8\x4e\x2b\x9f\x33\x55\xfe\xa6\x32\xe3\x13\x34\x5f\x0a\x4a\x65\x39\x3d\x26\x64\xd0\x11\x18\x83\xa5\x60\x73\x27\x82\xbc\x3d\x50\x40\xa9\x38\xc2\xa8\x8f\x30\x17\xde\x2e\x91\x89\x0a\x50\xf9\x6e\x15\xcb\x4f\x0d\x4c\x34\x09\x52\x0d\xdb\x71\x8c\x34\x33\xd5\xf8\xa7\xb8\x3f\x3d\x66\x87\x64\xe1\x4b\xae\x34\x94\xce\xe3\x6f\x13\x9a\x2d\x93\x77\xb4\xde\xbb\x4a\x85\x41\xde\x4e\x78\x4f\x75\x66\x9d\xbb\x48\x5a\x6c\x0a\x62\x8b\x02\xcb\xe8\x8b\x43\x96\x24\xec\xd9\xb3\x00\x28\x34\x84\x9b\x71\x21\xa4\xa1\xf6\x11\xb5\x9a\xf0\xe6\x2d\xa1\xa8\xed\x68\xa2\x28\x72\x02\x1e\x27\x07\x5b\xa5\xef\x2d\x69\x51\xa6\x2d\xd4\x76\xda\xd1\x7b\x1f\x60\xf7\xb5\x7e\x49\x87\xc2\xae\x5a\x10\x02\xb6\xd0\x92\xdd\x93\x77\x68\xf1\xe4\xea\xe2\xec\xec\xbb\xa3\xc1\x8f\xb7\x83\x8b\x9f\x2e\xcf\x4e\x6e\x4e\x48\xf7\x48\xde\xbd\x6f\x01\x59\xda\x2c\x6f\xa5\x4b\x7d\xda\x5e\x89\xaf\xde\x56\xed\x33\xe7\xfa\x55\x58\x76\xd1\x64\x45\x8f\x7c\x90\x40\x21\x55\x09\x2a\x3f\x22\x2f\x40\x69\x2b\xd8\x30\xb5\x82\x05\xae\xe9\xb9\x23\x64\x6e\xe7\xd1\x0e\x6f\xb8\x12\xae\x68\xda\x43\x62\x7e\x23\x07\x7c\x02\x75\xea\x8e\x89\xf9\x8d\x3c\x93\x33\x50\xe9\x0e\x36\xca\xb2\xb5\x9a\x75\xc9\xae\x80\x6b\x29\x48\x4d\x97\x15\xf6\xf0\xc0\xf1\x09\x9a\x7d\x9b\xfd\xa7\x36\x51\xd2\xa2\x3f\xe6\x80\x1d\x94\xd8\x89\x28\xd4\xf4\x7f\xc4\xa1\x11\xea\x78\xc8\x2a\x51\x28\x98\x80\x30\x57\xb2\xae\xef\x78\x71\x3f\x90\x53\x61\x36\x79\x63\x0f\xfb\x04\x65\xe4\x2f\xa7\x07\x4e\x3f\x97\x6a\xc2\xeb\x7f\x53\x10\xc7\x51\x44\x85\xa2\xd3\x17\x05\x5d\xde\xda\x0e\xde\x36\xd0\xa8\x1a\xab\xfc\xc9\x05\xbb\xf4\x79\x1c\xd9\x02\x22\xef\xde\xe7\xe9\xf3\xcd\x1d\x59\x96\x1f\x03\x34\x03\xd9\x3c\xa6\xd8\x92\xb9\x26\x4c\x48\xb3\xfe\xd6\xe3\xa8\x2c\xed\x8d\x47\x35\x64\xda\xd7\xd2\xae\x4a\x36\xbd\x24\xec\xf7\xdf\xb7\x12\xac\xc9\x3f\xf6\x3c\x30\xb4\xfd\x6b\xb7\x59\x47\x3f\xb5\x0d\x79\x66\xe9\xe4\xd4\x34\x53\xb3\x6c\x5e\x86\x26\x1f\xd8\xf3\x22\xd5\x0b\x22\x5a\x45\x8d\x3b\x89\xe5\x9d\xc6\x6d\x6d\x7b\x32\xf4\xfd\x89\x3d\x84\xae\x6d\x50\xa8\x2f\x62\x5f\x7f\xa5\xbf\x4e\x7a\x4c\xdb\xf7\x2c\xf6\xf7\x17\xe8\x3f\xfb\x42\xf7\x94\xab\x7b\x9e\x8a\xa1\x1c\xa6\x09\x2f\x4b\x28\xb7\xb2\x66\xb3\xca\x8c\x99\x76\xcd\x4a\x67\xb3\x9e\xeb\x2e\xd3\xe7\x64\x0c\xd7\xd2\x94\x59\xb6\x6d\xcb\x87\x0a\x66\x8c\x1b\x36\x36\xa6\xd1\x07\xfd\x7e\x21\x85\x96\x35\xe0\x79\x21\xe7\x13\xfe\x9b\x14\x74\x69\x59\xcb\x69\x39\xc4\x60\x40\xb7\xf7\xc7\x72\x02\xdf\xfe\x5f\xdf\xca\xd1\x2f\xc1\xf0\xaa\xfe\x96\x84\x2a\x0f\xbf\xd2\xc9\x16\x51\xe2\x28\xba\x25\x37\xad\x26\x97\xa5\x7b\xbd\x3a\xcb\xc6\x69\x23\xc3\x1e\x4b\x06\x57\x27\x47\x37\x27\xb7\xa7\xe7\xb7\x97\x57\x17\xaf\xae\x4e\xae\xaf\x93\x1e\x4b\x92\x0d\x3d\xe8\x2e\xde\xf6\xce\xb6\xf2\xa1\xb3\x35\x85\x30\x05\x25\xfa\x90\x8e\x5a\xb5\x06\xba\x2b\xda\x1c\x19\xa1\x5a\x9d\x13\xda\xe6\xe8\xa5\x03\x49\x2a\xeb\xf2\x02\xd3\x93\x80\xd9\xc5\x9a\x58\x96\xd2\x06\xb3\x25\xda\x3d\x9e\x23\x61\x97\x11\xcf\x7d\xd2\x80\x4d\x01\xfb\x85\x2f\xda\x9a\x72\x91\x90\x99\xb5\xda\xc7\xb3\x09\xe9\xee\x13\x8a\xfb\x7a\x90\x9f\x6a\xeb\xee\x81\x9c\x34\xd8\x02\xa6\x72\x83\x24\x3d\x36\xe4\xb5\x86\x0c\x5b\xe9\x2f\xfc\x9d\x21\xaa\x71\xf2\xeb\x94\xd7\xb8\xcc\xeb\x88\xa9\x7d\xb1\xe8\xa1\x4e\xdd\xa1\x6c\x5b\xde\x91\x72\xa7\xc4\xf3\xda\xc3\xba\xb8\x77\xaa\x7f\x96\xd4\xd3\xa2\xf1\xe3\xf9\x41\xc0\x8c\x35\x5c\xf1\x89\x66\x5f\x7d\xf3\x60\x4f\xbf\xb2\x2e\xf1\x39\xb1\x3a\x53\x7c\x09\x89\x2a\xed\x9b\xa3\x28\x68\x77\xc9\x52\x4e\x02\x97\xa4\xda\x6d\xff\x4b\xb2\x94\xf4\xfa\xb8\x87\x8f\xe6\xa9\xd7\x97\xc7\x7f\x7a\x9e\x62\x6c\x87\x44\xe5\x83\x6e\x05\xe9\x08\xdd\x9d\x12\x16\x5d\xbf\xfd\x69\x1d\x07\x6d\x87\xfa\xec\xd2\x1c\x44\x41\x60\x1e\x07\x27\xc8\x2c\x5e\x63\xf5\xbd\xe2\xd1\x9e\x47\xf7\x68\x05\x7c\x87\x4a\x2d\xdd\x06\xa0\xd3\x21\x77\x7b\x74\x85\x3c\x77\x70\xaf\xb7\xd7\xb6\x52\x14\x6f\xfb\x44\x6c\xdd\xbe\xf5\xe4\xf1\xe4\xd6\xcc\x5e\x6b\xf8\x5b\xe3\xd5\xeb\xaf\x5d\x2e\xcf\x3e\xe1\x2e\x4b\xb5\x57\x96\xbb\xdd\x64\x2d\x8f\x4b\xf9\x2b\x30\x29\x1d\x6e\xdc\xad\xd6\x2b\xe8\x5c\x6a\x3d\x45\xcc\x67\xbb\xd2\x0a\x4f\x2f\xee\xfe\xd6\x52\x61\x58\xd8\xdb\x2c\xbf\x2a\xac\xe6\x21\x91\x63\x1e\x7a\x27\xb8\x83\xda\x40\xf1\xcd\xcb\x38\x5e\x66\xba\xdd\x0d\xe5\x7a\x94\x90\xfb\x1f\x35\x4f\x9b\xbc\x3c\xb3\x35\xe6\x58\xb9\x18\xb6\x80\x5c\x73\x32\xdd\x0f\x86\xbd\xb6\x8d\x0e\xde\xb1\x9f\xf1\xaf\xca\x9e\xc5\x5a\xd0\x6e\x4b\x61\x3d\x82\x74\xf6\x17\xa6\xd7\x60\x5a\x54\x75\xef\x93\x80\xbd\xbe\xeb\x74\xee\xda\x85\x98\x4e\xd5\x36\x1c\xf0\x61\xfd\x92\xf6\x06\xd8\xe3\x82\xda\xde\x4d\x4d\xa8\xee\xb6\x9c\xb6\x7c\x04\xf5\xc8\x21\x30\xdc\xc8\x57\x7b\x6a\x0d\x74\xa7\x67\x44\xa7\xd0\x70\xfa\xe9\x97\xe7\x6e\x03\xd7\x07\xc4\x51\xf0\x35\xfb\xcb\x46\xc9\x06\x37\x25\xbf\x5d\xb8\x2e\xa5\x18\xc3\x84\xe7\x97\x4a\x36\xa0\x4c\xe5\xbf\x6d\xbb\xa2\x60\xd7\xe4\x37\x30\x69\x6a\x6e\xfc\xcf\x86\xa2\xf9\x9c\xc6\x7f\x84\x47\xc4\xbf\x45\xb4\xb7\xd5\x92\x36\x4d\x3c\x9d\x1b\xa3\x6f\x61\x9e\xf0\x98\x1b\x3e\xbf\xb8\x7b\x7f\xb0\x62\xa8\x81\xfb\x86\x4a\x66\x5c\x84\x5f\x28\xac\xcb\x9c\xdc\x9e\xf7\xb5\xfd\x16\x69\x25\x61\x87\xac\x23\x99\x53\xc3\xd7\xc8\x3d\xd8\x38\x2b\xbe\x7d\xa2\x81\x25\x48\xde\x79\xce\xa2\x5c\xee\xe2\x3e\xf8\xff\x67\x27\xd5\x6e\x3c\xee\x07\x71\xd2\x49\x3f\x8a\xe2\xa8\x2c\xed\xef\x51\x78\xed\x63\x4d\xaf\x45\xfd\x5e\x70\xc6\x2e\xc8\xeb\x80\x9b\x84\x2a\x78\x3c\x07\xf5\xa0\x0b\x9a\xb0\x3a\x3c\xfd\x58\xf0\x69\xc5\xe1\xaf\xf4\xff\x99\xd2\xff\x96\x1c\xe9\xf2\xa3\x0b\x9b\x5d\xfa\xf5\xb5\x3b\x2c\xd7\xbf\xe1\x95\x79\x2d\x4c\x55\x5b\x16\xc4\xad\xa4\xf2\xd2\x81\x0f\x2d\xf6\xd0\xd9\x0e\xee\x27\x30\xd2\x6c\xdb\xd1\x86\xa1\xf5\x02\x74\xfc\x89\x20\x08\xae\xcf\x08\x06\xfe\x5e\x6d\x57\x20\x3c\xe9\xc1\xda\xdf\x9c\xad\xab\xd9\x78\xa2\x58\x1e\x21\xd6\x18\xd0\xff\xbe\x06\x8b\x8b\xf6\x8f\x07\x87\xcb\x9f\x66\x2d\xe6\x73\x97\x40\x83\x7a\x15\x7e\xd1\xdd\x89\x75\x7b\x80\x84\x5f\x97\xab\xed\xed\x0e\x4b\x4e\x3e\x18\x50\x82\x2c\x94\x10\xe9\x97\x6a\xe5\x84\x15\x94\x31\xba\x41\x0e\x17\xb9\x7d\xd2\xe0\xb0\xb9\xf6\xb6\x32\x09\x19\x87\x0c\x6c\xf5\x9b\xcf\xd3\x4a\x94\xf0\x21\x90\xef\x52\x2a\xa3\xd9\xff\x67\xf6\x61\x81\xfe\x68\x4d\x74\xc8\x78\xd3\x80\x28\x53\x3f\x82\x0c\x56\xc5\xee\x1c\xe8\xba\x8f\xbb\xfa\xa4\xe5\xb8\x95\xee\x90\x79\x31\xbc\xb7\x3e\xee\x78\x82\xf0\x4f\xbc\xc1\x6d\xc8\x74\x13\x7c\xf9\x4c\xce\xef\xb0\x5f\xe7\x53\xec\x32\x36\x7e\x3b\x0f\xb6\xab\x0c\x4c\x90\x6e\xb9\xde\xae\xf4\xbf\xf3\xb3\xf3\xd4\x55\x24\x07\xac\x1d\xf8\x99\xd7\x53\xf7\xf3\xbf\x6e\x13\xb0\x13\xba\x5a\xe1\x3f\x8a\xa9\xf5\x7a\x65\xfe\x07\x7f\xd6\xa2\x2d\x56\x96\x63\xbb\xa0\x65\x77\xef\x6d\x05\x49\x40\xe9\xab\x1f\x4a\xd0\x02\x65\xf7\xae\xa8\x93\xbb\x56\xfa\xa2\x1d\x72\x95\xbf\xcd\x7c\xf2\xc5\xec\x5f\x01\x00\x00\xff\xff\xaf\xce\x6e\xce\x17\x30\x00\x00")

func controllerGoTemplBytes() ([]byte, error) {
	return bindataRead(
		_controllerGoTempl,
		"controller.go.templ",
	)
}

func controllerGoTempl() (*asset, error) {
	bytes, err := controllerGoTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller.go.templ", size: 12311, mode: os.FileMode(436), modTime: time.Unix(1538650854, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _template_functionsGoTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xc1\x8e\xdb\x36\x10\x3d\x8b\x5f\x31\x15\x02\xd4\x2a\xbc\x22\xf6\x56\x04\x4d\x0f\x6d\x36\x89\x91\x60\xb7\xc8\x1a\xc8\xb1\xa1\xa5\xb1\x34\x30\x45\x12\xe4\xc8\x8e\x21\xe8\xdf\x0b\x92\xb2\x03\x6c\x91\xee\x56\x17\x4a\x23\xce\x7b\xf3\x1e\x67\x28\x25\xfc\x9e\x1f\x78\xfb\x00\xf7\x0f\x5b\xb8\x7b\xbb\xd9\xc2\xf6\xc3\xe6\x11\xde\x6d\x3e\xdd\xc1\x6f\xd7\x47\x48\x09\xdb\x9e\x02\xec\x49\x23\x50\x00\x35\xb2\xed\xd0\xa0\x57\x8c\x2d\x1c\x49\xc1\x57\x75\x0a\x37\xd6\xc5\x88\xf5\x37\x8d\x6d\xb1\x43\x03\xce\xdb\x06\x43\xf8\x1a\x01\x36\x7b\x38\xdb\xf1\xe7\x16\x34\x1d\x10\xb8\x47\x68\x7a\x65\x3a\x04\x65\xce\xdc\x93\xe9\x40\xed\xec\xc8\xc0\x57\xa2\x41\x1d\x10\xb0\x25\x0e\xc0\x36\x65\xd4\x8c\x83\xd3\x11\x2d\x57\x62\x52\xd4\x1d\x3a\xb9\x30\x4a\x15\x02\x72\x80\x96\x3c\x36\x6c\xfd\xb9\x16\xc2\xa9\xe6\xa0\x3a\x84\x1e\xb5\x43\x1f\x84\xa0\xc1\x59\xcf\xb0\x12\x85\x3a\x85\x46\x13\x1a\x86\xb2\x23\xee\xc7\x5d\xdd\xd8\x41\xaa\x53\xd0\x6a\x17\xe2\x7a\x13\xd0\x1f\xa9\xc1\xab\x34\x99\xc8\x52\xce\xb2\x04\x64\x79\x44\x1f\xc8\x1a\x6c\x25\x9f\x1d\xb6\xf2\x69\x56\xad\x4e\x41\x1e\x6f\x95\x76\xbd\xba\x2d\x45\xf1\xff\xd8\xac\xd9\x53\x57\x8a\x62\x40\x56\xc7\x5b\x28\x0f\xbf\x86\x9a\xac\x54\x8e\x06\xd5\xf4\x64\xd0\x9f\xd3\x46\xe5\x28\xc8\xb8\x49\x1e\x6f\x4b\x51\x89\xe8\xd3\x3d\x9e\xe0\x0b\x69\x0d\x1e\x79\xf4\xe6\x62\x42\x34\x74\x87\xe0\xa2\x5d\x2d\x90\x59\xfc\x4d\xf6\x2a\xc6\x20\xf6\xa3\x69\x62\xf2\xaa\x82\x0f\x4b\xca\x24\x8a\x05\x64\x89\x4c\xa2\x00\xf8\x38\xee\xd0\x1b\x64\x0c\x9f\x31\xd8\xd1\x37\x78\xaf\x06\x7c\xfd\x83\xf8\x5a\x14\xc5\x34\xdd\x80\x4f\x27\xff\x8a\x4c\x8b\xdf\xd6\xf0\x0a\x35\x0e\xf1\x14\x5e\xbf\x81\x7a\xc3\x38\x04\x98\x67\x51\x14\xef\x91\xa7\xe9\xf2\xb3\x7e\x74\xd8\xd4\x1f\xc9\xb4\xf3\xfc\xc7\x39\x93\x3c\xb3\xe1\xc2\x86\x31\x24\x8a\x59\xcc\xc9\x93\x8b\xa0\x16\xf7\x64\x30\x80\xd2\x3a\xa9\xcf\x71\x88\xd2\x99\xac\x09\xc0\xbd\x62\x50\x1e\x01\xbf\x39\x1b\x8d\xfa\x97\x4d\xf1\xb8\xaf\x78\x81\xfd\xd8\x30\x4c\xe2\x47\xae\x24\xe8\x55\x60\x4f\xa6\xab\xe0\xf2\x22\x5e\xee\xc8\x33\x7a\x33\xfe\x2f\xb9\x61\xea\x3f\xd3\xb2\x86\x4c\x73\x59\x2b\x58\x91\x61\xf4\x7b\xd5\xe0\x34\xaf\x01\xbd\xb7\x7e\xa9\x21\xfb\x34\x0b\xf1\xd2\x82\xa4\x7c\xee\x08\xe0\x14\x9b\x6f\x4f\xa6\x4d\xce\xf9\xc5\x0d\xd8\x9d\xc1\xa8\x01\x73\x9f\x3d\x83\xb1\xca\x82\xe0\xa9\xb0\x08\x70\x55\x17\x3f\x82\x53\x0d\xfe\xa7\xce\xd8\xc4\xda\x76\x1d\xfa\x28\x64\xc1\xfb\x94\x02\xa2\xc8\xf3\xfc\x88\xbc\x86\xbf\xe3\xef\xeb\xed\x50\xdf\xe3\xe9\x9d\xf5\x99\x77\xa9\xa6\xfe\x7c\xf7\xb8\xcd\x91\x2a\x0e\x46\xd6\x95\x78\x12\xf4\x05\xab\x7e\xaa\xec\x2f\x3d\x7a\xa5\xa3\xae\x79\x5e\x5d\xab\xae\xea\xf7\xc8\xe9\x73\x0d\x79\xce\x63\xe0\xc1\xa5\x46\x9c\xe6\x4a\x14\xb4\x4f\xd8\x3f\xbd\x01\x43\x3a\xea\x58\x84\xd4\x5f\x88\xfb\xbb\xa8\x6e\x85\xde\x57\x75\x7e\x2d\x93\x5e\xe8\x90\x39\x5e\xad\x21\x76\xb3\x75\xd4\x94\x95\x28\x2e\x63\x5c\x96\xa9\xdc\x38\x18\xd7\xd1\xfe\x2e\xc4\x90\x16\xb3\xf8\xde\x16\xff\x04\x00\x00\xff\xff\xe5\x13\x3f\x32\x30\x06\x00\x00")

func template_functionsGoTemplBytes() ([]byte, error) {
	return bindataRead(
		_template_functionsGoTempl,
		"template_functions.go.templ",
	)
}

func template_functionsGoTempl() (*asset, error) {
	bytes, err := template_functionsGoTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template_functions.go.templ", size: 1584, mode: os.FileMode(436), modTime: time.Unix(1538540481, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typesGoTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x56\x4d\x6f\xe3\x36\x10\x3d\x4b\xbf\x62\x60\x04\x85\xd5\x26\x36\x7c\x5b\x18\xd8\xc3\x6e\x16\x2d\xb6\x69\x36\x45\x9c\xf6\xb2\x58\x20\x0c\x35\x92\xb9\x96\x48\x96\xa4\xd2\x1a\x86\xfe\x7b\xc1\x0f\xc9\xa6\x25\x6d\xd3\x43\x0e\xab\x93\x34\xf3\xe6\x91\x7c\xf3\x21\x4a\x42\x77\xa4\x44\x78\x5e\x91\x4a\x6e\xc9\x2a\x4d\x59\x2d\x85\x32\x30\x4f\x93\x1a\x0d\x79\x5e\xc1\x6c\xf7\x46\x2f\x98\x58\x12\xc9\x6a\x42\xb7\x8c\xa3\xda\x2f\xe5\xae\xb4\x06\xbd\xb4\xa0\xe5\xf3\x6a\x96\x26\x93\x38\xd5\x70\xc3\x6a\x9c\xa5\x59\x9a\x2e\x97\xf0\x53\x89\x9c\x56\x0c\xb9\x89\xbf\xd6\x5c\x6c\x0c\x31\x8d\x76\xe6\xdd\x1b\xbd\xce\x11\x25\x15\x72\x7f\x55\x22\x5f\x33\x6e\x50\x15\x84\xa2\x7e\xfb\x1f\x0b\x2d\xee\x9e\xbe\x22\x35\x6e\xb1\xc3\x61\xb1\x91\x48\x17\x37\x8c\xe7\x6d\x0b\x39\x16\x8c\xa3\x06\xb3\x45\x78\x22\x1a\x41\xa1\x16\x8d\xa2\x98\x9a\xbd\xc4\x73\xb4\x36\xaa\xa1\x06\x0e\x9d\x14\x8b\x87\xbd\xc4\x5b\x34\x04\x00\x1e\xbf\x6a\xc1\xd7\xb3\x4b\xc6\x2b\xc6\x71\xf6\xd8\x63\xfc\xe2\x0e\x15\x30\xd6\x91\x13\x43\x2c\xa8\x5b\xc1\x32\xb5\x2d\x00\x9c\xad\x39\xf0\x07\x8e\x53\xfb\xaf\x9b\xbb\x4f\x6d\x6b\xd9\xbc\x5e\x70\xf2\xc4\x6c\x03\x7f\x60\xd3\xce\x6e\x19\xee\x1a\x23\x1b\x33\xcd\x30\xf0\x07\x06\xe1\xec\x96\xe1\x5d\x9e\x33\xc3\x04\x27\xd5\x7d\xd0\x52\x9f\x71\x8c\x21\x02\x0d\x19\xba\x66\x8f\x69\x9b\xa6\x87\xc3\x15\xb0\x02\xe6\x25\xc2\xbc\x42\x0e\x9e\xef\xbd\xc8\xf7\x8b\x0d\xdd\x62\x4d\x16\x37\xb8\xbf\x25\x52\x32\x5e\x66\xb0\xca\xda\xd6\x85\x28\xc2\x4b\x84\x0b\xc6\x73\xfc\xe7\x12\x2e\xb0\xc2\x1a\xb9\x81\xf5\xdb\x6f\x12\xb4\xad\x2f\x95\x8b\x38\x13\x9e\xe6\xac\x6a\x4e\xcc\x5d\xed\x40\x21\xd4\x79\x74\x57\x4f\x53\x94\xc7\xd2\xea\x4e\x2a\x14\xcc\xf1\xaf\xb0\x77\x98\xd9\xb0\x59\x16\x99\x3e\xd8\x1a\xca\xec\x51\x93\xeb\x4a\x34\xf9\xcf\x42\xd5\xc4\xaa\xf7\x80\xb5\xac\x88\xc1\x4f\xa4\x46\xcb\xcc\x78\xd9\x09\x4c\x27\x81\x36\x77\xd3\x34\x5a\x12\xfa\x72\x2e\x87\xb6\x84\xf7\xa2\xaa\x9e\x08\xdd\x5d\x8b\x86\x1b\x60\xdc\x74\xb1\xea\xd4\xe1\x1b\xe1\x0a\xd0\x29\x65\x4b\xae\xcf\x9c\x54\x42\xda\x74\x75\xa9\xb3\x87\xed\x14\xe2\xc1\xef\x9a\x00\x66\x9d\xfc\xee\x30\x59\x20\xf2\x80\x8d\x53\xf7\x06\xf7\x6d\xdb\xdb\x6c\xd0\x5d\xe1\xbb\x13\xa2\xb4\x1c\x7b\xcc\x23\x5d\xdc\xd8\x1e\xfd\xbb\xad\x34\xff\xe6\x6c\x95\xc6\xc9\x74\x47\xbd\x7c\x4c\xf9\x77\x9f\xbd\x70\x74\x27\xc2\x70\xca\x86\x99\x71\xda\x35\x7e\x5c\x9c\x77\xcc\x58\xc3\x8c\x10\xf5\xc2\x8d\x56\x8a\x0f\xf0\xd0\xae\xb5\x7f\x57\x42\xa2\x32\x0c\x35\xbc\x42\x5d\xc4\xc5\x30\x3c\x7f\x98\xba\x5b\x51\xe5\xfe\xf4\x7e\xdc\x82\x28\xdc\x97\x4b\x5b\xd1\x25\x02\x4c\xc8\xc4\x98\x00\x81\xe9\x44\x80\x6e\x4c\x46\x93\x3d\x4e\xb4\x8a\x20\x6e\xbf\x71\xd4\x3d\x12\x2d\xf8\x37\xa3\x3c\xa4\x3f\x2b\x2b\x82\xcc\x7f\x68\x8c\xab\xce\xc9\xbb\x31\x84\xee\x3e\x7e\x80\xe8\x89\xe9\xb5\x87\xbc\x40\xbd\xb1\xbf\xc5\x51\xca\xe3\x0f\xa3\x2f\x26\x3d\xa6\xdc\x18\xcb\x60\xe6\xfa\x80\x11\xe8\x62\x83\xea\x99\x51\xd4\x76\xfc\x74\xef\x49\xf2\xf9\x8b\x3f\x54\xd2\x1d\x2a\x78\xe2\x59\xf1\x12\x76\xaa\xd0\x04\x72\xf7\x9a\x8c\x91\x3b\xcf\xff\xe5\xbe\x16\xbc\x60\xe5\x2d\x91\x8e\xfe\xf8\x35\xe0\xa7\xbd\x2b\x5e\xc2\xe7\xe4\x75\xee\x5f\xbf\x31\x1d\xcf\x85\xca\x1a\x88\x31\x8a\x3d\x35\xc6\x8f\x05\xff\x93\x8d\xfa\xd0\xe6\x77\x2c\xc9\x8e\x6e\xfa\x92\x36\x79\x45\xb3\x71\x53\x17\xb4\x8f\x06\xeb\xd3\x1b\xd3\xe7\x2f\x67\x9b\x09\x31\xcc\xe2\xfc\x5d\xa5\x68\x38\x05\xc6\x99\x99\x67\x76\x1b\x95\xa0\xa4\x72\x83\x08\xdf\x37\xac\xca\x51\x2d\xee\xb1\x64\xda\xa0\x9a\x93\x3c\x8f\xe9\xec\x66\x75\xd6\xb3\x8c\xfb\xe7\xda\xb1\xc1\x8f\x9d\xc0\x9e\x3d\x03\x54\xca\x0e\xd2\x34\xf1\x00\x5b\x10\x37\x5c\xfc\xcd\x7d\x94\x47\xfd\xa2\x44\x23\xff\x44\xa5\x99\xe0\x97\x69\x92\xfc\x70\x76\xe1\x6c\x47\x8c\x56\x20\xe7\xc8\x7a\xcd\xde\xe5\xf9\x83\x38\xe5\x0a\x9b\xba\x84\xe1\x32\x59\x9a\x28\x34\x8d\xe2\xc0\x59\x95\xb6\xe9\xbf\x01\x00\x00\xff\xff\x10\x11\xf5\x38\x65\x0c\x00\x00")

func typesGoTemplBytes() ([]byte, error) {
	return bindataRead(
		_typesGoTempl,
		"types.go.templ",
	)
}

func typesGoTempl() (*asset, error) {
	bytes, err := typesGoTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "types.go.templ", size: 3173, mode: os.FileMode(436), modTime: time.Unix(1538540481, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
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
	"cft.go.templ":                cftGoTempl,
	"controller.go.templ":         controllerGoTempl,
	"template_functions.go.templ": template_functionsGoTempl,
	"types.go.templ":              typesGoTempl,
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
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"cft.go.templ":                &bintree{cftGoTempl, map[string]*bintree{}},
	"controller.go.templ":         &bintree{controllerGoTempl, map[string]*bintree{}},
	"template_functions.go.templ": &bintree{template_functionsGoTempl, map[string]*bintree{}},
	"types.go.templ":              &bintree{typesGoTempl, map[string]*bintree{}},
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
