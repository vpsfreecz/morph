// Code generated by go-bindata.
// sources:
// data/eval-machines.nix
// data/options.nix
// DO NOT EDIT!

package assets

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

var _dataEvalMachinesNix = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\x4b\x6f\xdb\x38\x10\xbe\xeb\x57\x0c\xd4\x00\xb6\x90\x46\xe9\xee\x6d\x6d\xf8\xd0\x2d\xf6\x05\xf4\x11\xb4\x45\x2f\x41\xb0\xa0\xa5\x91\x45\x98\x22\x55\x92\x52\xe2\x35\xfc\xdf\x17\xc3\x87\x28\xbb\x39\x34\x3e\x24\xd2\x88\x1c\xce\x7c\xf3\xf1\x9b\x79\x05\xef\x54\xd7\x0b\xb4\x28\x0e\x60\xac\xe6\x7d\x8f\x35\xd4\xea\x51\xc2\x88\xda\x70\x25\x41\x35\x20\xf9\x93\xea\xcd\x02\x70\x64\x62\x60\x56\xe9\xec\x08\x12\xed\xa3\xd2\xfb\x3f\x9e\x7a\x0d\xa7\x55\x96\x09\xb4\x19\x44\x2b\x6c\x80\x77\xbd\xd2\x76\xbe\x6c\x9d\x01\xf4\xfb\x9d\x01\x00\xd8\xc4\x0f\x65\xfc\x4f\x5f\x68\x85\xe0\x5b\xf0\x2b\xc8\x52\x0a\xbe\x5d\x67\x5c\x66\x00\x8f\xdc\xb6\x10\x57\xb9\x17\xf7\x2d\xd3\x58\xc1\x31\x03\xf0\xa9\x0c\x16\xc1\xb6\x08\x35\x36\x5c\x72\xcb\x95\x34\x94\x01\x99\x3a\x56\xb5\x5c\xa2\x29\x29\x4e\x55\xa3\x81\x4d\x46\x27\x09\x6e\xec\x57\xf5\xd6\x5a\x6d\x60\xd9\xb1\x9e\xfe\xb8\x95\x1f\x59\x87\x2b\xb7\x04\xc0\xa7\xe7\x7f\xaf\xe0\x2f\xb4\xce\x65\xa5\x64\xc3\x77\x83\x66\x36\x20\x65\x5b\x6e\xe2\x41\xd0\x68\xd5\x01\xb2\xaa\x8d\xc9\xce\x3c\xe0\x53\xaf\xd1\x10\xc0\xaf\x81\x59\xeb\x76\xec\xe0\xdf\x86\x0b\xa4\x77\xcd\xb7\x83\x45\x03\x46\xb9\x73\x3e\xf2\xa7\x4f\x5f\xa0\x53\xf5\x20\x70\xe6\xc4\x1c\x8c\xc5\x0e\x2a\x26\x61\xc7\x47\x04\x83\xd2\xf0\xad\x40\x40\xad\x95\x86\x0e\x8d\x61\x3b\x9f\xb0\xff\x79\x0f\x06\x36\x70\x0f\xc7\x50\x22\xff\x16\xcb\x70\x75\x9c\x25\x7f\x82\x87\x35\x9c\x68\xa5\x6c\x51\x73\x0b\xcb\xb0\xac\xf0\x91\xd2\xc7\x87\x75\xf0\xee\x8a\x44\xbf\x23\x48\xd6\x21\x6c\x60\xe6\x69\x3d\x85\x40\x14\xc2\xc4\x8f\xfc\xea\x68\xd5\x17\xab\x29\x7b\x57\xf0\x9e\xd9\xf6\x74\x4b\x8c\x33\xb7\x82\x6f\x6f\x89\x73\x37\x1e\xe7\x52\xf2\xa7\xdc\xd5\x3a\xfe\x62\x5c\x91\x16\x3f\xe4\x39\xb3\x25\xeb\xf5\xf5\x99\x99\xa0\xd8\xe3\x01\x36\x90\xd7\xd8\x0b\x75\xb8\x31\x76\x68\x9a\x7c\x7d\xb6\xca\x1d\x37\x03\xac\xbc\x55\xbd\xa3\x17\x45\x95\x50\x48\xbf\x57\x70\xa7\xd5\xc8\x6b\x04\x46\x6c\x64\x83\xb0\xd0\x2a\x63\x1d\x38\x4c\xd6\xe0\x0f\xeb\x50\x5a\xb0\x4c\xef\xd0\x02\x7e\x1f\x98\x78\xc6\x91\xf5\x2c\x98\x78\xe1\x01\x3e\xa7\x35\x70\xe9\x5f\x55\x8d\xa2\xfc\xc1\x49\x28\x1c\x97\xbb\x92\x82\xf8\x18\x2a\xb4\xff\x34\xa2\xd6\x14\xe4\x6f\x6f\xde\x3c\x5f\xb0\xf8\x4b\xe1\x96\x3e\xdc\xbf\x95\xb1\x2f\xf2\x71\x3a\x7b\x3f\x83\x0c\x9f\xac\x66\x6f\xf5\x8e\xb0\x4d\x74\xf3\xb7\x74\xfd\x1c\x9f\xe0\x94\xb6\x4f\x8f\xfe\x80\x02\x96\x84\x14\x2d\x33\xb0\xd4\xd8\xa9\x11\xfd\xed\x8e\xea\x74\x0f\x79\x78\xcc\xa9\xe6\xae\x36\x26\x87\x5c\xa3\x51\x83\xae\xd0\x3f\x7f\x1f\xb8\xc6\x1c\x72\x47\xf5\x1c\x1e\x8a\xa2\x58\x67\x59\x36\x47\xe2\x1f\xd9\xa8\x0f\x8e\x55\x14\x77\x76\x0e\xd3\x64\x82\x98\x40\xfa\x36\x87\x87\xe9\xdd\x40\x36\xca\x9d\xe9\x44\xe4\x88\xc2\x30\xf0\xda\xdb\x5c\xa2\xa7\x75\xe6\xd4\xee\xae\x65\x06\xe1\x97\x55\x54\x65\x04\x25\xc5\x21\x88\xdf\x14\x43\xd2\x12\xe2\x04\x97\x8d\x8a\xaa\x37\x09\x5a\x80\x62\x91\x44\x39\x46\xe0\xf0\x77\x76\x7a\xf0\x56\x2e\xc1\x2b\xae\x7b\x8b\x9a\x3a\xdd\xb4\x46\xf0\x1e\x3a\xd6\x07\xc0\x5d\x01\x97\x72\x05\xe3\x62\x45\x47\xc2\x08\x1b\x30\x95\x1e\xb6\x9f\xdc\xed\xf9\xe6\xd4\x60\x5c\xac\x93\x7c\xc0\x5c\x6f\xc6\x32\x5c\xfd\x94\x52\x01\x33\xfa\x19\xac\x34\x5a\x03\x2d\x32\x61\xdb\x77\x2d\x56\x7b\x03\x23\xd5\x73\xce\xae\x00\xbf\x3c\xb3\x91\xc2\x7c\x46\x81\x84\xe2\x06\xa6\x83\xbc\xa6\x96\x67\x9f\x95\x8e\x3c\xfa\x32\x34\x0d\x7f\x7a\x7e\xf5\x37\xdf\x30\x7f\x62\x49\x31\xe3\x6e\x78\x2a\xd6\x67\x80\xbe\xe7\xee\x6e\xf9\x76\xb4\xc7\xc3\x0a\x76\x68\x09\x53\xa7\x54\x11\xf5\x33\xa6\x4f\xc6\xe0\x3d\xb5\xe2\x58\xe0\xd8\x6b\x29\x9f\xe3\x33\x5c\xfa\x75\x05\xdb\x81\x8b\x1a\xaa\x30\x13\x4c\xda\x72\xd6\xe7\x1c\x93\x52\xe1\x83\xde\x1b\x9a\x02\x02\xb1\x12\x71\x1a\x2e\x2c\xea\xd0\x5b\x89\x06\x2b\x40\x81\x1d\x48\xbf\xa7\x08\xcc\x8a\xd5\xd7\x83\x7c\xa7\xba\x8e\xd4\x31\xef\x94\xee\xdb\x7c\xea\x29\xbd\xc6\x06\xf5\x7b\x55\x31\xf1\xbb\x0b\x72\x03\x56\x0f\x24\x04\x61\xc9\x62\x91\xda\xdc\xbe\xe6\x1a\x6e\x7a\xb8\x52\x43\x6a\xdb\xb3\x4e\xe3\x60\x73\xdc\xf3\x2d\xff\x2c\xc0\x99\x23\x00\x21\xe1\xc6\xc0\xd5\xf1\xb2\x9c\x0e\xa8\xd2\xaa\x5e\xe0\x88\xe2\xe4\x4e\xba\xbd\x3a\xca\xd3\x8b\xf7\x96\xb5\x1e\xef\xa8\xe9\x25\x1f\x64\x9a\xfc\x2c\x16\x01\xa5\x45\x51\xa4\x5c\x43\xe1\xfe\x1c\x64\xe5\x66\x0f\x89\x58\x63\x4d\x9d\xa2\x62\xa2\x1a\x04\x0b\x73\x90\x1f\xdd\x92\xc6\x94\xf0\x95\x46\x14\xd3\xaa\x41\xd4\xe0\xd8\x80\x23\x4a\x78\x6c\x51\xa6\x55\xce\x37\xd3\x08\x52\xd1\x1d\xb3\x70\x40\xfb\x9a\x86\x91\x47\x84\x86\xed\x71\xa6\x59\x56\xc1\x16\x81\xd1\xd4\x61\x55\xd2\x22\x3a\x3b\x48\xe8\xac\x6f\xa9\xc6\x79\x9e\x05\x16\x29\x99\xe6\x21\x53\x3a\x8d\x1d\xba\xee\x10\x1a\x42\xe3\x79\x69\xb9\xa4\x79\xf0\x72\x58\x63\x2b\x1a\x04\x4b\xe2\x93\xab\xe8\x1d\xe3\x1a\x18\x34\x4c\x18\x2c\x60\x39\xed\x9c\xf5\x84\xc9\xd6\x04\xf8\xdc\x41\x8d\x17\x78\xa0\xab\x46\x03\x57\x6f\x68\x6a\xa5\x00\x84\x3f\x62\x90\xfc\xfb\x80\xb0\xa4\xe7\x46\x30\x6b\x51\xfa\x18\x76\x68\x3f\xfb\x5c\x0d\x88\xe4\x64\xb2\x51\x0a\x97\xa2\xeb\xd2\x77\xd3\xf3\x34\x0b\x35\x51\x16\xf4\xb4\x71\x62\x01\x6f\x12\x06\xdc\xa4\xb2\x27\x37\xb6\x45\x39\x23\xdf\x72\x39\xfb\xb6\x4c\x70\x26\x6b\x51\x94\xb1\x42\x4a\xc3\xfd\x43\x31\xed\x46\x61\x70\xee\x2a\xed\xb9\xd8\x11\x5b\x42\x58\x7c\x0f\x0d\x3c\xc0\xf5\x35\x5c\x82\x12\x13\x72\xc8\x50\x3b\xfd\xaa\xa6\xda\xfe\x3c\x30\xd3\x39\x2f\x04\x23\xd2\xe4\x48\xd3\x26\x3b\xe5\xb0\x49\xdb\x27\x49\x68\x68\x98\x7d\x11\x5f\xe6\x48\x66\x09\x38\xb8\x7f\x38\x67\x51\xcc\xd4\x78\x16\xfd\xc7\xa3\xdc\x5c\x72\xaa\x52\xb2\x62\xf6\x03\xeb\xe7\x10\x2d\x2f\xd8\xd8\x18\x47\xd3\x53\xf6\x7f\x00\x00\x00\xff\xff\xfa\xb3\xcc\x11\xbb\x0d\x00\x00")

