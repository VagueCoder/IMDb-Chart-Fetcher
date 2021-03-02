package main

import (
	"log"
	"os"
	"strconv"

	"github.com/VagueCoder/IMDb-Chart-Fetcher/app/fetcher"
)

func main() {
	chartURL := os.Args[1]
	itemsCount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid number %v error: %q", os.Args[2], err)
	}
	fetcher.FetchItems(chartURL, itemsCount, os.Stdout)
}
