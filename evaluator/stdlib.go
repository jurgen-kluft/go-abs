// Code generated for package evaluator by go-bindata DO NOT EDIT. (@generated)
// sources:
// stdlib/cli/index.abs
// stdlib/runtime/index.abs
// stdlib/util/index.abs
package evaluator

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

// Mode return file modify time
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

var _stdlibCliIndexAbs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x54\xd1\x6e\xda\x30\x14\x7d\xf7\x57\x9c\x25\x95\x48\x56\xc6\xc3\xfa\xd4\x4e\x68\xeb\x2a\x4d\xaa\x54\xed\x61\x7b\x64\x68\x32\xf1\x4d\xb0\x70\x9c\xc8\x76\x68\xbb\x8a\x7f\x9f\x9c\x38\x90\x40\x21\x4f\x90\xdc\x73\xae\xcf\xb9\xe7\x3a\xc6\xc3\xd3\x23\x78\x5d\xb3\x4c\x49\xcc\xf1\xb6\x63\x2c\xc6\x43\x55\x96\x5c\x0b\x0b\x43\x85\xb4\x8e\x0c\x09\x3c\x4b\xb7\x96\x1a\x6e\x2d\xed\x10\x33\xcb\xfa\xda\x1e\xfc\xa3\xd1\x99\x93\x95\x46\x63\x49\xc0\x55\x7b\x12\x70\x84\xe2\x0e\x58\x0a\xcc\x91\x27\x9a\x97\x34\x85\x20\x9b\x19\x59\x7b\xe0\x14\xb9\xe2\x45\x47\x98\xe2\x8d\x01\x80\x21\xd7\x18\x8d\x3c\xc9\x75\xff\xca\x3f\xc3\x03\x2c\x3c\xd1\xb2\x45\x7d\xb9\x50\xb0\xef\x3b\xe4\xf1\x4f\x8c\x07\x43\xdc\x51\xe8\xee\xf5\x42\x50\xce\x1b\xe5\xb0\xe5\xaa\x21\x3b\x2a\xcf\x2b\x83\xcd\x14\x7f\x21\x75\x40\x8c\xe9\xfc\xb3\xf5\x8d\x14\x2f\x92\x4d\xca\x4e\x3e\xca\x1c\xdb\x77\x30\x2d\xb7\xe7\x5b\x6c\xbc\x98\xed\x49\xc1\x8e\x9d\xff\x77\x2c\x88\x2b\x05\xb7\x26\x54\x46\x16\x52\x73\x85\xac\x14\xa3\x1a\x43\xd6\xeb\x9b\x23\xd7\xb3\x8c\x2b\x95\x2c\xb8\x29\x6c\x92\x2e\x6e\xee\x96\x61\x0e\xcb\xa3\xb3\xc7\x78\xcc\x3d\xa9\xa1\x89\x05\xef\x19\x9e\x09\xb5\x91\xda\x41\x3a\x54\x8d\x63\x47\x4a\x43\xd5\xa9\x5c\xca\xd6\x55\xd2\x7d\x4d\xcf\x08\xdb\xb1\x4b\xe3\x1c\x04\x07\xf3\x61\x8c\x58\x87\x6d\x33\xf9\xab\xd1\xad\x0f\xc3\xe0\x9a\x46\x8f\x72\x10\xe3\xfe\xfb\x6f\x58\x22\x8b\x88\xaf\x2c\x3a\xa2\x99\xff\xf9\xf2\xfa\x2f\x0a\x35\xb6\x6a\x89\xc2\x21\x20\x6d\xfb\xf7\xc6\x08\x70\x53\x34\x25\xe9\x4e\x7a\x17\x32\x6e\x8a\xe4\x73\xb0\x2f\xc6\xcf\xca\xa1\xe6\xd6\x4a\x5d\x1c\x56\xe1\x2b\x9e\xc8\x4d\x6c\x30\xcf\x73\xad\x49\xd5\x2c\xd8\xf6\xc1\xf3\x1c\x4c\x0b\x5b\x30\x72\x61\xe2\xeb\x27\x6d\xac\x93\x94\x0d\xfc\x6a\xe1\xc3\xca\xac\x14\xcb\x01\x19\xbd\x48\x97\xdc\xde\x4e\x11\xf5\x62\x26\x57\x6f\x59\x29\x76\x13\xe8\xca\x21\xaf\x1a\x2d\xa2\x11\xe3\x7b\xed\x3d\x69\xe8\xbd\x63\x9d\xaf\x54\xab\x91\xb1\xed\x88\xa3\xab\xc0\xe5\xf7\xc6\xab\x92\x1a\xd6\x09\xa9\x07\x27\x92\xf9\x29\xf7\x51\x64\xce\xf5\xde\x67\xa5\x5f\x82\x30\xf8\x7b\x21\xc0\xf7\x6b\xec\xad\xda\x8f\xce\xad\xb9\x43\xc6\x35\x56\xc4\x62\x54\x5b\x32\x46\x0a\x41\x1a\xab\xd7\x6e\xc4\x5c\x29\x32\xec\x5b\xb8\xab\x92\xc8\xa3\xa3\x29\xa2\x7e\x54\xd2\x76\x84\x25\x59\xcb\x0b\x8a\xa6\xfe\xbe\x62\x79\xfb\xf2\x48\xfc\xfd\x96\x4b\xc5\x57\x6a\x1f\x1c\x7b\xf7\x47\x47\x21\x19\x03\x47\x86\xf2\x66\x1b\x7a\xb5\x49\x3a\xb3\x95\x71\xa3\xcb\xca\xdf\x8c\x11\xf0\x11\xdd\xb8\x22\x76\xd1\xc0\xd1\x82\x8c\xcd\xb4\xb8\xf6\x4c\x9f\x10\xe1\xfa\x32\xf0\xbd\x5d\x6c\x95\xd9\xf4\xe0\xf6\x21\x1e\xec\x7f\x00\x00\x00\xff\xff\x5b\x6d\x77\x28\x57\x06\x00\x00")

