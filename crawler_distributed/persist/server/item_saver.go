package main

import (
	"flag"
	"fmt"
	"github.com/fisher310/goder/crawler_distributed/config"
	"github.com/fisher310/goder/crawler_distributed/persist"
	"github.com/fisher310/goder/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

var port = flag.Int("port", 0, "the item saver port")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Printf("must specify a port for the item saver")
		return
	}
	log.Fatal(serveRpc(
		fmt.Sprintf(":%d", *port),
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
