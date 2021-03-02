package scrapers

import (
	"bytes"
	"regexp"
)

// getGenre scrapes the genre of the movie from IMDb page
func (c *customSelector) getGenre() string {
	children := c.Find("div.subtext")
	selection1 := children.Find("span.ghost").Eq(1).NextAllFiltered("a")
	selection2 := children.Find("span.ghost").Eq(2).PrevAllFiltered("a")
	text := selection1.Intersection(selection2).Text()
	pattern := regexp.MustCompile("[A-Z][a-z]+")
	byteSlice := pattern.FindAll([]byte(text), -1)
	genre := string(bytes.Join(byteSlice, []byte(", ")))
	return genre
}
