package main

import (
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLoadJSON(t *testing.T) {
	expected := BookmarkList{
		Bookmark{
			URL:  "https://golang.org",
			Tags: []string{"go", "google", "docs", "code"},
		},
		Bookmark{
			URL:  "https://github.com",
			Tags: []string{"code", "github", "favorite"},
		},
	}

	var got BookmarkList
	got.load("test/load.json")

	if !cmp.Equal(expected, got) {
		t.Errorf("Expected: \n%v\nGot: \n%v\n", expected, got)
	}
}

func TestSaveJSONStruct(t *testing.T) {
	expected := BookmarkList{
		Bookmark{
			URL:  "https://blog.golang.org/concurrency-is-not-parallelism",
			Tags: []string{"go", "concurrency", "parallelism", "blog"},
		},
		Bookmark{
			URL:  "https://github.com/rmccorm4/gomarks",
			Tags: []string{"code", "github", "favorite", "bookmarks", "go"},
		},
	}
	expected.save("test/save.json")

	got := BookmarkList{}
	got.load("test/save.json")

	if !cmp.Equal(expected, got) {
		t.Errorf("Expected: \n%v\nGot: \n%v\n", expected, got)
	}
}

func TestSaveJSONFile(t *testing.T) {
	expected := `[
  {
    "URL": "https://blog.golang.org/concurrency-is-not-parallelism",
    "Tags": [
      "go",
      "concurrency",
      "parallelism",
      "blog"
    ]
  },
  {
    "URL": "https://github.com/rmccorm4/gomarks",
    "Tags": [
      "code",
      "github",
      "favorite",
      "bookmarks",
      "go"
    ]
  }
]`

	bookmarks := BookmarkList{
		Bookmark{
			URL:  "https://blog.golang.org/concurrency-is-not-parallelism",
			Tags: []string{"go", "concurrency", "parallelism", "blog"},
		},
		Bookmark{
			URL:  "https://github.com/rmccorm4/gomarks",
			Tags: []string{"code", "github", "favorite", "bookmarks", "go"},
		},
	}
	bookmarks.save("test/save.json")
	bookmarksJSON, err := ioutil.ReadFile("test/save.json")
	if err != nil {
		panic(err)
	}

	got := string(bookmarksJSON)
	if !cmp.Equal(expected, got) {
		t.Errorf("Expected: \n%v\nGot: \n%v\n", expected, got)
	}
}
