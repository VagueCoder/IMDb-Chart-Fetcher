package fetcher

import (
	"io"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/VagueCoder/IMDb-Chart-Fetcher/app/fetcher/scrapers"
)

// FetchItems scrapes movies details and returns JSON marshalled slice of bytes
func FetchItems(url string, total int, writer io.Writer) {
	logger := log.New(os.Stderr, "[IMDb-Chart-Fetcher] ", log.LstdFlags|log.Lshortfile)

	document, err := goquery.NewDocument(url)
	if err != nil {
		logger.Fatalf("goquery Document Creation Error: %v", err)
	}

	scraperObject := scrapers.NewScraper(document, nil, logger)
	scraperObject.GetMovieDetails(total)

	scraperObject.Encode(writer)
}
