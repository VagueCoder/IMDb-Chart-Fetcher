run:
	go run ./app/main.go 'https://www.imdb.com/india/top-rated-indian-movies/' 2

build:
	go build -o IMDb-Chart-Fetcher ./app

execute:
	./IMDb-Chart-Fetcher 'https://www.imdb.com/india/top-rated-indian-movies/' 2

test:
	go test ./app/fetcher/scrapers/ -v
	go test ./app/fetcher/ -v
	go test ./app/ -v

all:
	go run ./app/main.go 'https://www.imdb.com/india/top-rated-indian-movies/' 500
	go run ./app/main.go 'https://www.imdb.com/india/top-rated-tamil-movies/' 500
	go run ./app/main.go 'https://www.imdb.com/india/top-rated-telugu-movies/' 500