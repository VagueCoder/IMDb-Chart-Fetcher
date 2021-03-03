package scrapers

import (
	"regexp"
	"strconv"
	"strings"
)

// getTitle scrapes the movie title from IMDb page
func (c *customSelector) getRank() int {
	pattern := regexp.MustCompile(`^([0-9]+)`)
	text := strings.TrimSpace(c.Find("td.titleColumn").Text())
	rank, err := strconv.Atoi(pattern.FindString(text))
	if err != nil {
		c.logger.Fatalf("Rank not parsable to integer: %v", err)
	}
	return rank
}
