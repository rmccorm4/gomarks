# Go-marks

<img src=https://github.com/egonelbre/gophers/blob/master/vector/fairy-tale/witch-learning.svg height=150 width=150> 

📖 Store bookmarks with tags for fast lookup from the command line!

----

## Usage

via `go build`
```bash
go build .
./gomarks <url> <tag1> <tag2> ...
```

via `go run`
```bash
go run . <url> <tag1> <tag2> ...
```

## Testing

```bash
go test -v
```

## TODO
- [x] Add URLs/Tags without duplicates
- [x] List URLs by Tag, or ListAll with no args
- [x] Store in a ~/.config/gomarks location
- [ ] Switch to use CLI library
- [ ] Add subcommand
- [ ] List subcommand
- [ ] Ability to delete tags
- [ ] Ability to delete URLs
- [ ] Make config file location configurable
- [ ] Create package to import
- [ ] Auto generate keywords for links with some naive text processing
- [ ] Grab tags from github repos via GitHub API
- [ ] Unit Tests
- [ ] CI/CD (TravisCI, Github Actions?)
- [ ] Code Coverage
- [ ] Badges
- [ ] BadgerDB or other database?
