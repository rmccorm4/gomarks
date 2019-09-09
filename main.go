package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"
)

func usage() string {
	return `gomark is a tool to bookmark websites with tags for faster future lookup.

Usage:
	./gomark <url> [<tag1> [<tag2> <tag3>... <tagN>]]

Arguments:
	<url> string - Website to bookmark
	<tag> string - Tag(s) to apply to the URL for future look-up
`
}

func containsHelpFlag(args []string) bool {
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			return true
		}
	}
	return false
}

func getConfigDir() string {
	usr, _ := user.Current()
	homeDir := usr.HomeDir
	configDir := path.Join(homeDir, ".gomarks")

	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		log.Fatalf("[Error] %s", err)
	}
	return configDir
}

func unique(tags []string) []string {
	keys := make(map[string]bool, len(tags))
	list := []string{}
	for _, val := range tags {
		if _, exists := keys[val]; !exists {
			keys[val] = true
			list = append(list, val)
		}
	}
	return list
}

func main() {
	if len(os.Args) < 3 || containsHelpFlag(os.Args) {
		fmt.Println(usage())
		return
	}
	url := URL(os.Args[1])
	tags := []Tag(os.Args[2:])

	//helpText := "Comma separated list of tags to assign to given link. (mark URL -tags t1,t2,t3)"

	fmt.Println("URL:", url)
	fmt.Println("Tags:", tags)
	configDir := getConfigDir()
	fmt.Println(configDir)

	//bookmarkURLsFile := path.Join(configDir, "bookmarkURLs.json")
	bookmarkTagsFile := path.Join(configDir, "bookmarkTags.json")

	//bookmarkURLs := BookmarkURLs{}
	// NoOp if file doesn't exist
	//bookmarkURLs.load(bookmarkURLsFile)

	bookmarkTags := BookmarkTags{}
	// NoOp if file doesn't exist
	bookmarkTags.load(bookmarkTagsFile)

	//bookmarkURLs.add(url, tags)
	bookmarkTags.add(tags, url)

	//bookmarkURLs.save(bookmarkURLsFile)
	bookmarkTags.save(bookmarkTagsFile)

	bookmarkTags.listAll()
	queryTags := []Tag{"emulator"}
	bookmarkTags.list(queryTags)
}
