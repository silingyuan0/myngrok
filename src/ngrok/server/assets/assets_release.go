// Code generated by go-bindata.
// sources:
// assets/server/tls/snakeoil.crt
// assets/server/tls/snakeoil.key
// DO NOT EDIT!

// +build release

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _assetsServerTlsSnakeoilCrt = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x93\xbb\xb2\xaa\x48\x14\x86\x73\x9e\x62\x72\x6a\x0a\x04\x41\x09\x4e\xd0\x4d\x37\x97\x0d\x0d\x72\x13\x21\xdb\x62\x73\x57\x64\x0b\x34\xf2\xf4\xe3\x39\xc1\x24\x67\x45\xab\xbe\x5a\xc1\xaa\xfa\xff\xef\xdf\xdf\x03\xb1\x69\x7b\xff\xe8\x38\x8c\x6d\xc3\xd6\x41\x8c\xff\x50\x8e\xd8\xb6\xfe\x6a\x75\x1d\xe4\x83\xae\x07\xba\x55\xa9\xd1\xb0\x82\x95\x3f\xb4\xc0\x83\x55\x37\xd6\x5d\x63\x6a\x4c\x84\x20\x48\x0c\x80\xc0\x95\x84\x1d\x33\x58\x86\xce\x41\x80\x10\x84\x22\x77\x4b\x95\x39\x4f\xd9\x5c\x58\xa1\x7c\x95\xc1\x4c\xad\x5e\x25\x70\x7f\x41\x31\x96\x09\x22\xcc\x43\x40\x21\x6d\xb2\x79\xbb\xe1\xc3\x08\x23\x31\x5e\x49\xfc\x3f\x63\x9c\xb9\xe1\x9c\xc0\xc2\x04\xbb\x04\x03\xc6\x82\x9b\x75\x9e\xaf\xf7\xf3\xcb\x7d\x40\xf1\x26\x69\xcc\x7d\xd4\x0a\xfd\xfd\x61\x64\x33\x14\x64\x5f\xce\x90\xdb\xf5\x52\x78\x20\xc0\x10\x06\x00\x71\x55\x85\x4f\x00\x7d\x0e\x82\x41\xff\xec\x10\x90\xf7\x6c\xce\xc2\xd7\x12\x82\xee\xea\x9e\x3a\x33\xa9\x21\x9f\xe3\xad\xd0\x66\xb3\x1e\x7a\xe4\xc7\x9d\xf3\xbe\xce\xea\x98\x87\xcc\x6c\x8e\x05\x37\x2e\xa1\x1b\x18\xcf\x61\x1e\x6f\x95\x2b\xa0\xb4\x3f\xdb\x23\x89\x0b\xa9\xb2\x5c\xc5\x86\x89\x18\x8b\xda\xbc\x04\x65\x41\x0d\xf9\x9e\x85\x63\x76\xe8\xaf\x65\x2f\x5c\x34\xa9\x09\x16\xd7\x61\x0a\x07\x48\x52\xd5\x2b\xf6\x57\x27\x2e\x22\x4c\x75\xe3\x4e\x76\x47\x0d\xba\x6b\x84\x26\x34\xd4\x93\x6f\x22\xc7\x18\x0e\x81\x6d\x28\xbc\x52\x1b\xc0\x28\x56\xab\xf7\xbe\xf9\x90\xc0\xac\x5e\x63\x85\x8b\x9c\x47\x6d\xc2\xa2\x26\x6b\xba\x5a\x64\x4a\x81\xc6\x96\x5d\x1b\x65\x1b\x65\xb7\xb8\xba\x2b\x25\x5d\x97\xaa\x2a\xde\xaf\x76\xdc\xa1\x54\x63\xaf\xea\x84\x68\x7b\xdf\x2e\xc6\x8e\xb2\x7e\x06\x1c\x7f\xd2\xf5\x6c\x3d\x27\xe0\x50\x69\xc5\xb7\x71\x9d\x1a\xad\xba\xd1\x8d\x36\xfb\xec\xa9\x8e\x85\xd2\x39\x32\xd9\xcf\xd2\x7a\xcb\x61\xf2\x70\x2f\x1b\xb6\xde\xaf\x24\x17\x97\xc1\x34\x1a\xcf\x15\x31\x57\x0b\x73\x8f\x67\x9b\x84\x58\xa5\xdd\x9e\x58\xea\x7b\xd2\xe6\x09\x3f\xd4\x63\xf2\x1d\x81\x6e\x45\xed\x4b\x38\x9c\x89\x0e\x18\x06\x20\xfe\xbb\x1b\x1c\x00\xfe\x27\x01\x0c\x74\x8b\x6c\xf2\xe2\xfc\xb4\x72\xdd\x1f\x95\x1a\xf6\xf0\x71\x08\x5f\xe2\x42\x77\x6d\xf7\xb5\xaf\xb7\xfd\x8f\x48\x20\x02\x3b\xaf\x7f\xca\x9d\x78\x6f\x63\x0a\xd5\x75\xa8\x37\x0e\x1b\xba\xd9\x5c\x9f\xfa\x3d\xf1\xec\x63\x1a\xbb\xa5\xa0\x57\xa6\x30\x4a\x34\x3d\x4d\xcf\x8a\x12\xbc\x4b\x0b\xec\x09\x5e\x58\x6b\xeb\xf4\x78\xf3\x92\x06\xb2\x92\x4e\xa5\xe0\x07\xa1\xa4\x39\x39\xa7\xb8\xdd\xa1\xd0\x80\x5e\x59\x54\xb7\xd9\x7e\x16\x35\x81\x77\x3a\x37\xe3\x55\xe4\xd3\x68\xb8\x43\x75\xbb\x1b\xd3\xf9\x3a\xda\x38\xc7\x56\x7b\x3c\x5c\x34\xf7\x8b\x76\x09\x8d\x92\x8b\x10\x26\x33\x17\x47\x65\x6c\x6e\xeb\xbe\x9d\x2d\xfd\x92\x1b\xbb\xe0\xc5\xbe\xfa\x53\x3b\xfb\xb8\x3e\xf9\x7c\x7c\x34\x25\xbb\xf5\x29\x6f\x1f\x35\x97\x8f\x26\xef\xa1\x2a\x8e\xec\x9e\xe5\x3d\x7a\x1f\xa8\x17\x0c\x4f\x2e\x92\x88\x54\x7e\xfa\x25\x7f\xff\x94\xcd\xda\x95\x02\x78\xa5\x8d\x9e\x8e\xc8\x8f\xf2\xbe\x71\xf3\x19\xbb\xe4\x67\x59\xa6\x2a\x93\x86\x63\xa7\x2a\x69\x20\x02\xc3\xcd\x16\x43\xe4\xb1\x38\x40\xc4\xa9\xc7\x12\x5b\x77\xeb\x84\xd2\xa5\xf5\x0e\x32\x53\x9b\xeb\x94\xab\x69\xa1\xf9\xbe\x46\xc1\xaf\x5f\xdc\x1f\x69\xb1\x87\xfe\x16\xf9\xbf\x00\x00\x00\xff\xff\x8d\x69\x45\x8d\xe5\x03\x00\x00"

