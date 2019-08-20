package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func createConfig() string {
	usr, _ := user.Current()
	homeDir := usr.HomeDir
	configDir := path.Join(homeDir, ".gomarks")
	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		log.Fatalf("[Error] %s", err)
	}
	return configDir
}

func main() {
	if len(os.Args) < 3 || containsHelpFlag(os.Args) {
		fmt.Println(usage())
		return
	}
	url := os.Args[1]
	tags := os.Args[2:]

	//helpText := "Comma separated list of tags to assign to given link. (mark URL -tags t1,t2,t3)"

	fmt.Println(url)
	fmt.Println(tags)
	configDir := createConfig()
	fmt.Println(configDir)

	b := Bookmark{
		URL:  url,
		Tags: tags,
	}
	fmt.Println(b)

	bookmarkFile := path.Join(configDir, "bookmarks.json")
	// Open our jsonFile
	jsonFile, err := os.Open(bookmarkFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)
	var bs Bookmarks
	json.Unmarshal(bytes, &bs)
	bs = append(bs, b)
	jsonData, _ := json.MarshalIndent(bs, "", " ")
	_ = ioutil.WriteFile(bookmarkFile, jsonData, 0755)
	fmt.Println("Wrote JSON to", bookmarkFile)

}
