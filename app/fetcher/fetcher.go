package fetcher

import (
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/VagueCoder/IMDb-Chart-Fetcher/app/fetcher/scrapers"
)

// FetchItems scrapes movies details and returns JSON marshalled slice of bytes
func FetchItems(total int) []byte {
	logger := log.New(os.Stderr, "[IMDb-Chart-Fetcher] ", log.LstdFlags|log.Lshortfile)
	url := "https://www.imdb.com/india/top-rated-indian-movies/"

	document, err := goquery.NewDocument(url)
	if err != nil {
		logger.Fatalf("goquery Document Creation Error: %v", err)
	}

	scraperObject := scrapers.NewScraper(document, nil, logger)
	scraperObject.GetMovieDetails(total)

	return scraperObject.Encode()
}
