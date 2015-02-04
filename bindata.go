package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _include_buildpack_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\x5d\x53\xeb\x36\x13\xbe\xb6\x7f\xc5\xbe\x26\xc3\x01\x66\x4c\x5e\xce\xe9\x15\x87\x30\xe5\x23\xa7\xcd\x14\x08\x0d\xc9\x15\x65\x32\x8a\xbc\x89\x35\x28\x96\x2b\xc9\x04\x0a\xfc\xf7\xae\x1c\xc7\x1f\x24\x1c\xb8\xe8\x0c\x4c\x6c\x49\xbb\xfb\x3c\xbb\xcf\xae\xec\x4f\x32\x21\xa3\x94\xf1\xfb\x30\x7f\xda\xd9\x85\x67\xdf\x8b\x90\x4b\xa6\x11\x22\x34\xbc\x13\x9c\xba\x0d\x60\x09\xb0\x34\x95\x82\x33\x2b\x54\x02\x99\x11\xc9\x0c\x44\x62\x2c\x93\x12\x23\x28\xfd\x98\xc0\xf7\x30\x31\x99\xc6\x30\x65\x36\x36\xbe\x77\x7b\x0b\x41\x6b\x74\xd3\x1d\x04\x70\x77\x07\x2f\x2f\xa0\x59\x12\xa9\xb9\xf8\x07\xc3\x2c\x49\xb5\x78\x10\x12\x67\x18\xf9\x5e\x85\xc5\xa0\xcd\x52\x38\x86\x76\x84\x0f\xed\x24\x93\xb2\xbe\x89\x8f\xc8\x33\x8b\xf0\x42\xe1\x23\x4c\xac\xef\xa5\x5a\xf1\x29\x79\x09\xed\x53\x8a\xa6\xda\x78\xf5\x6b\xfc\x0a\xac\x9b\x18\xf6\x96\x5b\x15\x09\x98\x6a\x35\x87\xdf\x84\x85\xd1\xe0\x82\xa8\x47\xa0\x52\x47\x9b\x49\xe0\x6a\x3e\x17\xd6\x0a\x13\x07\x95\x9b\x4c\xcb\x4e\xd0\x3a\x08\x8a\x5d\x7a\xfe\x1a\x40\xc2\xe6\x48\x4f\xdf\xd6\x12\x22\xa6\x40\x39\xf9\x1f\x65\x85\xec\x5c\x52\xbe\x83\x8d\x31\xf1\x3d\x8f\x19\x62\x1e\x52\x8a\x89\x01\x97\x59\x84\xed\x2a\xaf\xfb\xf6\xd1\x12\xb5\x45\x4c\x44\x41\x23\x8b\x5c\xd4\x22\xe0\x77\x88\x14\x99\x7b\x6b\x6c\x57\x31\x82\xd6\xf2\x20\x61\xf1\x22\x95\x20\xfd\x68\xca\xb1\xa6\xa0\x53\xe1\x7b\x52\x71\xa2\x66\x99\x9e\xa1\x1d\x3b\x94\x84\xbb\xf4\x95\x2f\xb4\x5b\xcf\x8e\xcf\x61\xd8\xda\x99\x30\x83\xee\x19\x9c\xeb\xdd\xd7\x60\x45\xa8\x8c\x51\x67\x34\xa3\x1c\x72\x49\x11\x2b\x24\xb5\x30\x0e\x0e\x8f\xd6\xd7\x72\xab\x18\xf9\xbd\xca\x2c\x84\xe1\xdf\x99\x40\xdb\xe0\x40\x46\x61\x53\x1f\x28\x0d\x36\xe2\x85\x61\x84\x29\x31\x39\x78\x2f\xb2\x23\xae\xe7\x10\xea\x69\x73\xab\xbd\x3f\x73\x41\x1a\xe2\x91\xc2\xd8\x4d\xca\xb9\xa0\xf5\xcf\x75\x81\x34\x10\x3a\x28\xcd\xb4\xbe\x09\x93\xeb\x7e\x19\x67\x0b\x4e\x4b\x35\xe2\x63\x8a\xdc\xe6\x8d\x47\x9e\xe8\x4d\x69\x0b\x27\xd7\xd7\xe3\xf3\xde\x80\x2a\x45\x7d\xb9\xe2\x54\xec\xfd\xde\xbf\xec\x6e\xdc\x18\x74\xff\x1c\x75\x6f\x86\xe3\xde\x79\x27\xc8\xc3\x86\xad\xc1\xc9\xd5\x79\xff\xb2\x3a\x72\x33\x3c\x39\xfb\xa3\x13\x70\x8c\x98\x0e\x0f\x7e\xa1\x0d\x9e\x52\x92\xa0\x72\xd7\xde\x0f\x56\x44\x56\xee\x1d\xde\x6b\x8d\x69\x9e\x19\xad\xd2\x94\x92\x51\xb6\x36\x61\xce\x0c\xea\xb9\xa2\xa2\x85\xb1\x9a\x3b\x2d\x38\x84\xce\x4b\x7d\x02\x8c\xdd\xa9\xa0\x5e\x56\xf8\x7a\xbc\x7d\x40\x00\x62\xb5\x48\x20\x1c\x6c\x3a\x7f\xd8\x5c\x9a\x69\x95\xa5\x01\xfc\x45\x52\xa8\xf1\x2f\xde\x6b\x90\x8b\x15\xce\x48\x66\xeb\x67\x6a\xf5\x71\xcc\x46\x06\xa7\x99\x04\xaa\x8e\xa5\xa9\x67\xa0\x0d\x53\x64\xd4\x3d\x58\x55\xe3\x8c\x06\xc5\xf8\xac\x7f\x75\xd5\x3d\x1b\x8e\x87\xbd\xcb\x6e\x7f\x34\xec\x04\xdf\xfe\xbf\x74\x90\x97\xd2\x58\x4c\x61\x42\x9e\x17\x4c\x47\xc6\x75\x2e\x45\x10\x13\x21\x85\x7d\x5a\xf5\x50\x38\x6d\xe4\x19\x93\x87\x46\x37\x19\x95\x69\x8e\x6b\x47\x72\x2d\x37\x84\x54\xcc\xc8\xa5\x94\x0a\xd7\x09\xd9\x9d\x8e\x7a\x17\xe7\xd7\x54\xe0\x31\xe1\x6d\xb8\xb6\xc2\xd2\x5c\x09\x7e\xa0\xe5\xb1\x1b\xed\x3c\x33\x96\x66\x60\xe9\xd2\x75\x9d\x43\x80\x92\xa4\x88\xd1\x3b\x53\x62\x69\xe5\xce\x96\x9d\xd5\xb0\x70\xe9\xf0\x7a\x3f\x6e\x3a\x5f\xb6\xbe\xbc\x1d\x62\x70\x74\x74\xb4\x06\xd1\xff\xcc\x50\x5b\x81\xdd\xae\x4f\x84\x3a\xd8\x62\x14\xef\xd4\xa5\x02\x4d\x64\xed\x89\x48\xc8\xd8\xd2\x0a\xd4\x84\xb2\x1b\x94\xa3\x65\x0b\xa6\xca\x65\x3f\x46\xad\xee\xb3\xb0\x82\x35\xcf\xa4\x15\x60\x15\x14\xe6\x53\xa1\xdd\x54\x98\x52\xdf\xd2\x78\x20\x8d\xb8\x0a\x50\xff\xaf\x65\x6b\xb3\xa7\xf5\x0e\x28\x8b\xf4\x11\xa3\x4f\xf9\x7f\x97\x69\xde\x01\x9e\xb7\xbd\x0d\x1f\x54\xf9\x1d\xdc\x64\xed\x66\x6a\xed\x7e\x6b\xd6\xbe\x2e\xb7\xe2\xca\xa9\x06\x66\x67\xe7\x6d\x94\xbd\x5d\x77\x8e\x72\x5e\xbb\x97\x85\x13\xf1\x73\x65\x75\xfb\xeb\xdd\x6b\xb0\xba\xfd\x3e\x9d\x9d\x8f\x32\xf0\xd3\x14\xd4\x0f\x4d\x48\xc2\xf7\xee\xad\xb8\x53\x1d\x7d\xf7\xbf\xba\x10\xdf\xe7\x5f\xb4\x5b\xab\x01\xd9\x7d\x61\x15\x22\xc2\xa8\x12\x5e\x71\x76\x94\xb0\x09\xfd\x92\xce\x96\x46\xc0\x9a\xdd\x49\x6a\xb3\x70\x90\x03\x70\x33\x67\xd8\x3f\xef\x1f\x82\xc5\xa5\x14\x6d\x2c\x0c\xd0\x5f\x82\x1c\x8d\x61\xfa\x09\x28\xaf\xb4\xbe\x40\xe0\xee\xd3\x4e\x2e\xd8\x93\x81\x94\x3e\x40\xa0\x45\x23\x25\x07\x0c\x0b\x61\x63\x77\x03\x0b\x63\x32\xac\x58\xed\xb8\xbb\xec\xa4\x3a\xb7\xdb\x60\xd6\xc8\x79\xb0\xa1\xc9\xdc\xe0\xa3\xed\x37\x77\x08\x34\x87\x71\x50\x7a\xaf\xf2\xf0\x9f\x78\x5e\x16\xe8\x43\x57\x9a\x56\xe8\x3b\xe7\xe7\x20\x8f\x1b\xbb\xed\xfd\x95\x11\xe5\xdf\xc4\xf4\xc5\x08\xa1\x21\x69\xda\x99\x54\x13\x70\xcd\xec\x1e\xca\x6f\x8e\x6a\x86\xef\xf9\xde\xfc\xa1\x2e\xc4\xf6\x5e\xb5\x5b\xba\xca\xd6\x5d\xbd\xfa\xff\x06\x00\x00\xff\xff\x0c\xad\x20\xdf\xc0\x0b\x00\x00")

