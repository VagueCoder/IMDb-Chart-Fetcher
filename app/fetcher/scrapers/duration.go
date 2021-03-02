package scrapers

import "strings"

// getDuration scrapes the movie duration details from IMDb page
func (c *customSelector) getDuration() string {
	duration := strings.TrimSpace(c.Find("div.subtext time").Text())
	return duration
}
