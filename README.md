# Go-marks

<img src=https://github.com/egonelbre/gophers/blob/master/vector/fairy-tale/witch-learning.svg height=150 width=150> 

📖 Store bookmarks with tags for fast lookup from the command line!

----

## Testing

```bash
go test -v
```

## TODO
- [ ] Use BadgerDB, store both URL:Tags and Tag:URLs as key/value pairs for simple and fast lookups
- [ ] Search bookmark file for URL and update tags if exists, don't duplicate tags
- [ ] subcommand list tags
- [ ] subcommand search by tag
- [ ] Switch to use CLI library
- [ ] Create package to import
- [ ] Auto generate keywords for links with some naive text processing
- [ ] Grab tags from github repos via GitHub API
- [ ] Store in a ~/.config/gomarks location (configurable)
- [ ] Nice UI/CLI
- [ ] Unit Tests
- [ ] TDD
- [ ] CI/CD
- [ ] Code Coverage
- [ ] Badges

## Issues
* [ ] [Vimium's](https://github.com/philc/vimium) bookmark search capabilities are pretty darn good
