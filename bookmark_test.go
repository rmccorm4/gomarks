package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	expectedURL = BookmarkURLs{
		"https://golang.org": []Tag{"go", "google", "docs", "code"},
		"https://github.com": []Tag{"code", "github", "favorite"},
	}
	expectedTag = BookmarkTags{
		"go":       []URL{"https://golang.org"},
		"google":   []URL{"https://golang.org"},
		"docs":     []URL{"https://golang.org"},
		"code":     []URL{"https://golang.org", "https://github.com"},
		"github":   []URL{"https://github.com"},
		"favorite": []URL{"https://github.com"},
	}
)

func TestURLSaveJSONStruct(t *testing.T) {
	expectedURL.save("test/saveURL.json")

	var got BookmarkURLs
	got.load("test/saveURL.json")

	if !cmp.Equal(expectedURL, got) {
		t.Errorf("Expected: \n%v\nGot: \n%v\n", expectedURL, got)
	}
}

func TestURLLoadJSON(t *testing.T) {
	var got BookmarkURLs
	got.load("test/loadURL.json")

	if !cmp.Equal(expectedURL, got) {
		t.Errorf("Expected: \n%v\nGot: \n%v\n", expectedURL, got)
	}
}

func TestTagSaveJSONStruct(t *testing.T) {
	expectedTag.save("test/saveTag.json")

	var got BookmarkTags
	got.load("test/saveTag.json")

	if !cmp.Equal(expectedTag, got) {
		t.Errorf("Expected: \n%v\nGot: \n%v\n", expectedTag, got)
	}
}

func TestTagLoadJSON(t *testing.T) {
	var got BookmarkTags
	got.load("test/loadTag.json")

	if !cmp.Equal(expectedTag, got) {
		t.Errorf("Expected: \n%v\nGot: \n%v\n", expectedTag, got)
	}
}
