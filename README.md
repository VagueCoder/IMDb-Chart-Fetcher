# IMDb-Chart-Fetcher
Simple application that takes IMDb chart page URL (chart_url) and number of items (items_count) required, from command line, and returns the scraped movie details in JSON format.

## Base System Configurations :wrench:
**Sno.** | **Name** | **Version/Config.**
-------: | :------: | :------------------
1 | Operating System | WSL2 Ubuntu-20.04 
2 | Language | Go 1.13.8 linux/amd64
3 | IDE | VS Code 1.53.2 x64
4 | Script | GNU Make 4.2.1

## Files and Functionality :open_file_folder:
**Sno.** | **Filename** | **Comment**
-------: | :----------: | :----------
1 | [main.go](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/app/main.go) | Entrypoint of the application. Takes command line inputs of chart_url and items_count and calls the fetcher.
2 | [fetcher.go](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/app/fetcher/fetcher.go) | Package fetcher takes inputs from main and calls scraper controller, and finally the encoder.
3 | [controller.go](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/app/fetcher/scrapers/controller.go) | Controls all the operations of scraping, calls required scrapers and also has encoder function to finally stream the data to output.
4 | [duration.go](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/app/fetcher/scrapers/duration.go), [genre.go](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/app/fetcher/scrapers/genre.go), [rank.go](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/app/fetcher/scrapers/rank.go), [rating.go](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/app/fetcher/scrapers/rating.go), [summary.go](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/app/fetcher/scrapers/summary.go), [title.go](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/app/fetcher/scrapers/title.go), [year.go](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/app/fetcher/scrapers/year.go) | 7 scraper functions to scrape respective details for each item.
5 | *_test.go | Each of the 3 levels has respective test files to run unit tests.
6 | [go.mod](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/go.mod) and [go.sum](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/go.sum) | Has the dependency requirements of the application.
7 | [Makefile](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/Makefile) | Holds all the required commands to build/test/execute the application.
8 | [IMDb-Chart-Fetcher](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/IMDb-Chart-Fetcher) | Executable that runs wihout any dependency.

_Note: The program is divided into levels and files to achieve easy testability._

## Levels :1234:
**Sno.** | **Level** | **Comment**
-------: | :-------: | :----------
1 | App | Entrypoint of the application where main function resides. Functionality can be tested using [main_test.go](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/app/main_test.go).
2 | Fetcher | Package fetcher in the medium through which main calls the scrapers. In other words, it is an intermediate layer for functionality testing using [fetcher_test.go](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/app/fetcher/fetcher_test.go)
3 | Scrapers | All 7 scrapers (6 scrapers for fields of the exported JSON and 1 rank scraper to publish the output in order) resides here. Testing is done in one file, [scrapers_test.go](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/app/fetcher/scrapers/scrapers_test.go)

## How to use :question:
All the commands are already in [Makefile](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/Makefile) that can be used with `make` command or can be copied to terminal.

### 1. Direct run
Syntax: `go run path/to/main.go chart_url items_count`

```
go run ./app/main.go 'https://www.imdb.com/india/top-rated-indian-movies/' 2
```
> This works. But it is always recommended to use executable.

### 2. Build
Syntax: `go build -o executable_name path/to/main/directory`

```
go build -o IMDb-Chart-Fetcher ./app
```
[`IMDb-Chart-Fetcher`](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/IMDb-Chart-Fetcher) is now an executable and runs without any dependency.

### 3. Test
Syntax: `go test path/where/go/files/exists -v`
> **-v** is for enabling verbosity.

> chart_url & items_count are already specified in the test cases.

```
go test ./app/fetcher/scrapers/ -v
go test ./app/fetcher/ -v
go test ./app/ -v
```
>One for each level. Each test again contains multiple sub test cases as per requirement.
```
    go test ./app/fetcher/scrapers/ -v
    === RUN   TestScraper
    --- PASS: TestScraper (2.59s)
    === RUN   TestScraperFuncs1
    --- PASS: TestScraperFuncs1 (0.01s)
    === RUN   TestScraperFuncs2
    --- PASS: TestScraperFuncs2 (1.75s)
    === RUN   TestEncoding
    --- PASS: TestEncoding (0.00s)
    PASS
    ok      github.com/VagueCoder/IMDb-Chart-Fetcher/app/fetcher/scrapers   (cached)
    go test ./app/fetcher/ -v
    === RUN   TestFetcher
    --- PASS: TestFetcher (4.26s)
    PASS
    ok      github.com/VagueCoder/IMDb-Chart-Fetcher/app/fetcher    (cached)
    go test ./app/ -v
    === RUN   TestApplication
    --- PASS: TestApplication (7.79s)
    PASS
    ok      github.com/VagueCoder/IMDb-Chart-Fetcher/app    (cached)
```

### 4. Execute :fast_forward:
Syntax: `executable chart_url items_count`

```
./IMDb-Chart-Fetcher 'https://www.imdb.com/india/top-rated-indian-movies/' 500
./IMDb-Chart-Fetcher 'https://www.imdb.com/india/top-rated-tamil-movies/' 500
./IMDb-Chart-Fetcher 'https://www.imdb.com/india/top-rated-telugu-movies/' 500
```
> This is the final usage of the application. Here, items_count is 500, i.e., any big number, greater than the chart size. The application warns when maximum size is exceeded and limits the execution to available records size.

## Return Data Format :page_with_curl:
The output data is expected to be a collection of JSON objects as follows:
```
[
  {
    "title":"movie_title_here",
    "movie_release_year":"yyyy",
    "imdb_rating":10.0,
    "summary":"Summary of the movie...",
    "duration":"XXh YYmin",
    "genre":"one or more genres seperated by comma"
  },
  .
  .
  .
  {
    ...
  }
]
```

## [Makefile](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/Makefile) :floppy_disk:
GNU Make is used for ease of calling run, build, execute and test commands as above. This is default for many Linux operating systems and can be handy. However, if non-Linux OS or in case of no GNU Make, copy the commands directly to terminal/command prompt.

Syntax: `make COMMAND_NAME`
The possible commands here are:
```
  make run
  make build
  make execute
  make test
  make all
```
> Here, `run`, `build`, `execute`, `test` and `all` are user-defined.

## Sample Execution Output :runner:
```
go run ./app/main.go 'https://www.imdb.com/india/top-rated-indian-movies/' 2
[{"title":"Pather Panchali","movie_release_year":"1955","imdb_rating":8.5,"summary":"Impoverished priest Harihar Ray, dreaming of a better life for himself and his family, leaves his rural Bengal village in search of work.","duration":"2h 5min","genre":"Drama"},{"title":"Drishyam 2","movie_release_year":"2021","imdb_rating":8.5,"summary":"A gripping tale of an investigation and a family which is threatened by it. Will Georgekutty be able to protect his family this time?","duration":"2h 32min","genre":"Drama, Thriller"}]
```

## Key Features :key: :trollface:
`Concurrency` is used in [controller.go](https://github.com/VagueCoder/IMDb-Chart-Fetcher/blob/master/app/fetcher/scrapers/controller.go) for scrapeMovieDetails() method to achieve better performance from the application. The scrapeMovieDetails() method is the one calls 7 scrapers for n individual items for collect and send to encoder. By moving the calls (tested for 250 goroutines so far as that was the max chart size) makes the execution a lot easier and returns results much quicker (39 seconds for 250 items).

The serialization is achieved with the help of rank that the chart has. But to manage the goroutine, `mutex` and `waitgroups` are used rather than buffered channels.
