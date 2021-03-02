package scrapers

import (
	"regexp"
)

// getYear scrapes the movie release year from IMDb page
func (c *customSelector) getYear() string {
	yearPattern := regexp.MustCompile("[0-9]{4}")
	text := c.Find("td.titleColumn span").Text()
	year := string(yearPattern.Find([]byte(text)))
	return year
}