func include_buildpack_bash_bytes() ([]byte, error) {
	return bindata_read(
		_include_buildpack_bash,
		"include/buildpack.bash",
	)
}

func include_buildpack_bash() (*asset, error) {
	bytes, err := include_buildpack_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/buildpack.bash", size: 3008, mode: os.FileMode(420), modTime: time.Unix(1423042511, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _include_buildpacks_txt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x91\x49\x0e\x83\x30\x0c\x45\xf7\x9c\x82\x0b\x50\x44\x42\xa9\xda\xdb\x38\x83\x08\x10\x48\x94\xc1\x12\xb7\xef\x06\x15\x77\xe9\xbf\xc9\xea\xfd\x17\xdb\xae\x94\x98\x3f\x7d\x3f\x2f\xc5\x55\xf5\xd0\x61\xef\x9d\x4d\x61\xab\xd7\xd3\xa9\xba\x78\x13\x41\x6f\x9d\xf6\x61\xad\xc9\xb6\x57\x70\x92\x0d\x83\x9e\x43\x4b\xa3\xc4\x34\x00\x28\x56\x43\x02\xe3\x7f\xfa\x16\x07\xc1\xa1\x57\x40\x20\x7e\x94\xac\xdf\xef\xd5\x97\xe5\xa6\xb5\x31\x56\xcb\x91\xd3\x70\x04\x63\xd7\x7c\xfb\x27\x16\x1d\x5d\xa4\xdb\xc3\xe7\x8b\x45\x7b\x38\x29\x2d\x58\xb3\xc7\xb3\xb8\x70\x10\x37\x8b\x4e\x55\xfd\xb9\x07\xf1\xe6\xe0\x59\x83\x27\x77\xc3\x51\x36\xdf\x00\x00\x00\xff\xff\xb5\xa5\x85\x86\xb3\x02\x00\x00")

