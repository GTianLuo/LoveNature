package conf

import (
	"context"
	"github.com/olivere/elastic/v7"
)

var (
	esClient *elastic.Client
)

func es(url string) {
	client, _ := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
	if _, _, err := client.Ping(url).Do(context.Background()); err != nil {
		panic(err)
	}
	esClient = client
}

func NewEs() *elastic.Client {
	return esClient
}
