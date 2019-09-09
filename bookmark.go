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

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

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

func (bu *BookmarkURLs) add(url URL, tags []Tag) {
	// Update tags for URL or create entry if it doesn't exist
	if existingTags, ok := (*bu)[url]; ok {
		(*bu)[url] = unique(append(existingTags, tags...))
	} else {
		(*bu)[url] = unique(tags)
	}
}

func (bt *BookmarkTags) add(tags []Tag, url URL) {
	// Add URL for tags or create entry if it doesn't exist
	for _, tag := range tags {
		if existingURLs, exists := (*bt)[tag]; exists {
			if !contains(existingURLs, url) {
				(*bt)[tag] = append(existingURLs, url)
			}
		} else {
			(*bt)[tag] = []URL{url}
		}
	}
}
