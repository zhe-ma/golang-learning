package main

import (
	"zhenai-spider/database"
	"zhenai-spider/engine"
	"zhenai-spider/fetcher"
	"zhenai-spider/scheduler"
	"zhenai-spider/util"
)

func main() {
	util.InitLogging()

	dataSaver := database.NewElasticSaver("http://10.196.102.145:9200")
	// simpleEngine := engine.SimpleEngine{DataSaver: dataSaver}
	// simpleEngine.Run(&fetcher.CityFetcher{URL: "http://www.zhenai.com/zhenghun"})

	// scheduler := &scheduler.SimpleScheduler{}
	scheduler := &scheduler.QueuedScheduler{}

	concurrentEngine := engine.ConcurrentEngine{WorkerCount: 10, Scheduler: scheduler, DataSaver: dataSaver}
	concurrentEngine.Run(&fetcher.CityFetcher{URL: "http://www.zhenai.com/zhenghun"})
}
