package main

import (
	"fmt"

	"github.com/VagueCoder/IMDb-Chart-Fetcher/app/fetcher"
)

func main() {
	encoded := fetcher.FetchItems(10)
	fmt.Println(encoded)
}
