package engine

import "fmt"

func Run(seeds ...*Fetcher) {
	fmt.Println("Engine Run")

	// fetchers := make([]*Fetcher)
	fetchers := []*Fetcher{}

	for _, seed := range seeds {
		fetchers = append(fetchers, seed)
	}

	for len(fetchers) > 0 {
		fetcher := fetchers[0]
		fetchers = fetchers[1:]

		fetcher.Run()
	}
}
