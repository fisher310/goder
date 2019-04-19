package main

import (
	"github.com/fisher310/goder/crawler/engine"
	"github.com/fisher310/goder/crawler/model"
	"github.com/fisher310/goder/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	// start ItemSaverServer
	go serveRpc(host, "test1")

	time.Sleep(time.Second)
	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)

	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Url:  "http://www.zhenai.com/biejing/123131",
		Id:   "123131",
		Type: "zhenai",
		Payload: model.Profile{
			Name:     "非诚勿扰",
			Gender:   "男士",
			Age:      30,
			Height:   180,
			Weight:   66,
			Income:   "1.2-2万",
			Hukou:    "四川阿坝",
			Marriage: "未婚",
			Xinzuo:   "天秤座",
			Car:      "已买车",
		},
	}
	var result string
	err = client.Call("ItemSaverService.Save", item, &result)

	if err != nil || result != "ok" {
		t.Errorf("save error, result: %s; err: %s", result, err)
	}
}
