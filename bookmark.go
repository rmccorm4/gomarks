package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type URL = string
type Tag = string
type BookmarkURLs map[URL][]Tag
type BookmarkTags map[Tag][]URL

func (bu *BookmarkURLs) load(filename string) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Couldn't read file %s: %v\n", filename, err)
		return
	}
	err = json.Unmarshal(bs, bu)
	if err != nil {
		log.Fatalf("Couldn't unmarshal JSON file %s: %v\n", filename, err)
	}
}

func (bu *BookmarkURLs) save(filename string) {
	bs, err := json.MarshalIndent(bu, "", "  ")
	if err != nil {
		log.Fatalf("Couldn't json.Marshal %v: %v\n", bu, err)
	}
	err = ioutil.WriteFile(filename, bs, 0644)
	if err != nil {
		log.Fatalf("Couldn't write JSON file %s: %v\n", filename, err)
	}
}

func (bt *BookmarkTags) load(filename string) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Couldn't read file %s: %v\n", filename, err)
		return
	}
	err = json.Unmarshal(bs, bt)
	if err != nil {
		log.Fatalf("Couldn't unmarshal JSON file %s: %v\n", filename, err)
	}
}

func (bt *BookmarkTags) save(filename string) {
	bs, err := json.MarshalIndent(bt, "", "  ")
	if err != nil {
		log.Fatalf("Couldn't json.Marshal %v: %v\n", bt, err)
	}
	err = ioutil.WriteFile(filename, bs, 0644)
	if err != nil {
		log.Fatalf("Couldn't write JSON file %s: %v\n", filename, err)
	}
}

func (bt *BookmarkTags) listAll() {
	for tag, urls := range *bt {
		fmt.Println(tag)
		for _, url := range urls {
			fmt.Printf("\t%s\n", url)
		}
	}
}

func (bt *BookmarkTags) list(tags []string) {
	for _, tag := range tags {
		fmt.Println(tag)
		urls := (*bt)[tag]
		for _, url := range urls {
			fmt.Printf("\t%s\n", url)
		}
	}
}
