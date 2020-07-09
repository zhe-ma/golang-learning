package main

import (
	"zhenai-spider/engine"
	"zhenai-spider/fetcher"
	"zhenai-spider/scheduler"
)

func main() {
	// simpleEngine := engine.SimpleEngine{}
	// simpleEngine.Run(&fetcher.CityFetcher{URL: "http://www.zhenai.com/zhenghun"})

	scheduler := &scheduler.SimpleScheduler{}

	concurrentEngine := engine.ConcurrentEngine{WorkerCount: 10, Scheduler: scheduler}
	concurrentEngine.Run(&fetcher.CityFetcher{URL: "http://www.zhenai.com/zhenghun"})
}
