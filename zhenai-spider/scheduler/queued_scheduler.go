package scheduler

import (
	"zhenai-spider/fetcher"
)

type QueuedScheduler struct {
	fetcherIn       chan fetcher.Fetcher
	workerChannelIn chan chan fetcher.Fetcher
}

func (s *QueuedScheduler) Run() {
	s.fetcherIn = make(chan fetcher.Fetcher)
	s.workerChannelIn = make(chan chan fetcher.Fetcher)

	var fetcherQueue []fetcher.Fetcher
	var workerChannelQueue []chan fetcher.Fetcher

	go func() {
		for {
			if len(fetcherQueue) > 0 && len(workerChannelQueue) > 0 {
				f := fetcherQueue[0]
				fetcherQueue = fetcherQueue[1:]

				w := workerChannelQueue[0]
				workerChannelQueue = workerChannelQueue[1:]

				w <- f
			}

			select {
			case fetcher := <-s.fetcherIn:
				fetcherQueue = append(fetcherQueue, fetcher)
			case workerChannel := <-s.workerChannelIn:
				workerChannelQueue = append(workerChannelQueue, workerChannel)
			}
		}
	}()
}

func (s *QueuedScheduler) Submit(f fetcher.Fetcher) {
	s.fetcherIn <- f
}

func (s *QueuedScheduler) FetchInChannel() chan fetcher.Fetcher {
	return make(chan fetcher.Fetcher)
}

func (s *QueuedScheduler) WorkerReady(workerChannel chan fetcher.Fetcher) {
	s.workerChannelIn <- workerChannel
}
