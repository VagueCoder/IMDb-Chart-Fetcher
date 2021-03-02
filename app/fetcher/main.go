package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/VagueCoder/IMDb-Chart-Fetcher/app/fetcher/scrapers"
)

func main() {
	logger := log.New(os.Stderr, "[IMDb-Chart-Fetcher] ", log.LstdFlags|log.Lshortfile)
	url := "https://www.imdb.com/india/top-rated-indian-movies/"
	// response, err := http.Get(url)
	// if err != nil {
	// 	logger.Fatalf("http GET error: %v", err)
	// }
	// if response.StatusCode != http.StatusOK {
	// 	logger.Fatalf("http GET Status Expected: %v. Received: %v.", http.StatusOK, response.Status)
	// }
	// defer response.Body.Close()

	// document, err := goquery.NewDocumentFromReader(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	document, err := goquery.NewDocument(url)
	if err != nil {
		logger.Fatalf("goquery Document Creation Error: %v", err)
	}

	// selector := document.Find("div.lister")

	scraperObject := scrapers.NewScraper(document, nil, logger)
	scraperObject.GetMovieDetails(10)

	fmt.Printf("%+v", scraperObject)
}
