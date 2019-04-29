package client

import (
	"github.com/fisher310/goder/crawler/engine"
	"github.com/fisher310/goder/crawler_distributed/config"
	"github.com/fisher310/goder/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	//client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	//if err != nil {
	//	return nil, err
	//}

	return func(r engine.Request) (result engine.ParseResult, e error) {
		sReq := worker.SerializeRequest(r)

		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(sResult), nil
	}
}
