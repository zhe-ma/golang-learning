package engine

import (
	"fmt"
	"zhenai-spider/fetcher"
	"zhenai-spider/util"
)

func Run(seeds ...fetcher.Fetcher) {
	util.TraceLog.Println("Engine Run")

	// fetchers := make([]*Fetcher)
	fetchers := []fetcher.Fetcher{}

	for _, seed := range seeds {
		fetchers = append(fetchers, seed)
	}

	for len(fetchers) > 0 {
		fetcher := fetchers[0]
		fetchers = fetchers[1:]

		result := fetcher.Run()
		if result.SubFetchers != nil {
			fetchers = append(fetchers, result.SubFetchers...)
		}

		for _, item := range result.Items {
			fmt.Printf("item : %s\n", item)
		}
	}
}
