package main

import (
	"fmt"
	"github.com/fisher310/goder/crawler/engine"
	"github.com/fisher310/goder/crawler/scheduler"
	"github.com/fisher310/goder/crawler/zhenai/parser"
	"github.com/fisher310/goder/crawler_distributed/config"
	itemSaver "github.com/fisher310/goder/crawler_distributed/persist/client"
	worker "github.com/fisher310/goder/crawler_distributed/worker/client"
)

func main() {
	itemChan, err := itemSaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(nil)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.CreateFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}
