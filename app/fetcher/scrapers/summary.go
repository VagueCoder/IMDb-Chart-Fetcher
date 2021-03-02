package scrapers

import "strings"

// getSummary scrapes the movie summary from IMDb page
func (c *customSelector) getSummary() string {
	summary := strings.TrimSpace(c.Find("div.summary_text").Text())
	return summary
}
