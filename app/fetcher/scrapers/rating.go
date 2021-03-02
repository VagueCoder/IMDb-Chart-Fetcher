package scrapers

import (
	"strconv"
	"strings"
)

// getRating scrapes the rating of the movie from IMDb page
func (c *customSelector) getRating() float32 {
	text := strings.TrimSpace(c.Find("td.imdbRating strong").Text())
	rating, err := strconv.ParseFloat(text, 32)
	if err != nil {
		c.logger.Fatalf("String parsing failed with error: %v", err)
	}
	return float32(rating)
}