func dataEvalMachinesNixBytes() ([]byte, error) {
	return bindataRead(
		_dataEvalMachinesNix,
		"data/eval-machines.nix",
	)
}

func dataEvalMachinesNix() (*asset, error) {
	bytes, err := dataEvalMachinesNixBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/eval-machines.nix", size: 3515, mode: os.FileMode(420), modTime: time.Unix(1523347553, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataOptionsNix = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x58\x4d\x6f\xdb\x38\x13\xbe\xeb\x57\x0c\xd4\x43\x12\xa0\x55\x9c\xb4\x2f\xde\xc5\x1a\x39\x2c\xba\xe8\xa6\x40\xd1\x14\x8d\x77\x2f\x45\x0f\x34\x35\xb6\x08\x53\x1c\x2d\x39\xb2\x63\x18\xf9\xef\x0b\x4a\xb2\x2c\xc9\x92\xe2\x24\x97\x05\xd6\x17\xcb\xa4\x66\xe6\x99\x0f\x3e\x33\xf4\x0e\x24\x99\x85\x5a\xbe\x05\xad\xe6\x6f\x21\x5b\x2d\xdd\x5b\x88\xa2\x08\x1e\x7f\x0d\x82\x8d\xe2\xc4\xaf\x4f\xeb\xa7\x88\xb7\x19\xba\x69\x10\x68\xe4\x20\xa0\x8d\x41\x7b\x97\xb1\x22\xe3\x66\xdb\x0c\xe1\x06\x5c\x3e\x4f\x29\xce\x35\xc2\xf9\xae\xd2\x03\xbb\x00\x00\x80\xca\xf7\xe0\xa6\xfa\xed\x3f\x4b\x4b\x79\x06\x37\x90\xae\x4a\x2d\x8d\x2d\xff\xe1\x4a\x27\xdb\x69\x6b\x3d\x46\x27\xad\x2a\x25\x6e\x20\xfc\xa3\xd0\xc2\x89\x60\xd8\x28\xad\x81\x36\x06\x38\x41\x70\x28\x2d\x72\x14\x76\x85\x17\x22\xd7\xec\x05\x2d\x11\x37\x76\x1f\xa7\x41\xfd\x9c\x3b\xb4\xaf\x06\xf6\xa7\x57\xb2\x49\xe8\xd5\xb0\xaa\xaf\xc7\x8b\x69\x10\xac\x70\x7b\x5a\xc8\xbb\x01\x8f\xd1\xb1\x32\xa2\x02\x77\xe4\xd9\xb1\x4f\x1d\x6f\xbe\x63\x4a\x8c\x90\x09\x4e\xc2\x1a\x52\xf1\xed\x28\xb7\x12\x5f\xa4\xf4\x0b\x49\xa1\xfb\x74\x16\xb5\xd5\xa7\xf2\x10\xa8\xdd\xe3\xb4\x6d\xa6\x5b\x8f\xfd\x36\xcf\xce\xea\xd0\xde\x15\x46\x68\xd1\xcc\x4b\xb5\x79\x76\xd6\xc6\x93\xa1\x4d\x95\x73\x55\x44\x47\x50\x85\x93\x0f\x93\x49\x38\x3d\x35\x00\xdf\x1a\x7a\xf1\x21\xb3\xe8\x1c\xc6\x20\x1c\x90\x64\xa1\xa3\x43\x54\x0e\x15\xb0\xf6\x96\x4e\xa9\x81\x4e\x11\x14\x9a\x98\xf5\xa9\x89\xaa\x3d\xfa\xf0\xfe\x7a\x32\x49\xc3\x29\xbc\x81\xf7\x13\x88\xc5\xd6\xf5\xfb\x32\x9b\x7d\x81\x05\xd9\x2a\x92\xc0\xb4\x42\xe3\x8a\x15\x4e\x94\x83\x84\x1c\x47\x9d\x34\x4b\x15\xdb\xde\x80\x56\x80\xb4\x72\x7c\xb7\xe8\xc7\xf5\xe3\xe7\x7e\x0d\x1f\x44\x9a\x69\xff\xfe\x8f\xf0\xea\xff\xd7\xd1\xf5\x24\xba\xba\x8a\xae\xae\x2f\xdf\x5f\x87\x3f\x07\x02\xff\xf9\xdb\xfa\x03\x7c\xfc\xfc\xfb\x77\x98\x6b\x92\xab\x92\x3f\xa4\x30\xa0\x69\xa9\x0c\xe4\x4e\x99\xe5\x73\x3c\xc9\x48\x2b\xa9\xf0\xc5\xce\x84\xd5\x73\xd8\xeb\xd5\xea\x17\x17\x42\x28\xc9\xb0\x25\xfd\x2e\xd3\xc2\xe0\xa0\x67\x7f\x15\x2a\x85\x94\xe8\xdc\x01\x16\x13\x88\x2c\xd3\xdb\x71\x2f\x1a\x24\xf1\x49\xe9\xd1\x33\x5d\x97\xdc\xee\x40\x9b\x7d\x24\x0f\xc5\xf1\x7e\x09\x99\xd6\x05\x78\xb9\x16\xf6\xb2\xcc\x85\xbb\x2c\xca\x3f\x42\xb3\x3e\x62\xd0\x56\x14\x3e\xe5\xba\x24\x16\x38\x57\x46\xea\x3c\xf6\xf9\x5c\x28\x8d\x46\xa4\x78\xb1\x3f\xf4\xb6\x24\x35\x34\xeb\x77\x7e\x0f\x36\x89\x92\x49\xc9\xd7\x09\xe9\x18\xca\x58\x4a\x8b\x31\x1a\x56\x42\xbb\x36\x6f\x37\x5b\xc6\x08\x67\x9d\xc4\x52\x5d\xaf\x77\xfb\x26\x54\x35\x84\xba\x5b\xee\x7f\x3f\x3e\xdd\x7c\x8e\x9b\x62\xe1\xa5\xaf\x23\xa1\x8c\x8f\xc8\x91\x83\x40\xa6\x19\x99\x66\x91\xf4\xf9\xfc\x04\x2f\x9e\x9c\xe2\x16\x6b\xf6\x3a\x74\x1a\x53\xd6\x18\x8f\x1e\xeb\x87\xc1\x46\xb2\x67\xd9\x92\x66\x13\x14\x9a\x93\x8f\x09\xca\xd5\x33\x3b\xad\x4c\xe3\xa7\x59\x40\xa6\xf1\x6d\xdb\xc2\x18\xc3\x75\xdb\xa6\x72\xec\x4b\x58\x52\x9a\x0a\x13\x43\x89\x15\xa4\x57\xe5\x9a\x3d\x03\x20\x61\xee\x1d\xb1\xda\x68\xfc\x5b\xaf\x87\x73\x3b\x9b\x7d\x1b\xc6\xb2\xef\x5f\x3d\xb6\xe0\xa6\xc0\xe3\xa2\x67\x0c\x33\x0d\x00\x3d\x65\xd7\x57\x72\x1d\xd4\xb7\x0d\xa0\xcd\xbd\x4e\xfc\xc8\xf1\x48\xfc\x4c\xae\xf5\x9d\x1d\x69\xf0\xb7\x5e\xde\x73\x4e\x78\x1c\x50\x2f\xbc\x5f\x7d\x73\x58\x2e\x47\xf2\xc8\x20\x6f\xc8\xae\x94\x59\x46\x1e\xc4\x57\x91\x62\x0b\x99\x93\x09\xa6\x2f\x1b\xbb\xee\x0b\xd1\x1e\x44\xa1\x4f\x4e\x3b\x00\x19\xd9\xb1\x00\x28\xc3\x43\xa3\x8d\x17\x34\x79\x3a\x47\xdb\xd1\x38\xd0\x0f\x2a\x8d\x7e\x7b\x28\x98\xbe\xc0\x2c\xfe\x9d\xa3\xe3\xe6\xd4\xd8\xee\x16\x9d\x04\xa2\x88\x71\x74\xc8\x10\xcc\xd6\x75\x1b\x73\xcb\xae\x21\x06\xe5\x9b\x71\x8a\x86\x31\xee\xb1\xda\xe2\x90\x82\x16\x15\x8d\xb1\xc0\x70\xd8\xee\x51\x92\x89\x1d\xcc\x91\x37\x88\xa6\x7d\x90\x9a\x26\xaf\x5b\x16\x59\xa5\x48\xf9\xcb\x32\x35\xab\x64\x95\xf1\x13\x8f\xb7\xde\x63\xee\x7f\x2d\x73\xca\x38\x94\xb9\xc5\xfb\xfb\x2f\x23\x26\xe7\x44\x7a\x68\xfe\x5a\x1a\xb2\x08\x5e\x1e\xad\x25\xdb\x67\x72\x21\xb4\xc3\x1e\x06\x39\xe6\xce\x7f\x31\x81\x0c\x74\x83\x2e\x83\x9c\x1f\xa6\xc3\x8b\x41\x8b\x1f\x2b\xbe\x67\x02\x9b\x1b\xdf\xfc\xbc\x54\xd8\x7c\xff\x98\x5b\xfe\x1b\x25\xb9\xaf\x0d\x65\x82\x46\xbe\xa3\x18\x33\x4d\x5b\x7f\x6c\xeb\xd4\x97\xd8\x84\x5d\x22\xdf\x8e\x93\x7b\x9d\xfc\xfa\xa6\x5b\xce\xa0\xa7\xde\x4b\x0f\x13\xfc\x0e\x32\xe1\xdc\x86\x6c\x1c\x31\x3e\x14\x34\xb5\x20\x9a\x0b\xdb\x9c\xe3\x3a\x74\xd4\xbe\xe9\x3f\x79\x99\xfd\xcd\x8b\x21\xc3\x26\x41\x8b\x80\x42\x26\x85\x26\x35\xcf\x19\x2b\xa9\x39\x3a\x10\x5e\xaf\x2f\xa0\xb9\x9f\x03\x33\x85\x31\xac\x95\x00\xe7\x92\x5a\x93\x32\x8e\x51\xc4\xe5\x8c\x6c\x29\x5f\x26\xc5\x44\xf8\x55\x3d\x80\xd4\xe4\x72\x8b\x70\xbe\x42\xcc\xfc\xfc\xa8\x18\x7c\x9a\xaa\x71\xda\xbf\xe2\x98\x2c\x46\x17\x95\xb6\xee\x3d\xba\x31\x55\x8d\x51\x72\x32\x34\x8b\x0c\x38\xdf\x3a\x8a\x65\xff\xcc\x6d\x71\x8d\xe9\x5e\xe8\x07\x48\xbb\xf8\x5e\x57\xeb\xcf\xf9\xd3\xa1\x7b\x1b\x7f\x1a\xaa\x70\x89\x92\x64\xb3\x6a\xf6\xde\x33\x93\xbf\x97\xed\x91\xfb\xc8\x8a\x2c\xb3\xa4\xb1\x79\xf9\xf4\x63\x80\xeb\xf9\x87\xa2\xac\xfe\xe0\x9f\x00\x00\x00\xff\xff\x37\xc9\x10\x9a\xcd\x13\x00\x00")

func dataOptionsNixBytes() ([]byte, error) {
	return bindataRead(
		_dataOptionsNix,
		"data/options.nix",
	)
}

func dataOptionsNix() (*asset, error) {
	bytes, err := dataOptionsNixBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/options.nix", size: 5069, mode: os.FileMode(420), modTime: time.Unix(1523448769, 0)}
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
	"data/eval-machines.nix": dataEvalMachinesNix,
	"data/options.nix": dataOptionsNix,
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
	"data": &bintree{nil, map[string]*bintree{
		"eval-machines.nix": &bintree{dataEvalMachinesNix, map[string]*bintree{}},
		"options.nix": &bintree{dataOptionsNix, map[string]*bintree{}},
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

