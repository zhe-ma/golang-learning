package scheduler

import "zhenai-spider/fetcher"

type SimpleScheduler struct {
	fetcherIn chan fetcher.Fetcher
}

func (s *SimpleScheduler) Run() {
	s.fetcherIn = make(chan fetcher.Fetcher)
}

func (s *SimpleScheduler) Submit(f fetcher.Fetcher) {
	s.fetcherIn <- f
}

func (s *SimpleScheduler) FetchInChannel() chan fetcher.Fetcher {
	return s.fetcherIn
}

func (s *SimpleScheduler) WorkerReady(workerChannel chan fetcher.Fetcher) {
}
