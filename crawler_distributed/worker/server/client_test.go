package main

import (
	"fmt"
	"github.com/fisher310/goder/crawler_distributed/config"
	"github.com/fisher310/goder/crawler_distributed/rpcsupport"
	"github.com/fisher310/goder/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://m.zhenai.com/u/1320662004",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			ARgs: "不诚请勿扰",
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Errorf("%v", err)
	} else {
		fmt.Println(result)
	}
}