func stdlibCliIndexAbsBytes() ([]byte, error) {
	return bindataRead(
		_stdlibCliIndexAbs,
		"stdlib/cli/index.abs",
	)
}

func stdlibCliIndexAbs() (*asset, error) {
	bytes, err := stdlibCliIndexAbsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "stdlib/cli/index.abs", size: 1623, mode: os.FileMode(436), modTime: time.Unix(1650014557, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _stdlibRuntimeIndexAbs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4a\x2d\x29\x2d\xca\x53\xa8\xe6\x52\x50\x50\x50\x50\xca\x4b\xcc\x4d\x55\xb2\x52\x50\x4a\x4c\x2a\x56\xd2\x81\x08\x95\xa5\x16\x15\x67\xe6\xe7\x29\x59\x29\x38\x3a\x05\xc7\x87\xb9\x06\x05\x7b\xfa\xfb\x41\xe5\x32\xf3\x4a\x52\x8b\x12\x93\x4b\x32\xcb\x52\xa1\xf2\x9e\x7e\x21\xae\x41\x8e\xce\x21\x9e\x61\xae\x3a\x5c\xb5\x5c\x80\x00\x00\x00\xff\xff\x68\x41\xac\x26\x5e\x00\x00\x00")

func stdlibRuntimeIndexAbsBytes() ([]byte, error) {
	return bindataRead(
		_stdlibRuntimeIndexAbs,
		"stdlib/runtime/index.abs",
	)
}

func stdlibRuntimeIndexAbs() (*asset, error) {
	bytes, err := stdlibRuntimeIndexAbsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "stdlib/runtime/index.abs", size: 94, mode: os.FileMode(436), modTime: time.Unix(1744912536, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _stdlibUtilIndexAbs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\x4d\xcb\xdb\x30\x10\x84\xef\xfa\x15\x83\x73\xb1\xdb\xb7\xc2\x6f\xa1\x97\x50\xf7\xd4\x63\x8f\x81\x1e\x4a\x09\xaa\xb3\x4a\x04\xb2\x14\xa4\x35\x29\x09\xfe\xef\x45\xfe\x08\xfe\x48\x7d\x30\x68\xf6\xd9\xd5\x8c\xa4\x1d\xbe\x53\xed\x83\x62\x1f\xc0\x1e\x0d\x35\xde\xdc\x49\xec\xc0\x17\x42\xa0\xd8\x5a\x86\xd7\x50\xd0\xad\xab\xd9\x78\x27\x85\x9e\xa8\x9c\xd9\x16\x78\x08\x00\xd8\xe1\x27\xe1\xa6\x1c\xa7\x29\x91\x7d\x20\xb0\x69\x28\xad\xd2\xa4\xc6\x58\x6b\x22\xd5\xde\x9d\xde\x46\x5e\x59\xeb\x6f\xc6\x9d\xa1\x7d\xc0\xe1\xf0\x23\x26\xf6\x0f\xe1\x2b\xde\x23\x72\x3a\xa3\x94\x9f\xbf\x94\x85\xec\x71\x66\x8b\xaa\xff\x7f\xc0\x7b\x59\x96\xbd\x98\x6c\x1c\x1b\x75\x45\x85\x47\x27\x7a\x29\x10\xb7\xc1\x41\xe7\xda\x4d\xce\x16\xf2\x5c\x4c\x5f\xa4\x60\x94\x35\x77\x3a\x1d\x55\x38\x47\x54\x90\x52\xca\xc8\x21\x2f\x16\x5c\xad\xea\x0b\x9d\x50\x3d\xf7\xfc\xb5\xea\xfc\xbd\xc0\x9d\xbf\xa1\x42\xeb\xcc\xdf\x63\x13\xf3\x42\x2c\x8a\x46\x4f\xe3\x96\x5e\x16\x35\xc9\x11\x1f\xfb\xc0\x9f\xfa\x69\xdf\x50\xbe\xc0\x67\xd9\xc6\xb6\xe1\xca\x36\x60\x27\x36\xd2\x94\x44\x5e\xfd\x35\x5f\xa5\x59\x86\x5f\x35\x07\x4a\xe7\xa4\x9d\xac\x95\xb5\xb9\x94\x72\x95\xef\xbf\x47\x94\xee\x69\x63\x23\xe3\x98\xed\x53\xc4\xb7\x6d\x69\x08\x93\xed\xd3\x96\x2b\x47\xf3\x95\x78\x71\x1c\xf3\x8e\x81\xee\x44\x27\xc4\x58\x1d\x6c\x64\xe3\x3b\xce\xf6\xcf\x77\xdf\x89\x7f\x01\x00\x00\xff\xff\x0a\xd8\xb2\x18\x11\x03\x00\x00")

func stdlibUtilIndexAbsBytes() ([]byte, error) {
	return bindataRead(
		_stdlibUtilIndexAbs,
		"stdlib/util/index.abs",
	)
}

func stdlibUtilIndexAbs() (*asset, error) {
	bytes, err := stdlibUtilIndexAbsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "stdlib/util/index.abs", size: 785, mode: os.FileMode(436), modTime: time.Unix(1650014557, 0)}
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
	"stdlib/cli/index.abs":     stdlibCliIndexAbs,
	"stdlib/runtime/index.abs": stdlibRuntimeIndexAbs,
	"stdlib/util/index.abs":    stdlibUtilIndexAbs,
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
	"stdlib": &bintree{nil, map[string]*bintree{
		"cli": &bintree{nil, map[string]*bintree{
			"index.abs": &bintree{stdlibCliIndexAbs, map[string]*bintree{}},
		}},
		"runtime": &bintree{nil, map[string]*bintree{
			"index.abs": &bintree{stdlibRuntimeIndexAbs, map[string]*bintree{}},
		}},
		"util": &bintree{nil, map[string]*bintree{
			"index.abs": &bintree{stdlibUtilIndexAbs, map[string]*bintree{}},
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
