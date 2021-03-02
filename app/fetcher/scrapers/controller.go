package scrapers

import (
	"bytes"
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Scraper struct {
	Logger   *log.Logger
	Selector *goquery.Document
	Movie    Movies
}

// MovieDetails object holds the details of movie
type MovieDetails struct {
	Title    string `json:"title"`
	Year     string `json:"movie_release_year"` // Add Validator
	Rating   string `json:"imdb_rating"`
	Summary  string `json:"summary"`
	Duration string `json:"duration"`
	Genre    string `json:"genre"`
}

type Movies []MovieDetails

func NewScraper(resp *goquery.Document, movie Movies, logger *log.Logger) *Scraper {
	return &Scraper{
		Selector: resp,
		Movie:    movie,
		Logger:   logger,
	}
}

func (s *Scraper) GetMovieDetails(total int) {
	s.Selector.Find("tbody.lister-list").Find("tr").EachWithBreak(func(i int, tr *goquery.Selection) bool {

		if i+1 == total {
			return false
		}

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

		subdoc, err := goquery.NewDocument(url)
		if err != nil {
			s.Logger.Fatalf("goquery Document Creation Error: %v", err)
		}

		// Summary
		summary := strings.TrimSpace(subdoc.Find("div.summary_text").Text())
		movie.Summary = summary
		// fmt.Println(summary)

		// Duration
		duration := strings.TrimSpace(subdoc.Find("div.subtext time").Text())
		movie.Duration = duration
		// fmt.Println(duration)

		// Genre
		children := subdoc.Find("div.subtext")
		selection1 := children.Find("span.ghost").Eq(1).NextAllFiltered("a")
		selection2 := children.Find("span.ghost").Eq(2).PrevAllFiltered("a")
		text = selection1.Intersection(selection2).Text()
		pattern := regexp.MustCompile("[A-Z][a-z]+")
		byteSlice := pattern.FindAll([]byte(text), -1)
		genre := string(bytes.Join(byteSlice, []byte(", ")))
		movie.Genre = genre
		// fmt.Println(genre)

		s.Movie = append(s.Movie, *movie)

		return true
	})
}

func (s Scraper) encode(collection Movies) []byte {
	encoded, err := json.Marshal(collection)
	if err != nil {
		s.Logger.Fatalf("Error Encoding JSON: %v", err)
	}
	return encoded
}
