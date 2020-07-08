package main

import (
	"zhenai-spider/engine"
	"zhenai-spider/fetcher"
)

func main() {
	engine.Run(&fetcher.CityFetcher{URL: "http://www.zhenai.com/zhenghun"})
}
