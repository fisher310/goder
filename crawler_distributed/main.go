package main

import (
	"flag"
	"fmt"
	"github.com/fisher310/goder/crawler/engine"
	"github.com/fisher310/goder/crawler/scheduler"
	"github.com/fisher310/goder/crawler/zhenai/parser"
	"github.com/fisher310/goder/crawler_distributed/config"
	itemSaver "github.com/fisher310/goder/crawler_distributed/persist/client"
	"github.com/fisher310/goder/crawler_distributed/rpcsupport"
	worker "github.com/fisher310/goder/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()

	itemChan, err := itemSaver.ItemSaver(fmt.Sprintf(":%s", *itemSaverHost))
	if err != nil {
		panic(nil)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.CreateFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}


func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client

	for _, h := range hosts {
		client, err := rpcsupport.NewClient(fmt.Sprintf(":%s",h))
		if err == nil {
			clients = append(clients, client)
			log.Printf("connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}