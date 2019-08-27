package main

import (
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLoadJSON(t *testing.T) {
	example := BookmarkList{
		Bookmark{
			URL:  "https://golang.org",
			Tags: []string{"go", "google", "docs", "code"},
		},
		Bookmark{
			URL:  "https://github.com",
			Tags: []string{"code", "github", "favorite"},
		},
	}

	var bookmarks BookmarkList
	bookmarks.load("test/load.json")

	if !cmp.Equal(example, bookmarks) {
		t.Errorf("Expected: \n%v\nGot: \n%v\n", example, bookmarks)
	}
}

func TestSaveJSON(t *testing.T) {
	exampleJSON := `[
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

	if !cmp.Equal(exampleJSON, string(bookmarksJSON)) {
		t.Errorf("Expected: \n%v\nGot: \n%v\n", exampleJSON, string(bookmarksJSON))
	}
}
