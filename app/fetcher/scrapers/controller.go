package scrapers

import (
	"encoding/json"
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
	Title    string `json:"title"`
	Year     string `json:"movie_release_year"` // Add Validator
	Rating   string `json:"imdb_rating"`
	Summary  string `json:"summary"`
	Duration string `json:"duration"`
	Genre    string `json:"genre"`
}

// movies is collection of movieDetails objects
type movies []movieDetails

// customSelector is a wrapper for calling scrapers methods alone
type customSelector struct {
	*goquery.Selection
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
		if i+1 == total {
			return false
		}

<<<<<<< HEAD
		// This is master

		movie := &MovieDetails{}
		// Title
		title := tr.Find("td.titleColumn a").Text()
		movie.Title = title
		// fmt.Printf("%d, %s\n", i, strings.TrimSpace(title))

		// Year
		yearPattern := regexp.MustCompile("[0-9]{4}")
		text := tr.Find("td.titleColumn span").Text()
		year := string(yearPattern.Find([]byte(text)))
		movie.Year = year
		// fmt.Printf("%d, %s\n", i, year)

		// Rating
		rating := tr.Find("td.imdbRating strong").Text()
		movie.Rating = rating
		// fmt.Printf("%d, %s\n", i, strings.TrimSpace(rating))

		path, ok := tr.Find("td.titleColumn a").Attr("href")
		if !ok {
			s.Logger.Fatalf("Scrapping error: Couldn't find path in td.titleColumn")
		}
		url := s.Selector.Url.Scheme + "://" + s.Selector.Url.Hostname() + path
=======
		plainSelector := &customSelector{tr}
		selectorWithDoc := &customSelector{s.pageSelector(tr)}
>>>>>>> current

		movie := &movieDetails{
			Title:    plainSelector.getTitle(),
			Year:     plainSelector.getYear(),
			Rating:   plainSelector.getRating(),
			Summary:  selectorWithDoc.getSummary(),
			Duration: selectorWithDoc.getDuration(),
			Genre:    selectorWithDoc.getGenre(),
		}

		s.MoviesCollection = append(s.MoviesCollection, *movie)
		return true
	})
}

// Encode method marshals the MovieCollection object and returns JSON
func (s scraper) Encode() []byte {
	encoded, err := json.Marshal(s.MoviesCollection)
	if err != nil {
		s.Logger.Fatalf("Error Encoding JSON: %v", err)
	}
	return encoded
}