func include_buildpacks_txt_bytes() ([]byte, error) {
	return bindata_read(
		_include_buildpacks_txt,
		"include/buildpacks.txt",
	)
}

func include_buildpacks_txt() (*asset, error) {
	bytes, err := include_buildpacks_txt_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/buildpacks.txt", size: 691, mode: os.FileMode(420), modTime: time.Unix(1423042819, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _include_cedarish_txt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x33\xe2\x02\x04\x00\x00\xff\xff\x9f\x70\x98\xa2\x03\x00\x00\x00")

func include_cedarish_txt_bytes() ([]byte, error) {
	return bindata_read(
		_include_cedarish_txt,
		"include/cedarish.txt",
	)
}

func include_cedarish_txt() (*asset, error) {
	bytes, err := include_cedarish_txt_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/cedarish.txt", size: 3, mode: os.FileMode(420), modTime: time.Unix(1423042840, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _include_cmd_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\x41\x6f\x9c\x3c\x10\x3d\xc3\xaf\x98\xcf\x71\xa2\xe4\x80\xf8\x58\xf5\x44\xb4\x55\xa2\xb6\xb7\xb6\x97\x1c\xc3\x46\x72\xc0\x04\x14\xaf\x59\x61\xd8\xa6\x22\xfc\xf7\xce\x18\x9b\x40\xb2\xbd\xf4\x04\x9e\x19\xcf\x7b\xf3\xfc\xec\xb0\x90\xb9\x12\xad\x84\xe8\x16\xbe\xfc\xf8\x7a\x17\x86\xf9\xbe\x88\x54\x6d\xba\xcb\x2b\x18\xc2\xc0\xa7\x0b\x69\xf2\x2d\xfb\x8e\x71\x03\xe2\x28\x6a\x25\x1e\x95\x84\xbc\xd9\xef\x85\x2e\x0c\x7b\x2b\xd4\x66\xcb\x78\x82\x01\xdf\x27\x7a\x96\xbf\x0d\x30\xae\x0d\x83\x57\x30\xb2\x00\x66\x62\x5c\xa5\x71\xcc\xc2\xf1\x0d\xcf\xd6\xad\x41\xe7\x5e\x65\xd3\xc2\x33\xd4\x1a\xdb\x0c\xff\x11\xcd\xfb\x9b\xdd\xc8\xae\xa1\x68\xc2\x20\x90\x79\xd5\x60\xe2\x99\x48\x34\x5a\x22\xc8\x53\x2b\x0f\xc0\x1e\x08\xc4\x62\x36\x6d\xb7\x42\xd2\x0e\xe7\x5f\xda\x46\x47\x48\xdf\xf5\x94\x2f\x07\x5c\x9d\xd2\xeb\x9b\xcd\xa0\x62\x50\xf6\x3a\xef\xea\x46\x83\xa0\x95\xd3\x6d\x21\x5b\xa9\xed\xa8\x98\xc6\xef\xb0\x49\x23\x9e\x8c\x98\x56\x4d\x2e\x94\xd5\xc1\xa9\xa0\x89\x2e\xbf\x5c\x8c\x72\xb5\xe6\x5b\x6a\xb6\x54\x20\x62\x70\xf1\x19\xe2\x42\x1e\x63\xdd\x2b\x05\x17\x17\x93\xaa\xda\x8d\x15\x06\x76\x6e\x3a\x9e\x94\x0f\xc2\xc4\x67\xf8\x17\xc5\x23\xdb\x6d\x6d\xaf\xf5\x8c\xb3\x72\xef\x4e\xc8\x8d\xcb\x37\xd8\x55\x1e\x91\x31\x06\xa9\x10\x90\xd7\x4a\x91\x8c\x71\xfa\x66\x58\x67\x0d\xa2\x0d\xf0\x04\x32\x96\xf1\x9b\x0c\x85\x0f\x83\xd1\x39\x67\xc2\x83\xe9\xf8\x1d\xc5\xc4\x92\x4a\x66\x4e\x9e\xcc\x9b\x48\x98\xbc\x06\x53\xd5\x65\x07\x3e\x8c\x85\xab\xf8\xeb\x2b\x74\x6d\x2f\x7d\xda\x74\xa2\xeb\xcd\xf6\xff\x30\xa8\x4b\xf0\xaa\xce\x6e\xb5\x32\x3e\x70\x8c\x67\x7c\xa5\xe3\x35\x74\x95\xd4\x38\x04\x1f\x16\xfa\x61\x1d\xdb\x8d\xb8\xfb\x86\x74\x50\x06\x51\xa8\xed\xfd\x3d\x86\x28\x07\xbb\xdd\xbc\xd1\x9d\xd7\xcf\x06\x4c\x9f\x57\xde\x11\x29\xd8\x42\xca\x3b\x66\x1b\x3a\x5a\xe5\xbb\x10\xad\x8f\x4d\xf8\x65\xa9\x23\xd2\x75\xaa\xb8\xa2\xfd\x65\xed\x3c\x31\x5b\xe3\xf6\xc3\xb5\x4d\x6d\x21\xba\x0a\x31\xd7\xbe\x72\x8d\x9c\xb5\x82\x43\x5b\xeb\xae\x04\x06\x70\x1e\x6d\x3e\x19\x38\x37\x19\xda\xcc\x0d\xb5\x82\xff\x28\xc7\x44\xc7\xc2\x98\xfe\xf1\x14\x12\xd5\xcd\x50\x0b\xac\xf7\x68\xd3\xf6\xbf\x00\x62\x26\xf5\x15\x33\xe8\xe4\x70\xff\xf1\x6a\xbc\xd4\x1d\xf0\x49\xde\x90\x74\x72\x76\xaa\xa4\x3a\x9c\xba\xc4\x77\x55\xf3\xcb\x00\x65\x91\x39\x8e\xb1\x17\xf6\x22\xd3\x40\xa7\x6e\xb2\x68\x9f\xc8\x89\x64\x01\x7f\x6a\x14\x5a\x9c\x1b\x4d\xe6\x65\xa7\xb1\xad\x21\xe9\x07\x5f\x99\x33\x50\x02\x45\xc1\x1d\x53\xd9\xc2\xdb\x03\xb5\x89\xcf\x69\xd2\x98\x5e\xa9\x29\xa8\x31\x04\xeb\x35\xc4\x11\x5e\xa3\xc5\x76\xfb\xb8\x9c\x38\x19\xc7\x44\x47\x34\x97\x7b\x3b\x92\xd9\xb9\xee\x7e\xda\x97\x87\x44\x5a\x5c\x4b\x2f\x97\x55\x25\xfc\x13\x00\x00\xff\xff\xd3\x23\x50\x2b\x44\x06\x00\x00")

