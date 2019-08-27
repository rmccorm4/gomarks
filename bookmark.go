package main

import (
	"encoding/json"
	"io/ioutil"
)

type Bookmark struct {
	URL  string
	Tags []string
}

type BookmarkList []Bookmark

func (bl *BookmarkList) load(filename string) {
	bs, err := ioutil.ReadFile(filename)
	err = json.Unmarshal(bs, bl)
	if err != nil {
		panic(err)
	}
}

func (bl *BookmarkList) save(filename string) {
	bs, err := json.MarshalIndent(bl, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, bs, 0644)
	if err != nil {
		panic(err)
	}
}