func assetsServerTlsSnakeoilCrtBytes() ([]byte, error) {
	return bindataRead(
		_assetsServerTlsSnakeoilCrt,
		"assets/server/tls/snakeoil.crt",
	)
}

func assetsServerTlsSnakeoilCrt() (*asset, error) {
	bytes, err := assetsServerTlsSnakeoilCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/server/tls/snakeoil.crt", size: 997, mode: os.FileMode(436), modTime: time.Unix(1488619537, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsServerTlsSnakeoilKey = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x95\xb9\xae\xab\xda\x12\x45\x73\xbe\xe2\xe4\xe8\x09\x63\xfa\xe0\x04\x0b\x58\xb4\x06\x4c\xdf\x64\x80\x69\x0d\x18\x4c\xcf\xd7\xbf\x7d\xf7\x4d\x6f\xa5\xa5\x2a\x0d\x4d\x4d\x69\xfc\xef\x9f\xe1\xa1\xac\x9a\x7f\x1c\x17\xfc\x79\x3a\x6a\x00\x3c\xf8\x47\x87\xf1\xef\x06\x31\x54\x15\x8e\x40\xe5\x01\xd0\x05\x60\x43\x70\xe9\x64\xca\xb0\x3d\x07\x05\x67\x9e\x51\x3b\x71\xa1\x42\xbd\x3d\xf3\xba\x93\xa9\xdc\xf8\x06\x65\xd9\x5f\x6d\x65\xa6\xb1\x53\x40\xfa\x38\xce\x09\xe1\xe0\x02\xc2\xbe\xa1\xc7\x3b\xd8\x58\x33\x0d\xfc\xe6\x73\x98\x57\x0a\xf2\xed\x5d\x49\x9e\xf3\xf4\xee\x34\xc7\x73\x07\x19\x15\x49\x2d\x8f\xed\x1a\x2c\x68\xc4\x95\x44\xaa\x09\xec\xfc\x15\xdf\x00\x39\x5c\x41\x81\x36\x03\x47\xf3\x50\x1d\x52\x0d\xe3\x2b\xba\x6e\x70\x93\x54\xeb\x66\x35\xf2\x8d\x8c\x8d\x4f\xd8\x2e\xa0\x8a\x06\xa6\x87\x3e\x08\x2e\x58\xf8\xf8\x97\x82\xbb\xd4\x28\xd2\xb3\x53\x91\xa9\x80\xb1\x74\xc2\x4b\xca\x60\x7e\xe2\xb1\x48\x08\x58\x68\x6a\xad\xc9\xf0\xb8\x25\x64\x1d\xc7\xc8\xa8\xc0\x5f\xfa\x69\x7d\x7d\x33\x23\x84\x13\xb0\x26\x6d\x65\x66\x1e\x05\x8c\x10\x92\x22\x89\xb0\xaa\xd6\x4a\x81\x2d\xae\x22\x7e\x7c\x82\xf5\xfe\xb8\xf3\x38\x63\xd2\x8f\xba\x19\xa6\xf1\xea\xed\x6f\x7e\x35\x4c\x26\xe1\x6f\xc9\xcd\x97\xd2\xb0\x4b\xfd\x70\x06\x17\xad\xe2\x50\xbd\x6d\xb6\xab\x20\x28\xed\xbb\x64\x75\x40\x6f\xa4\xbd\x6a\x2f\xbf\x3a\x71\xa7\x6f\x6e\xf9\x3d\x9c\x51\x15\x3c\x68\x59\x17\xb6\xf8\xbb\x2a\x02\x1b\xf0\xe0\xf3\x13\xb6\xe2\xce\xf9\xa0\x8d\xdb\xf5\x99\x52\xdf\xd5\x11\x9d\x7b\x09\x8b\x4c\x5b\x33\x9a\x24\xf1\xba\x6c\x87\xd2\x70\x13\x79\xc3\x17\x19\x2a\x4a\x15\x8d\x80\xd2\x74\xfd\x60\x0a\x5b\xac\x7c\x62\x60\x73\x76\x0b\x46\x9f\x11\xae\xea\xfb\x56\x9c\x5c\xb0\x91\x6e\xbb\x35\x9d\xa5\xe7\x9e\xda\xea\xb6\x6c\x87\x9c\x13\xf9\xfe\x81\xab\x91\x9f\x6d\x92\x37\x60\x0f\x99\x3d\xed\x78\xa3\xbe\xae\x0e\xa9\x2b\x7b\xb0\x8b\x61\xa1\xa2\xbe\x46\xfd\x64\xbc\x0e\xcd\x41\x3e\x42\x91\xbb\xfb\xea\x0c\x50\xe5\x0b\x58\x1d\x62\xb6\x4d\x6b\xc1\x29\xc5\x07\xaf\xe3\xf3\x5c\x8f\xd3\x4a\x47\xae\xa2\x58\xf1\x41\x3a\x25\x14\x52\xc9\xeb\x52\x37\xa4\x92\xea\x89\xbb\xe2\x25\x22\xde\xed\x0e\xfd\x3e\xd8\x6b\x8c\xb4\x5c\x06\x2d\x5f\xc7\x95\x44\xda\x32\x3e\x2a\x8b\x66\x74\xd3\x23\xcc\xef\x67\x1c\x47\x59\xd2\xec\x3d\xea\xf9\x36\xf1\x08\xa1\xec\x36\xe0\x7c\x2e\xfd\x8a\xf3\x12\xb9\xa8\x58\xec\xfd\x66\x58\x62\x26\x84\x81\x3e\x98\xc1\x0f\x93\xba\x44\x84\x39\xad\xc7\x6a\xe8\xc6\x4d\x58\x9b\x44\xc1\x88\xfd\x5c\x44\x74\xea\xb8\x35\x4d\x32\x14\x1d\x7c\x9a\x7e\xcb\x69\x94\x20\x86\x9f\x8f\xfb\x01\x85\x2a\x86\x00\x33\x79\x86\x5b\xd5\x09\xbd\x96\x1b\x16\x24\x59\x62\xc3\xfb\x95\x8f\x92\xfd\x20\xd5\x9d\x9f\xbb\xac\x87\x43\x3a\x85\xea\x5b\x4b\xd7\x81\x7b\x56\xac\xf6\x78\x20\x9b\x57\xef\xb5\x72\xbf\x28\xad\x68\xb1\xd7\x3e\xd7\x9a\x05\x9c\xef\xf6\x4c\x8a\xb9\x6c\x5b\xc9\x91\xda\x80\xa1\xb6\x29\xc0\x48\x6a\x96\x80\xed\xa3\xe9\xd7\x19\xc3\x9c\x0c\xd3\x30\x25\x0f\x0b\x43\x9a\x99\x7d\xd0\x95\x6f\xc6\x1e\x25\x14\x3a\x16\x76\x54\xa2\x08\xa7\x39\x83\xb7\xaf\x71\x5e\xbb\x6a\x67\x55\x97\x59\xda\xbc\xb8\x64\x37\x65\xf0\xf3\xfd\xdc\xf0\xd6\xff\x25\xbe\x6e\x5a\x8a\xc0\x46\x6a\x70\x0c\xee\xd8\x75\xa8\x4f\x7f\x7f\x8a\x53\xe3\x27\x8c\xfc\x0c\xf1\xe4\xbd\xd2\xb6\x29\xfa\x15\x99\x72\xf3\x2d\x65\x3d\xc6\x71\x86\xac\x50\x5a\x7b\xb5\xbf\x3e\xe7\x7f\xa3\xf6\x36\x71\x88\xc7\x61\x73\x33\xdb\x07\x39\x5b\xf9\xb5\x27\xd8\xc1\x6f\x8d\x71\xf3\xd9\x61\xd7\xd9\xdb\x20\x7f\xd2\x3e\x43\x8b\x39\xdc\xed\xe2\xb5\x07\xae\xe8\xa0\xf7\x28\xc4\x5c\xb5\x06\x7b\x48\x88\x63\x34\x23\x16\xa6\x75\x36\x8b\xe6\xbe\xc7\x26\x8e\x1c\x27\xf4\x3d\x1a\xdb\x47\x88\x33\xf3\x25\xe1\x84\x4d\xc8\xb4\xc7\xd5\x7d\xfe\x8b\xdc\xef\x4e\xe4\x4d\xf7\xae\xbc\xc1\x98\xa8\xd5\x69\x78\xbf\x5e\x88\x27\x55\x65\xff\x00\x4b\xf1\x73\x78\x33\x64\x6c\x1f\x9f\xd1\x58\xb7\x16\x29\xa3\xc7\xfa\x90\xe9\x67\xc6\x9e\xcb\xf6\x05\xb4\xea\x85\xb4\x74\x5b\xe8\xb5\x9d\xd6\xd2\x95\xe4\x8e\xb7\x06\xd0\x60\x48\xa8\x97\x31\xd1\x2f\xe8\x2a\xc8\xad\x9d\xd2\x17\x7c\x95\x2f\x71\x29\x48\xea\x91\x13\xdb\xcc\x7e\x94\xef\xfe\xa4\x5d\x45\x63\x88\x35\xec\xb4\x0d\x37\x84\xa7\xef\x57\x6c\x62\x76\xb9\x4d\xda\xb8\x8f\xe4\xcd\x1e\x3b\x4e\x62\xb4\xb3\xb0\x18\x3d\xaf\x74\x95\x47\x0b\xda\xbf\x29\x0f\x87\xbd\x37\x55\x74\x4a\x55\x9b\x1b\xfa\x1a\x7a\x6b\x1c\x43\xf5\x43\x8a\x6d\x36\x33\x14\x98\x24\xe7\x59\x22\xd3\xa3\x9a\xf9\xe6\x6a\xef\xe1\xce\xd3\x9a\xd2\x1b\xd2\x9a\xd2\xc6\x81\x77\x42\x7d\x2d\xf8\x47\x8d\xe1\x51\xf3\xb0\x43\x0b\xfd\x1a\x26\x23\xbf\x3f\x5f\x72\x2d\x6d\xfe\xd0\xc8\x97\x57\x7a\x73\x86\x8c\x17\xea\x3a\xf4\x1d\xbc\x33\xdf\x36\xc3\xcc\xa9\x6d\x75\x5b\xee\x2c\xf8\x90\x2c\x95\x82\xc9\x1a\x92\x83\xb2\xfc\x77\xda\x81\x9e\x8c\x9c\x80\x02\xe4\xa3\x7c\xe5\xcc\x44\x13\xc1\xfb\x7a\x02\x12\x79\xf1\x12\xf5\xfc\x2d\x06\x90\xcd\x9f\xfa\x11\x6b\x91\xc1\x04\x86\x69\xa5\x63\x92\xa6\x0d\xf1\x76\xd7\x48\xb0\xef\x71\xd9\x10\xc6\x30\x13\x51\xe1\x5a\x37\x2c\x5c\xbf\x35\xd5\xb5\x64\x1a\x23\x7c\xb2\xe0\xb9\x32\xc5\x1b\x6a\xd4\x6c\xe3\x39\x8f\xa5\x9b\xd1\xca\x6f\x50\x0f\x56\xa7\x28\xcc\x64\xbf\xed\xf8\x84\xf5\x21\xeb\xbf\x51\x8a\xc6\xf4\xf3\x2c\x6b\x6d\x96\x8b\x6d\x7f\xa3\xf1\x35\x20\xcc\x24\x41\x82\xa0\xc4\x53\x70\x83\x8f\x41\x6b\x93\xf1\x58\x57\xe2\x9e\x72\x71\xf0\x81\x4c\xbf\x4d\xf3\x41\x5e\x93\x49\xaf\xa8\x68\x51\xfc\xc5\x3a\x4a\x4b\xbe\xec\xbf\x7f\x91\x5f\xad\x40\x53\xfc\x6f\xdd\xfc\x3f\x00\x00\xff\xff\x93\x5c\x16\x24\x8f\x06\x00\x00"

func assetsServerTlsSnakeoilKeyBytes() ([]byte, error) {
	return bindataRead(
		_assetsServerTlsSnakeoilKey,
		"assets/server/tls/snakeoil.key",
	)
}

func assetsServerTlsSnakeoilKey() (*asset, error) {
	bytes, err := assetsServerTlsSnakeoilKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/server/tls/snakeoil.key", size: 1679, mode: os.FileMode(436), modTime: time.Unix(1488619537, 0)}
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
	"assets/server/tls/snakeoil.crt": assetsServerTlsSnakeoilCrt,
	"assets/server/tls/snakeoil.key": assetsServerTlsSnakeoilKey,
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
	"assets": &bintree{nil, map[string]*bintree{
		"server": &bintree{nil, map[string]*bintree{
			"tls": &bintree{nil, map[string]*bintree{
				"snakeoil.crt": &bintree{assetsServerTlsSnakeoilCrt, map[string]*bintree{}},
				"snakeoil.key": &bintree{assetsServerTlsSnakeoilKey, map[string]*bintree{}},
			}},
		}},
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
