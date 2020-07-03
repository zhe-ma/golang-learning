package engine

import "fmt"

func Run(seeds ...*Fetcher) {
	fmt.Println("Engine Run")

	fetchers := make([]*Fetcher)
	for _, seed := range seeds {
		fetchers = append(fetcher, seed)
	}
}
