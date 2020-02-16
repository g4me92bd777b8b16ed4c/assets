// Assets package contains gzipped assets
// or drop the NewReader in most cases..
package assets

import (
	"io/ioutil"
	"strings"
)

//go:generate go get -v github.com/shurcooL/vfsgen/cmd/vfsgendev
//go:generate vfsgendev -source="github.com/g4me92bd777b8b16ed4c/assets".Assets
//go:generate bash rename.bash

// Assets contains project assets.
//var Assets http.FileSystem = http.Dir("assets")

func MustAsset(name string) []byte {
	name = strings.ToLower(name)
	f, err := Assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return b
}
