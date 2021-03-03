package main

import (
	"log"
	"os"
	"strconv"

	"github.com/VagueCoder/IMDb-Chart-Fetcher/app/fetcher"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Insufficient arguments provided. Require chart_url and items_count.")
	}
	chartURL := os.Args[1]
	itemsCount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid number %q error: %v", os.Args[2], err)
	}
	fetcher.FetchItems(chartURL, itemsCount, os.Stdout)
}