func include_cmd_bash_bytes() ([]byte, error) {
	return bindata_read(
		_include_cmd_bash,
		"include/cmd.bash",
	)
}

func include_cmd_bash() (*asset, error) {
	bytes, err := include_cmd_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/cmd.bash", size: 1604, mode: os.FileMode(420), modTime: time.Unix(1422884875, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _include_fn_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x51\x41\x6e\xc2\x30\x10\x3c\xc7\xaf\x18\xad\x22\x01\xaa\xa2\x08\xae\x34\x3d\x56\xea\x1b\x28\x07\xcb\xac\x89\xd5\xd4\x89\x6c\x03\xaa\x28\x7f\xef\x1a\xd2\x92\x43\x55\x55\xb9\x64\x77\x66\x67\x76\xd6\xca\xfa\x4a\x87\x7d\x9c\x2f\x70\x56\xc5\x8e\x4d\xa7\x03\x63\xc7\xd1\x34\xf4\xe2\xe3\xc0\x26\x41\xc3\x1e\xbc\x49\xae\xf7\xb3\x08\x21\x1f\xde\xd9\xa7\x48\xaa\xe8\x7a\xa3\xbb\xdc\xe9\x9c\xe7\xa6\x9c\xa7\x8f\x81\x51\x2e\xf1\x89\x7d\xe0\x01\xdf\x6a\x63\x59\x1d\x41\x53\x03\x12\xa0\x65\xbd\x43\xb5\x5c\xa8\x82\x4d\xdb\xa3\x62\x50\x79\x1e\x05\xeb\x1a\x35\xbd\x7a\xba\x64\xa2\x3e\xbd\xa1\x7a\x6e\x30\xab\x9b\xfa\x3c\x04\xe7\x13\xe8\x91\xca\x25\x3d\xd1\x65\x26\x78\x0a\xc8\x5c\xc8\xa7\x2e\x2a\xa7\xca\x16\xff\x4e\x95\xa1\xe0\x86\x5c\x51\x1e\xc8\x44\xf9\xe1\xa3\xe4\xa3\x5f\x82\x45\x33\x59\x9e\xd6\xb8\x6e\x5f\xe6\xfe\xe8\xee\xbc\xed\xff\x70\x8f\x13\x7b\xba\x73\xac\x6f\x72\x26\xc4\xb6\x3f\xc5\xfe\x10\x0c\x4b\xbd\xa2\xf1\x3a\x54\x5a\x8f\x72\x3e\xbe\x18\xa4\x5a\xfc\x40\xb8\x01\xd7\xcd\x26\x80\x2a\x9c\xc5\x66\x23\xa3\x77\x49\xc2\x76\xbb\x46\x6a\xd9\xab\xa2\xb8\x25\x13\x5d\x39\xa1\x76\x1d\x2a\x8f\x87\x95\xf4\x6f\xc3\xd6\x49\x9c\xaf\x00\x00\x00\xff\xff\x5c\xaf\x12\xe0\x23\x02\x00\x00")

func include_fn_bash_bytes() ([]byte, error) {
	return bindata_read(
		_include_fn_bash,
		"include/fn.bash",
	)
}

