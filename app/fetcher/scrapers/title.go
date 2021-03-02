package scrapers

import "strings"

// getTitle scrapes the movie title from IMDb page
func (c *customSelector) getTitle() string {
	title := strings.TrimSpace(c.Find("td.titleColumn a").Text())
	return title
}
