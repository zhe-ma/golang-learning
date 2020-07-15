package engine

import (
	"zhenai-spider/database"
	"zhenai-spider/fetcher"
	"zhenai-spider/model"
	"zhenai-spider/scheduler"
	"zhenai-spider/util"
)

type ConcurrentEngine struct {
	WorkerCount int
	Scheduler   scheduler.Scheduler
	DataSaver   chan database.ElasticItem
}

func (e *ConcurrentEngine) Run(seeds ...fetcher.Fetcher) {
	util.TraceLog.Println("Engine Run")

	resultOut := make(chan fetcher.Result)

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.FetchInChannel(), e.Scheduler, resultOut)
	}

	for _, seed := range seeds {
		go func(s fetcher.Fetcher) {
			e.Scheduler.Submit(s)
		}(seed)
	}

	for {
		result := <-resultOut

		for _, item := range result.Items {

			if profiles, ok := item.(*[]model.Profile); ok {
				for _, profile := range *profiles {
					util.InfoLog.Println("Spide profile:", profile)

					if memberExist(profile.MemberId) {
						util.InfoLog.Println("Member eixsts, ID:", profile.MemberId)
						continue
					}

					elasticItem := database.ElasticItem{Index: "zhenai", Type: "profiles", Data: profile}

					if e.DataSaver != nil {
						e.DataSaver <- elasticItem
					}
				}
			}
		}

		subFetchers := result.SubFetchers
		for _, subFetcher := range subFetchers {
			if f, ok := subFetcher.(*fetcher.ProfileFetcher); ok {
				if urlExist(f.URL) {
					continue
				}
			}

			e.Scheduler.Submit(subFetcher)
		}
	}
}

func createWorker(fetcherIn chan fetcher.Fetcher, s scheduler.Scheduler, resultOut chan fetcher.Result) {
	go func() {
		for {
			s.WorkerReady(fetcherIn)

			fetcher := <-fetcherIn

			result := fetcher.Run()
			resultOut <- result
		}
	}()
}

var memberIdCache = make(map[float64]bool)

func memberExist(id float64) bool {
	if memberIdCache[id] {
		return true
	}

	memberIdCache[id] = true
	return false
}

var urlCache = make(map[string]bool)

func urlExist(url string) bool {
	if urlCache[url] {
		return true
	}

	urlCache[url] = true
	return false
}
