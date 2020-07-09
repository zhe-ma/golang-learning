package engine

import (
	"fmt"
	"zhenai-spider/fetcher"
	"zhenai-spider/scheduler"
	"zhenai-spider/util"
)

type ConcurrentEngine struct {
	WorkerCount int
	Scheduler   scheduler.Scheduler
}

func (e *ConcurrentEngine) Run(seeds ...fetcher.Fetcher) {
	util.TraceLog.Println("Engine Run")

	resultOut := make(chan fetcher.Result)

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.FetchInChannel(), resultOut)
	}

	for _, seed := range seeds {
		go func(s fetcher.Fetcher) {
			e.Scheduler.Submit(s)
		}(seed)
	}

	for {
		result := <-resultOut

		for _, item := range result.Items {
			fmt.Printf("item : %s\n", item)
		}

		subFetchers := result.SubFetchers
		for _, fetcher := range subFetchers {
			e.Scheduler.Submit(fetcher)
		}
	}
}

func createWorker(fetcherIn chan fetcher.Fetcher, resultOut chan fetcher.Result) {
	go func() {
		fetcher := <-fetcherIn

		result := fetcher.Run()
		resultOut <- result
	}()
}
