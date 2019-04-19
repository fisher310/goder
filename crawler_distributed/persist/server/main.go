package main

import (
	"github.com/fisher310/goder/crawler_distributed/persist"
	"github.com/fisher310/goder/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func main() {
	log.Fatal(serveRpc(":1234", "dating_profile"))
}

func serveRpc(host string, index string) error {

	client, err := elastic.NewClient(elastic.SetURL("http://10.252.19.55:9200"), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
