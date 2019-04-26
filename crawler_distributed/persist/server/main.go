package main

import (
	"fmt"
	"github.com/fisher310/goder/crawler_distributed/config"
	"github.com/fisher310/goder/crawler_distributed/persist"
	"github.com/fisher310/goder/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func main() {
	log.Fatal(serveRpc(
		fmt.Sprintf(":%d", config.ItemSaverPort),
		config.ElasticIndex))
}

func serveRpc(host string, index string) error {

	client, err := elastic.NewClient(elastic.SetURL(config.ElasticAddress), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
