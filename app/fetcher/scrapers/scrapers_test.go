package scrapers

import (
	"bytes"
	"log"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

// Common variables that are created in one test and may be used in other tests
var (
	URL        string = "https://www.imdb.com/india/top-rated-indian-movies/"
	scraperObj *scraper
	selector   *goquery.Selection
	movie      *movieDetails
	logger     *log.Logger
)

// Not test function. Creates and returns reference to scraper object
func createScraper() (*scraper, error) {
	document, err := goquery.NewDocument(URL)
	if err != nil {
		return nil, err
	}
	return NewScraper(document, nil, &log.Logger{}), nil
}

// Basic tests of scraper object
func TestScraper(t *testing.T) {
	var err error // Required definition as direct := will overwrite global scraperObj too
	scraperObj, err = createScraper()
	if err != nil {
		assert.FailNowf(t, "goquery Document Creation failed", "Error: %v", err)
	}
	assert.NotNil(t, scraperObj, "Scraper Object is Nil")
}

// Tests of first 3 items of the scraper object
func TestScraperFuncs1(t *testing.T) {
	selector = scraperObj.Selector.Find("tbody.lister-list").Find("tr").First()
	logger = log.New(os.Stderr, "", log.LstdFlags)
	plainSelector := &customSelector{selector, nil, logger}

	title := plainSelector.getTitle()
	assert.NotNil(t, title, "title is Nil")

	year := plainSelector.getYear()
	assert.NotNil(t, year, "year is Nil")

	rating := plainSelector.getRating()
	assert.NotNil(t, rating, "rating is Nil")

	movie = &movieDetails{
		Title:  title,
		Year:   year,
		Rating: rating,
	}
	assert.NotNil(t, movie, "movie Object is Nil")
}

// Tests of last 3 items of the scraper object
func TestScraperFuncs2(t *testing.T) {
	url, err := url.Parse(URL)
	if err != nil {
		assert.Failf(t, "URL parsing failed", "URL parsing error: %v", err)
	}
	selectorWithDoc := &customSelector{scraperObj.pageSelector(selector), url, logger}

	// The shared objects shouldn't be nil
	assert.NotNil(t, selectorWithDoc, "selectorWithDoc Object is Nil")
	assert.NotNil(t, movie, "movie Object is Nil")

	movie.Summary = selectorWithDoc.getSummary()
	assert.NotNil(t, movie.Summary, "summary is Nil")

	movie.Duration = selectorWithDoc.getDuration()
	assert.NotNil(t, movie.Duration, "duration is Nil")

	movie.Genre = selectorWithDoc.getGenre()
	assert.NotNil(t, movie.Genre, "genre is Nil")
}

// All types of tests on encoded string
func TestEncoding(t *testing.T) {
	// The shared objects shouldn't be nil
	assert.NotNil(t, movie, "movie Object is Nil")
	assert.NotNil(t, scraperObj, "scraperObj Object is Nil")

	// The MoviesCollection should be nil as not assigned yet
	assert.Empty(t, scraperObj.MoviesCollection, "The scraperObj.MoviesCollection slice should be empty as not inputted yet")

	// Creating a bytes.Buffer as writer so we can read the encoded JSON back to variable
	writer := &bytes.Buffer{}
	scraperObj.MoviesCollection = movies{*movie}
	scraperObj.Encode(writer)
	encodedString := strings.TrimSpace(writer.String())

	// Tests on encoded JSON (already parsed as string)
	assert.NotNil(t, writer, "writer Object is Nil")
	assert.NotEqual(t, "null", encodedString, "Encoded Value Shouldn't be null")
	assert.Contains(t, encodedString, `"title"`, `encoded json dats should contain "title" key with quotes`)
	assert.Contains(t, encodedString, `"movie_release_year"`, `encoded json dats should contain "movie_release_year" key with quotes`)
	assert.Contains(t, encodedString, `"imdb_rating"`, `encoded json dats should contain "imdb_rating" key with quotes`)
	assert.Contains(t, encodedString, `"summary"`, `encoded json dats should contain "summary" key with quotes`)
	assert.Contains(t, encodedString, `"duration"`, `encoded json dats should contain "duration" key with quotes`)
	assert.Contains(t, encodedString, `"genre"`, `encoded json dats should contain "genre" key with quotes`)
}