func include_fn_bash() (*asset, error) {
	bytes, err := include_fn_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/fn.bash", size: 547, mode: os.FileMode(420), modTime: time.Unix(1422884875, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _include_herokuish_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x56\x6f\x73\x13\xb7\x13\x7e\x6d\x7f\x8a\x45\x3f\x07\x27\xf9\x8d\x30\x50\xa6\x2f\x92\x71\xa7\x26\x84\xc2\x14\x82\x27\x21\x74\x3a\x84\xc9\xc8\xa7\xf5\x59\x93\x3b\xe9\x2a\xe9\xc0\x29\xf0\xdd\xbb\xd2\xfd\xf1\xd9\xe7\xe4\x4d\x4e\xbb\xcf\xae\x56\xbb\x8f\x1e\x79\xa8\x96\xf0\xf9\x33\xb0\xd1\xcb\xd9\xd5\x9b\xdb\x4f\xe7\x97\x57\x6f\x2f\x5e\x7f\x60\xc0\x33\x0f\xec\x05\x83\x2f\x5f\x4e\xc1\xaf\x50\x0f\x07\x98\xac\x0c\xb0\x47\x8f\xe0\x6f\x53\x5a\x70\xf7\xce\x63\x0e\x2f\x85\x5b\x81\x72\x60\x4a\x0f\x66\x09\x52\x78\x3c\x81\x4e\xae\x0f\x17\xac\x13\x39\xcf\x50\x38\x84\xb2\x48\xad\x90\x08\xde\x54\xf1\x2f\xc0\x58\x48\x2d\x52\xb0\x7d\x12\xf0\x6b\xe5\xe1\xf9\x70\xa9\x86\x43\x32\x4a\xa3\xb3\x7b\x10\x45\x71\x5b\x08\xbf\x9a\xb2\xd1\xf7\xd9\x7c\x7e\x3b\x9f\x7d\x7c\x73\xc2\x27\x64\xfe\xc9\x36\x28\xd4\x5f\x5b\xd4\xf9\xc5\xa7\x06\xe5\xf3\x62\x42\xae\x2e\x72\x51\xaa\x4c\xb6\xd8\x97\xd7\x6f\xdf\xbd\xea\xa2\xa3\xbb\x8b\x4f\x44\xb2\xc2\x16\x7f\x36\x3b\x7b\x73\xde\xc5\x47\x77\x2f\x7f\x21\x92\xbb\xed\x3d\xe6\xb3\xb3\x3f\x7b\xfb\x04\x98\xa3\xe0\xce\x6e\x28\x85\x55\x6e\x75\xfb\x15\xad\x53\x46\x53\xfc\xa1\x70\x0e\x3d\x4f\x84\x07\xa5\x93\xac\x94\x38\x69\x50\x4f\xfc\xda\x1f\x51\xb8\xc4\x24\x13\x96\xfa\xab\x0b\xab\xbe\xaa\x0c\x53\x94\xb7\xa5\x43\x4b\xe1\xd7\x57\xe7\x97\x6c\x3f\x22\xb5\xa6\x2c\x42\x85\x01\x33\xd1\x66\x61\xe4\x3d\xfd\x8b\xe6\x50\x15\x6a\x57\x5a\xe4\xe1\x1c\xee\xf0\x08\xbe\x0f\x07\xf9\x9d\x54\x16\x78\x01\x37\xc3\xc1\x80\x8d\x9a\xd9\xb0\x7a\xdd\x4c\xa1\x59\x6f\x7a\xdd\x58\x36\xdd\xdc\xc2\xb4\xfd\x62\xc3\x9f\xc3\x61\x67\xc3\xa6\x6e\x89\x2e\x99\xb2\xab\x95\xf9\xe6\x20\xb8\x81\x5a\xe2\x95\x4e\x1d\xb1\x86\x4e\xa4\xfd\x12\xd8\x01\xff\xe5\xb9\x83\xff\xc1\x81\xbb\xd1\x75\xf6\x86\x30\xd3\x4e\xa9\xc1\x5c\x14\x99\xa2\x86\x52\x83\xab\x6c\xb2\xa4\x1c\x29\xd8\x52\x7b\x95\x63\x1d\xdc\xf0\x68\xda\x39\x17\x99\xe7\x21\x80\x18\xbc\xa4\x2e\x3a\x58\x12\x83\x25\x2e\x95\x0e\xf1\x8b\x40\x72\x02\x2b\x6b\x74\x8e\xda\xd7\x89\x36\x24\x9b\x6e\xb5\x64\xc0\xfe\x32\xf6\x2e\x04\x52\x57\x31\xf1\xc6\xde\x37\x95\x44\x98\xab\xe3\x37\xa4\x9b\x6e\x35\x90\x32\x37\xcd\xab\x68\x0a\x99\xa9\x4e\xd5\xdd\xb8\x65\xde\x74\xb7\xd7\xd0\x9e\x45\x69\xe7\x45\x96\xa1\xdc\xd0\xd7\xc5\x51\xd4\x34\x7c\x68\x18\x50\xfb\x41\x68\x09\xae\x2c\x0a\x63\x3d\x25\x69\xac\x4a\x2f\x4d\x2b\x03\x2b\xb4\xe6\xae\x24\xd6\x36\x6e\xd2\x8b\xef\xb5\x56\x9c\x70\x89\xe1\x96\xd6\xd0\xc4\xe4\x54\xa0\x5a\x64\xd8\x5e\x08\x02\xef\xde\x8d\x7d\xf0\x4d\xf9\x27\xe4\xee\xdf\x9c\x8d\x3f\xdc\x1d\xf8\x41\x3c\x92\x30\x76\x93\x27\xc7\x55\x79\x37\x93\xc9\x64\x4c\xe6\xb5\xb0\x29\x31\xad\x66\x16\xc0\x01\x7f\xfe\xab\xab\x98\x15\xda\xe2\x95\xcf\xb0\x6a\x4a\x2c\x61\x34\xbe\xc1\xcf\xcf\xfe\xe0\xe1\xef\xb7\x31\x8c\x8e\x03\x48\x69\x49\x1c\xa8\x50\xdf\x56\xc4\x16\x08\xd7\x1c\x32\xa5\xf1\x14\xa4\xa1\x01\x35\x12\x1c\x4c\x0c\xa6\x53\xe0\xfc\x98\x84\x17\x7e\xfc\xd8\xb1\x4f\xa7\xc7\x1d\x41\x1e\x6c\x6d\x3a\x8e\x30\xb2\x62\xe6\x70\xd7\x09\xf1\x6f\xdc\xe4\x22\x37\x89\xeb\x80\xb4\x06\x43\x85\x5d\x35\xa8\xea\xa4\x7e\x95\x4a\xa6\x4a\x52\x44\x4f\x4d\x88\x30\xa3\xdf\xe3\xf9\x25\x7a\xe2\x2b\xef\xc7\xef\x53\xa0\x43\xe2\x96\x07\x9e\xc0\xc1\x35\x74\x34\x83\x84\x6b\xb0\x5f\x8e\x7a\x49\xfa\xda\x44\x35\x58\xe2\x9c\xc9\xd5\xbf\xb8\xa7\x8c\x70\x0f\x32\x08\xa1\x4a\x86\x0a\x0e\x2f\x67\x17\xaf\x3e\xbc\xff\xff\xb3\xa7\x4f\x9f\x1e\x85\x8d\x37\x00\x2d\x72\x9c\xb2\x92\x76\x8d\xe8\xa0\x7b\x03\x21\x65\xdc\x88\xe6\xf1\x4f\xa9\x90\x8a\xe7\x75\x4b\x22\x86\xd5\x5f\x21\x94\x45\x74\x58\xc5\x2b\xc7\xb9\x5b\x61\x96\xc1\x64\xa1\xf4\x64\x11\xde\xb8\xca\x2a\x95\x13\xc4\x4f\x49\x62\xea\xdc\x37\x63\x65\x6d\x27\xf9\x48\x90\x2f\x84\x0c\xb9\x6a\x9b\x36\x3c\x89\x6f\x22\x5f\x99\xd6\x58\x6e\xed\x5f\xd9\xd2\x7d\x36\x4c\x8c\x83\xf1\xb8\x5e\x56\xe5\x57\xdf\x31\x5b\x5f\xb4\x3b\x27\xd9\x3f\xbf\x0e\x60\xef\xbc\x36\x7e\x9a\x4a\x2e\x94\x6e\xa9\x04\x1c\x0d\x14\xaa\xc0\xa5\x50\xd9\x69\xc5\xe9\x8f\x97\xb3\xb3\xf3\xf0\xf3\x02\x1e\x3f\x86\x88\x59\x53\xc7\x93\x5c\x72\x5c\x07\xf9\x88\x82\xec\xb6\x2c\xf5\x7d\xdf\x82\x71\xed\x36\x97\x1d\xd8\x35\x09\x6f\xd0\xa0\x5a\xc8\xb6\x64\xac\x9b\xaa\xb5\xf3\xf8\xf5\x80\xaf\xce\xf2\x80\x37\x53\xce\xf7\x6a\x71\x59\x99\x02\x7b\x2f\xb4\x48\x31\xfc\x66\x69\xdf\x97\xe0\xd8\x29\x22\x98\xb8\xca\xc3\x77\xdf\x9e\xa2\x46\x4b\xc3\xef\x7b\xaa\xef\xde\xce\x85\x35\x49\x78\x8b\xaa\x26\xcc\xeb\x95\x8b\xed\xa0\x07\x2d\x14\x03\xa4\x90\x39\xad\x77\xea\x68\x22\x39\x9d\x76\xa7\x94\xd6\x85\x6b\x4c\xf6\x7b\x0a\x61\x83\xde\x90\x33\xbc\x7a\x6c\x74\x75\xfe\xee\x35\xa3\x01\x00\xb1\x6a\x12\x33\x1e\x0d\x06\xdb\x5b\x44\xfd\x38\x3d\x0d\x80\x90\xb7\xeb\x0f\xeb\x8e\x3b\x76\x9b\xfc\x3b\xf3\x8a\xce\x63\xb2\xc7\x8a\xe8\xf0\xac\xd2\x24\x32\xa3\x13\xc9\xf0\xe7\x7f\x01\x00\x00\xff\xff\x36\x14\xbb\x37\xd6\x0a\x00\x00")

