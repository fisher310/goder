package main

import (
	"github.com/fisher310/goder/crawler/engine"
	"github.com/fisher310/goder/crawler/persist"
	"github.com/fisher310/goder/crawler/scheduler"
	"github.com/fisher310/goder/crawler/zhenai/parser"
)

const (
	url     = "http://www.zhenai.com/zhenghun"
	cityUrl = "http://www.zhenai.com/zhenghun/guangxi"
)

func main() {

	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{Url: url, ParserFunc: parser.ParseCityList})
	//e.Run(engine.Request{
	//	Url:        cityUrl,
	//	ParserFunc: parser.ParseCity,
	//})

}
