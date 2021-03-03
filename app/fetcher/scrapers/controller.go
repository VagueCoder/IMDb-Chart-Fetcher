package scrapers

import (
	"encoding/json"
	"io"
	"log"
	"net/url"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// Variables that are shared between all the goroutines
var (
	wg    sync.WaitGroup
	mutex sync.Mutex
)

// scraper is a common data struct to define all the methods. Needless to export structure.
type scraper struct {
	Logger           *log.Logger
	Selector         *goquery.Document
	MoviesCollection movies
}

// movieDetails object holds the details of movie
type movieDetails struct {
	rank     int     // Required to order the result items due to goroutines. Not exported.
	Title    string  `json:"title"`
	Year     string  `json:"movie_release_year"` // Add Validator
	Rating   float32 `json:"imdb_rating"`
	Summary  string  `json:"summary"`
	Duration string  `json:"duration"`
	Genre    string  `json:"genre"`
}

// movies is collection of movieDetails objects to avoid calling slice in parameters
type movies []movieDetails

// customSelector is a wrapper for calling scrapers methods alone
type customSelector struct {
	*goquery.Selection
	url    *url.URL
	logger *log.Logger
}

// NewScraper initiates new scraper obkect and returns reference. Return object type is needless to export.
func NewScraper(resp *goquery.Document, movie movies, logger *log.Logger) *scraper {
	return &scraper{
		Selector:         resp,
		MoviesCollection: movie,
		Logger:           logger,
	}
}

// pageSelector scrapes the page URLs from chart page
// Creates and returns a goquery generic Selection object of full page
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
	for i := 0; i < total; i++ {
		s.MoviesCollection = append(s.MoviesCollection, movieDetails{})
	}

	var counter int

	// Create required number of goroutines to run concurrently
	wg.Add(total)
	s.Selector.Find("tbody.lister-list").Find("tr").EachWithBreak(func(i int, tr *goquery.Selection) bool {
		counter = i
		go s.scrapeMovieDetails(tr, &wg)
		if counter == total-1 {
			return false
		}
		return true
	})

	// When items_count is greater than available
	if counter != total-1 {
		s.Logger.Printf("Note: There total available number of items in the chart are %d\n", counter+1)
		wg.Add(-total + counter + 1)
		mutex.Lock()
		s.MoviesCollection = s.MoviesCollection[:counter+1]
		mutex.Unlock()
	}
	wg.Wait()

}

// scrapeMovieDetails scrapes all scraper functions and updates the scraper's MovieDetails
// This is utilized as a goroutine to improve the performance
func (s *scraper) scrapeMovieDetails(tr *goquery.Selection, wg *sync.WaitGroup) {
	plainSelector := &customSelector{tr, nil, s.Logger}
	selectorWithDoc := &customSelector{s.pageSelector(tr), s.Selector.Url, s.Logger}
	movie := movieDetails{
		rank:     plainSelector.getRank(),
		Title:    plainSelector.getTitle(),
		Year:     plainSelector.getYear(),
		Rating:   plainSelector.getRating(),
		Summary:  selectorWithDoc.getSummary(),
		Duration: selectorWithDoc.getDuration(),
		Genre:    selectorWithDoc.getGenre(),
	}
	mutex.Lock()
	s.MoviesCollection[movie.rank-1] = movie
	mutex.Unlock()
	wg.Done()
}

// Encode encodes the MovieCollection object and streams through writer
func (s scraper) Encode(writer io.Writer) {
	err := json.NewEncoder(writer).Encode(s.MoviesCollection)
	if err != nil {
		s.Logger.Fatalf("Error Encoding JSON: %v", err)
	}
}