func include_herokuish_bash_bytes() ([]byte, error) {
	return bindata_read(
		_include_herokuish_bash,
		"include/herokuish.bash",
	)
}

func include_herokuish_bash() (*asset, error) {
	bytes, err := include_herokuish_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/herokuish.bash", size: 2774, mode: os.FileMode(420), modTime: time.Unix(1422999670, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _include_procfile_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x94\x4b\x4f\xdc\x30\x10\x80\xcf\x9b\x5f\x31\x8a\x22\xb4\x54\x0d\x11\x3d\x16\x2d\xea\x01\x54\x2e\x7d\x88\xaa\x27\xba\x5a\x99\x78\xf2\xd0\x3a\x76\x64\x3b\x0b\x08\xf8\xef\x1d\x3b\xce\x12\xb3\xcb\x29\xb1\x3d\x8f\x6f\x9e\x49\xaf\x55\x59\xb5\x02\xf3\x9e\x69\x83\xcb\x53\x78\x4e\x16\x1c\x4b\xc1\x34\x02\x47\x53\xae\xd2\xef\x68\xa1\x54\x5d\xc7\x24\x07\x63\x75\x2b\x6b\xa8\x94\x06\x06\x4e\x15\x8d\x01\xfb\xd4\x23\x54\x5a\x75\xf0\x3b\x18\x4b\xdf\x6c\xb8\xc7\x55\x9a\x9d\xd3\x55\xc9\x2c\xa4\x19\xeb\xfb\x4d\xcf\x6c\x53\xec\x85\xe1\x05\x9e\x58\x27\xf2\x1a\xdd\xbb\x53\x48\x93\xd7\xe4\x8d\xcc\x58\xa6\xed\x31\xb2\xdb\x41\xc6\x10\x13\x66\x04\x03\xb6\xd1\x6a\xa8\x1b\xc0\x47\x2c\x8f\x93\xed\x5d\x39\x11\x62\x58\xc6\x59\x99\xa8\x4e\x63\x2e\x27\xfc\x11\x16\x33\x30\xc8\x5e\xb7\x3b\x12\xac\x91\xc3\x60\x50\xc3\x43\x6b\x1b\xb8\x41\xad\xb6\x43\x2e\xda\x2d\x02\xca\x1d\x79\xbf\xbb\x23\x07\x7f\xff\x5c\xdf\xa6\xb0\x5e\xc3\xcb\x0b\x99\xb1\x58\xda\x7c\x6e\x60\xc6\x68\xd0\x0e\x7d\xde\xa8\x0e\x67\x97\x42\x31\x9e\x93\xb9\xf7\x57\x74\x72\x07\x4a\x3e\x9f\xe5\x9e\x9c\x46\x74\xc5\x7d\x2b\x8b\x7b\x66\x1a\xc8\x7d\xf8\xb8\x63\x02\xb0\x6c\x14\x64\xdf\xde\x05\xed\x12\x61\xc6\xa8\x6d\x6b\x29\xbb\xe9\x55\x6b\x4a\xb5\x43\xdf\x18\xf3\x6a\x18\xf2\xd2\x56\x40\xd1\xe5\xd5\xf1\xba\xaf\xd7\x17\x54\x1c\x94\xc9\x62\x21\x54\x49\x2e\xbd\x1a\x9d\xfc\x97\x6a\xb3\x74\x2d\x73\xa8\x39\x35\xcc\x16\x9f\x0c\xfd\x3f\x32\x5d\x1b\x8f\x4b\xac\x8b\x85\xe7\x4e\xf7\xb2\xa1\x36\x81\x09\xf2\x4b\xc8\x9e\xfd\x6f\x51\x40\xf1\x19\x5e\x9d\x8a\xa6\x94\x6a\xc2\xa8\xda\x3d\xb2\x99\x23\x9f\x69\x14\xc8\xcc\x31\x64\x8e\x15\x1b\x84\xdd\x4c\xe8\xd1\xf9\x30\x84\xc9\x52\x14\xc2\xa4\x13\xb2\x37\xea\x1e\x06\xe6\xfb\x24\xb2\xef\x1b\xe6\xe4\x04\xfe\xd1\x6b\x88\xfb\x6a\x7c\x8f\x2b\xe1\x07\x36\x33\xe4\xba\xb4\xc8\x37\x92\x75\x38\x26\x22\xb2\xf6\x41\x42\x46\xbb\x3f\xd5\x81\xc9\x41\xf2\xb8\x37\xa6\x26\x1c\xdb\x23\x24\xd2\xf5\x1d\xdd\x8d\x7d\x37\xcf\x9f\x63\x42\x68\x25\x64\x4b\x61\x60\x2f\x73\x7a\x01\x5c\xf9\x78\x1e\x7b\xa5\xdd\x46\xc0\x55\xc8\xe2\x24\x52\x64\xe8\x13\xc2\x95\x44\xcf\x78\x00\x11\xda\x7e\x04\x31\x8d\xea\xad\xab\xa8\x1c\x84\xa8\x85\xba\x4f\x16\xdd\x96\xb7\x1a\xf2\x3e\x2a\x72\x50\x3a\xa3\xa8\x3c\x9c\x6f\x1f\xc7\x77\x44\xa4\xf8\x74\x66\x9a\x40\x6a\xd4\xa0\x4b\xb7\x24\xa6\xd5\xe7\xb1\x82\xd7\x61\xe6\xb5\xf1\x13\xa6\xe3\xe5\xb6\x9f\xe6\x11\x36\x44\x7d\xf3\xeb\xc7\xf5\x2a\x1e\x58\xda\x20\x9d\xe2\x90\x7b\xe1\xf9\x30\xd3\xff\x7c\x9a\x37\x4e\x32\x85\x4b\x28\x38\xee\x0a\xe7\x1d\xbe\x5c\x9e\x9c\xd3\x0a\x68\xd4\x83\x84\xfc\xf6\x98\xfc\xd7\xf8\xaa\xa6\x8d\xd9\xa7\xd1\xc6\x78\x4d\xfe\x07\x00\x00\xff\xff\x00\x71\x67\x30\x2c\x06\x00\x00")

