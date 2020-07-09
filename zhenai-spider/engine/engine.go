package engine

import "zhenai-spider/fetcher"

type engine interface {
	Run(seeds ...fetcher.Fetcher)
}
