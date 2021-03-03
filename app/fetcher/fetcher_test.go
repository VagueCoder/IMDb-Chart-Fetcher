package fetcher

import (
	"bytes"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetcher(t *testing.T) {
	url := "https://www.imdb.com/india/top-rated-indian-movies/"
	writer := &bytes.Buffer{}
	FetchItems(url, 3, writer)

	encodedString := strings.TrimSpace(writer.String())

	assert.NotNil(t, writer, "writer Object is Nil")
	assert.NotEqual(t, "null", encodedString, "Encoded Value Shouldn't be null")
	assert.Contains(t, encodedString, `"title"`, `encoded json dats should contain "title" key with quotes`)
	assert.Contains(t, encodedString, `"movie_release_year"`, `encoded json dats should contain "movie_release_year" key with quotes`)
	assert.Contains(t, encodedString, `"imdb_rating"`, `encoded json dats should contain "imdb_rating" key with quotes`)
	assert.Contains(t, encodedString, `"summary"`, `encoded json dats should contain "summary" key with quotes`)
	assert.Contains(t, encodedString, `"duration"`, `encoded json dats should contain "duration" key with quotes`)
	assert.Contains(t, encodedString, `"genre"`, `encoded json dats should contain "genre" key with quotes`)

	pattern := regexp.MustCompile(`^\[{.*},{.*}\]$`)
	assert.Regexp(t, pattern, encodedString, "Not enough items found in the result.")
}
