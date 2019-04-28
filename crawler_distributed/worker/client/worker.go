package client

import (
	"fmt"
	"github.com/fisher310/goder/crawler/engine"
	"github.com/fisher310/goder/crawler_distributed/config"
	"github.com/fisher310/goder/crawler_distributed/rpcsupport"
	"github.com/fisher310/goder/crawler_distributed/worker"
)

func CreateProcessor()  (engine.Processor, error){

	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}

	return func(r engine.Request) (result engine.ParseResult, e error) {
		sReq := worker.SerializeRequest(r)

		var sResult worker.ParseResult
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(sResult), nil
	}, nil
}
