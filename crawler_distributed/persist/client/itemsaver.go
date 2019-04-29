package client

import (
	"github.com/fisher310/goder/crawler/engine"
	"github.com/fisher310/goder/crawler_distributed/config"
	"github.com/fisher310/goder/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			// Call RPC to save item
			var result string
			err := client.Call(config.ItemSaverRPC, item, &result)
			if err != nil {
				log.Printf("Item Saver: err saving item %v", err)
			}
		}
	}()
	return out, nil
}
