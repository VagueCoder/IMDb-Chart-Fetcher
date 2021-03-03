package scrapers

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// getSummary scrapes the movie summary from IMDb page
func (c *customSelector) getSummary() string {
	summary := strings.TrimSpace(c.Find("div.summary_text").Text())

	// Special case when only a part of summary is deplayed on the page. Redirect and scrape.
	if strings.HasSuffix(summary, "See full summary »") {
		path, ok := c.Find("div.summary_text a").Attr("href")
		if !ok {
			c.logger.Fatalf("Scrapping error: Couldn't find path in td.titleColumn")
		}
		url := c.url.Scheme + "://" + c.url.Hostname() + path

		redirect, err := goquery.NewDocument(url)
		if err != nil {
			c.logger.Fatalf("goquery Document Creation Error: %v", err)
		}

		summary = strings.TrimSpace(redirect.Find("li.ipl-zebra-list__item p").First().Text())
	}
	return summary
}
