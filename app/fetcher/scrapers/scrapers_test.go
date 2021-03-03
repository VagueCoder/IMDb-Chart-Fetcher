package scrapers

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

var (
	scraperObj *scraper
	selector   *goquery.Selection
	movie      *movieDetails
	logger     *log.Logger
)

func createScraper() (*scraper, error) {
	url := "https://www.imdb.com/india/top-rated-indian-movies/"
	document, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	return NewScraper(document, nil, &log.Logger{}), nil
}

func TestScraper(t *testing.T) {
	var err error
	scraperObj, err = createScraper()
	if err != nil {
		assert.FailNowf(t, "goquery Document Creation failed", "Error: %v", err)
	}
	assert.NotNil(t, scraperObj, "Scraper Object is Nil")
}

func TestScraperFuncs1(t *testing.T) {
	selector = scraperObj.Selector.Find("tbody.lister-list").Find("tr").First()
	logger = log.New(os.Stderr, "", log.LstdFlags)
	plainSelector := &customSelector{selector, logger}

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

func TestScraperFuncs2(t *testing.T) {
	selectorWithDoc := &customSelector{scraperObj.pageSelector(selector), logger}
	assert.NotNil(t, selectorWithDoc, "selectorWithDoc Object is Nil")
	assert.NotNil(t, movie, "movie Object is Nil")

	movie.Summary = selectorWithDoc.getSummary()
	assert.NotNil(t, movie.Summary, "summary is Nil")

	movie.Duration = selectorWithDoc.getDuration()
	assert.NotNil(t, movie.Duration, "duration is Nil")

	movie.Genre = selectorWithDoc.getGenre()
	assert.NotNil(t, movie.Genre, "genre is Nil")
}

func TestEncoding(t *testing.T) {
	assert.NotNil(t, movie, "movie Object is Nil")
	assert.NotNil(t, scraperObj, "scraperObj Object is Nil")
	assert.Empty(t, scraperObj.MoviesCollection, "The scraperObj.MoviesCollection slice should be empty as not inputted yet")

	writer := &bytes.Buffer{}
	scraperObj.MoviesCollection = movies{*movie}
	scraperObj.Encode(writer)
	encodedString := strings.TrimSpace(writer.String())

	assert.NotNil(t, writer, "writer Object is Nil")
	assert.NotEqual(t, "null", encodedString, "Encoded Value Shouldn't be null")
	assert.Contains(t, encodedString, `"title"`, `encoded json dats should contain "title" key with quotes`)
	assert.Contains(t, encodedString, `"movie_release_year"`, `encoded json dats should contain "movie_release_year" key with quotes`)
	assert.Contains(t, encodedString, `"imdb_rating"`, `encoded json dats should contain "imdb_rating" key with quotes`)
	assert.Contains(t, encodedString, `"summary"`, `encoded json dats should contain "summary" key with quotes`)
	assert.Contains(t, encodedString, `"duration"`, `encoded json dats should contain "duration" key with quotes`)
	assert.Contains(t, encodedString, `"genre"`, `encoded json dats should contain "genre" key with quotes`)
}