func include_procfile_bash_bytes() ([]byte, error) {
	return bindata_read(
		_include_procfile_bash,
		"include/procfile.bash",
	)
}

func include_procfile_bash() (*asset, error) {
	bytes, err := include_procfile_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/procfile.bash", size: 1580, mode: os.FileMode(420), modTime: time.Unix(1422999670, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _include_slug_bash = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x53\xef\x6f\xd3\x30\x10\xfd\xdc\xfc\x15\x87\x55\x69\xed\x07\x37\x8c\xaf\x28\x48\x68\x20\x34\x09\x01\x62\xad\x84\x34\xa6\xca\x24\xd7\xc4\x92\x13\x5b\xb6\x03\x5d\x80\xff\x9d\x73\x7e\x2c\xc9\xda\x4d\xf4\x53\x73\xbe\xf7\xfc\xfc\xee\x5e\x64\x51\x64\xba\x52\xf7\xe0\x54\x9d\xef\x8d\xf0\x45\xc2\x62\x5f\x9a\x38\x7c\x6f\x7c\xde\xb0\x28\x0a\x7f\xb9\x2c\x8d\xb6\x7e\xb5\x86\xdf\xd1\x22\xc3\x54\x09\x8b\x90\xa1\x4b\x13\x76\xdd\x9e\x80\x80\xbc\x91\xc6\x60\xd6\x52\x81\x17\xf6\x87\x50\x0a\x0e\x56\x97\xb0\xfb\xfa\x11\xb4\x85\x9b\xed\xbb\xeb\x4f\xc0\x46\x82\xda\xaa\x84\x2d\x2f\xa9\x82\x95\xab\x2d\xf2\x20\xc0\x45\x0b\x79\x80\xdb\x5b\x60\xcb\x95\x72\xc0\xdf\xc2\x52\x18\xd3\x6a\x5b\x33\xb8\xbb\x7b\x0d\xbe\xc0\x2a\x5a\x2c\x2c\xfa\xda\x56\x70\x49\x68\x35\x20\x88\x71\xd6\x93\x52\x01\x38\xb1\x70\xea\xb6\xf7\xf0\x6a\xe8\xf9\x13\x14\x02\x3f\x36\x57\x54\x19\xf8\x83\x10\xe5\x30\xe0\x84\x7f\xb2\xe5\x20\xa3\xbf\xbd\x2b\x39\x56\x68\x85\xc7\x73\xbe\x7c\xe8\xcf\x9e\x75\x86\x64\x02\x69\xb4\x58\x91\x83\xc6\x9c\x38\xa1\x74\x2a\x14\xa4\xba\x34\x16\x9d\xdb\x6b\xe3\xa5\xae\x12\xc6\x69\x2e\xc1\xa5\x5f\x85\x4c\x0b\x30\x32\x6f\xe0\x0d\xc4\x19\xfe\x8c\xab\x5a\xa9\xf1\xf5\x27\x38\x5e\x3b\xe4\x43\x99\x1b\xab\x73\x2b\xca\x24\x10\x74\x0f\xeb\x2f\x0c\x42\x65\x5e\x69\x8b\x3d\x74\x98\x09\x3f\x4c\xbc\x88\x37\x63\xdf\xcc\xf5\x13\x38\xdd\xfc\x0d\xce\xe3\xba\x6b\x83\xd3\xcb\x47\x72\x61\x79\xc2\x03\xdf\x89\x9d\x73\x3c\xa6\xaa\xce\x30\xb9\xd8\xe4\xd2\x5f\x74\xc5\xd9\x90\xba\x52\x1a\xc4\x3e\x2c\x76\x28\x02\xfd\x16\x9b\xe9\x2b\xf7\x4e\x36\x48\x4b\xb8\xca\x6a\xe0\x37\x05\x8c\xfd\x34\xff\xb4\xf6\xf4\xe2\xcb\x35\x89\xf4\xd2\x2b\x04\x76\x45\x12\xa5\x1a\x46\x19\xb0\x20\x5d\x0f\x0a\x5f\xec\x61\x35\xf0\xf8\x54\x60\xde\xb7\x27\x30\xec\xce\xa3\xb5\xf0\xba\x8d\xcb\xea\xcb\x6e\xbb\xee\x43\xf3\x79\xb7\xfd\xef\xd0\xbc\xe8\x46\x34\x79\xf5\xd9\xc4\x04\xcf\x9f\xcf\xcc\xcb\x36\x36\x7a\xdc\xaa\x49\x86\x68\x98\x24\x0f\xf8\x76\x7e\x53\x47\x35\xcb\xd0\xf4\xbc\x4b\xce\xbf\x00\x00\x00\xff\xff\x5e\x30\xc4\xb4\x73\x04\x00\x00")

