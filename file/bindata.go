// Code generated by go-bindata.
// sources:
// ../config/config.yaml
// ../config/geth.config
// DO NOT EDIT!

package file

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

var _ConfigConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func ConfigConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_ConfigConfigYaml,
		"../config/config.yaml",
	)
}

func ConfigConfigYaml() (*asset, error) {
	bytes, err := ConfigConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../config/config.yaml", size: 0, mode: os.FileMode(420), modTime: time.Unix(1692007045, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ConfigGethConfig = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xbb\x6e\xeb\x30\x10\x44\x7b\x7e\x05\xcb\x7b\x0b\x53\x72\x80\x24\x15\x8b\xbc\x60\xa4\x49\x11\xd9\x48\xe1\xb8\x58\x53\x63\x8b\x09\x4d\x0a\xcb\x95\x6d\xfd\x7d\x20\xcb\x79\x76\x7b\x06\x33\x8b\xc1\x2c\x17\xd1\xcb\x4a\xdd\x23\x3b\xf6\xad\xf8\x14\xed\x0c\xd2\xe8\x87\x23\x5c\x37\xa0\xbe\x0b\x1e\x51\xf4\xbf\x59\x02\x07\xaf\xe7\xc8\xa2\x9f\x20\x87\xc4\xef\xff\xd5\xcd\x46\xc0\x36\x8e\x68\x84\x78\x0b\x51\x2f\x14\x25\xff\x15\x97\x15\x78\xef\x1d\x56\x6a\x91\xc1\x76\x0b\x69\xd4\x8c\x53\xd7\x8e\xe7\xbc\x6f\x61\xb3\xdf\xb5\x01\xea\x19\x59\x88\xc5\x52\x38\x50\x9f\x3f\xb1\x82\xb3\x97\x6a\xee\x77\x48\x9d\x54\x92\xda\x41\xb8\x2a\x4b\x35\x54\xad\x4e\x81\xa2\xcb\x5c\x84\xe4\x28\x14\x6b\x1f\x8b\xe1\xb1\x7e\x55\x5a\x4f\x26\xdb\xb1\xfc\x08\x35\x09\xd5\x9e\x75\xb1\x27\x2e\x82\x5f\xff\x34\x52\x27\x0d\xb7\xce\xbc\x1d\x24\xc3\x31\xe4\xdb\xf5\x25\x0d\x97\x69\x70\x3c\x47\x76\x10\xf6\x2e\xff\x26\x43\x75\xcd\x7a\x7a\x71\x6d\x4a\x53\x9a\xa9\x5a\x3e\xc6\x2c\x14\xc2\xea\xb4\x0e\xea\xdb\xde\xd6\xd8\x50\x17\xe4\x3c\xd0\x47\x00\x00\x00\xff\xff\xc8\x39\x63\x14\x89\x01\x00\x00")

func ConfigGethConfigBytes() ([]byte, error) {
	return bindataRead(
		_ConfigGethConfig,
		"../config/geth.config",
	)
}

func ConfigGethConfig() (*asset, error) {
	bytes, err := ConfigGethConfigBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../config/geth.config", size: 393, mode: os.FileMode(420), modTime: time.Unix(1692268171, 0)}
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
	"../config/config.yaml": ConfigConfigYaml,
	"../config/geth.config": ConfigGethConfig,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"..": &bintree{nil, map[string]*bintree{
		"config": &bintree{nil, map[string]*bintree{
			"config.yaml": &bintree{ConfigConfigYaml, map[string]*bintree{}},
			"geth.config": &bintree{ConfigGethConfig, map[string]*bintree{}},
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
