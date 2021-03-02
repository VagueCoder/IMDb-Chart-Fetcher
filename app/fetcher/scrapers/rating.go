package scrapers

import "strings"

// getRating scrapes the rating of the movie from IMDb page
func (c *customSelector) getRating() string {
	rating := strings.TrimSpace(c.Find("td.imdbRating strong").Text())
	return rating
}
