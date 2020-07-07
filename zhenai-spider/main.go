package main

import "zhenai-spider/engine"

func main() {
	engine.Run(&engine.CityFetcher{"http://www.zhenai.com/zhenghun"})
}
