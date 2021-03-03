package scrapers

import (
	"encoding/json"
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type scraper struct {
	Logger           *log.Logger
	Selector         *goquery.Document
	MoviesCollection movies
}

// movieDetails object holds the details of movie
type movieDetails struct {
	Title    string  `json:"title"`
	Year     string  `json:"movie_release_year"` // Add Validator
	Rating   float32 `json:"imdb_rating"`
	Summary  string  `json:"summary"`
	Duration string  `json:"duration"`
	Genre    string  `json:"genre"`
}

// movies is collection of movieDetails objects
type movies []movieDetails

// customSelector is a wrapper for calling scrapers methods alone
type customSelector struct {
	*goquery.Selection
	logger *log.Logger
}

// NewScraper initiates new scraper obkect and returns reference
func NewScraper(resp *goquery.Document, movie movies, logger *log.Logger) *scraper {
	return &scraper{
		Selector:         resp,
		MoviesCollection: movie,
		Logger:           logger,
	}
}

// pageSelector scrapes the page URLs from chart page, creates a goquery Selection object and returns
func (s *scraper) pageSelector(tr *goquery.Selection) *goquery.Selection {

	path, ok := tr.Find("td.titleColumn a").Attr("href")
	if !ok {
		s.Logger.Fatalf("Scrapping error: Couldn't find path in td.titleColumn")
	}
	url := s.Selector.Url.Scheme + "://" + s.Selector.Url.Hostname() + path

	subdoc, err := goquery.NewDocument(url)
	if err != nil {
		s.Logger.Fatalf("goquery Document Creation Error: %v", err)
	}
	return subdoc.Find("*")
}

// GetMovieDetails is the controller of all scrapers.
// Takes number of items required and updates the MovieCollection object of scraper accordingly.
func (s *scraper) GetMovieDetails(total int) {
	s.Selector.Find("tbody.lister-list").Find("tr").EachWithBreak(func(i int, tr *goquery.Selection) bool {

		plainSelector := &customSelector{tr, s.Logger}
		selectorWithDoc := &customSelector{s.pageSelector(tr), s.Logger}

		movie := &movieDetails{
			Title:    plainSelector.getTitle(),
			Year:     plainSelector.getYear(),
			Rating:   plainSelector.getRating(),
			Summary:  selectorWithDoc.getSummary(),
			Duration: selectorWithDoc.getDuration(),
			Genre:    selectorWithDoc.getGenre(),
		}

		s.MoviesCollection = append(s.MoviesCollection, *movie)
		if i == total-1 {
			return false
		}
		return true
	})
}

// Encode encodes the MovieCollection object and streams through writer
func (s scraper) Encode(writer io.Writer) {
	err := json.NewEncoder(writer).Encode(s.MoviesCollection)
	if err != nil {
		s.Logger.Fatalf("Error Encoding JSON: %v", err)
	}
}
