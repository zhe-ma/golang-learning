package engine

import (
	"fmt"
	"zhenai-spider/database"
	"zhenai-spider/fetcher"
	"zhenai-spider/model"
	"zhenai-spider/util"
)

type SimpleEngine struct {
	DataSaver chan database.ElasticItem
}

func (e *SimpleEngine) Run(seeds ...fetcher.Fetcher) {
	util.TraceLog.Println("Engine Run")

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

			if profiles, ok := item.(*[]model.Profile); ok {
				for _, profile := range *profiles {
					fmt.Println(profile)
					elasticItem := database.ElasticItem{Index: "zhenai", Type: "profiles", Data: profile}
					e.DataSaver <- elasticItem
				}
			}
		}
	}
}
