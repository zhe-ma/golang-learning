package database

import (
	"context"
	"zhenai-spider/util"

	"github.com/olivere/elastic"
)

type ElasticItem struct {
	Index string
	Type  string
	Data  interface{}
}

func NewElasticSaver(host string) chan ElasticItem {
	client, err := elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false))
	if err != nil {
		util.ErrorLog.Println("Failed connect elastic. Host: ", host)
		return nil
	}

	itemChannel := make(chan ElasticItem)

	go func(c *elastic.Client) {
		for {
			item := <-itemChannel
			if client == nil {
				return
			}

			_, err := client.Index().Index(item.Index).Type(item.Type).BodyJson(item.Data).Do(context.Background())
			if err != nil {
				util.ErrorLog.Println("Failed insert data into Elastic. Error: ", err)
			}
		}
	}(client)

	return itemChannel
}