func include_slug_bash_bytes() ([]byte, error) {
	return bindata_read(
		_include_slug_bash,
		"include/slug.bash",
	)
}

func include_slug_bash() (*asset, error) {
	bytes, err := include_slug_bash_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "include/slug.bash", size: 1139, mode: os.FileMode(420), modTime: time.Unix(1422884875, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"include/buildpack.bash": include_buildpack_bash,
	"include/buildpacks.txt": include_buildpacks_txt,
	"include/cedarish.txt": include_cedarish_txt,
	"include/cmd.bash": include_cmd_bash,
	"include/fn.bash": include_fn_bash,
	"include/herokuish.bash": include_herokuish_bash,
	"include/procfile.bash": include_procfile_bash,
	"include/slug.bash": include_slug_bash,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"include": &_bintree_t{nil, map[string]*_bintree_t{
		"buildpack.bash": &_bintree_t{include_buildpack_bash, map[string]*_bintree_t{
		}},
		"buildpacks.txt": &_bintree_t{include_buildpacks_txt, map[string]*_bintree_t{
		}},
		"cedarish.txt": &_bintree_t{include_cedarish_txt, map[string]*_bintree_t{
		}},
		"cmd.bash": &_bintree_t{include_cmd_bash, map[string]*_bintree_t{
		}},
		"fn.bash": &_bintree_t{include_fn_bash, map[string]*_bintree_t{
		}},
		"herokuish.bash": &_bintree_t{include_herokuish_bash, map[string]*_bintree_t{
		}},
		"procfile.bash": &_bintree_t{include_procfile_bash, map[string]*_bintree_t{
		}},
		"slug.bash": &_bintree_t{include_slug_bash, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

