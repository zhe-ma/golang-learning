package scheduler

import "zhenai-spider/fetcher"

type Scheduler interface {
	Run()
	FetchInChannel() chan fetcher.Fetcher
	Submit(f fetcher.Fetcher)
	WorkerReady(workerChannel chan fetcher.Fetcher)
}
