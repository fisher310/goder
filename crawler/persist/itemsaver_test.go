package persist

import (
	"context"
	"encoding/json"
	"github.com/fisher310/goder/crawler/engine"
	"github.com/fisher310/goder/crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSave(t *testing.T) {

	client, err := elastic.NewClient(elastic.SetURL("http://10.252.19.55:9200"), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	var expected = engine.Item{
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

	const index = "dating_test"
	err = save(client, index, expected)

	if err != nil {
		panic(err)
	}

	// TODO: try to start up a elastic search
	// here use docker go client
	resp, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	var actual engine.Item

	err = json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	if actual != expected {
		t.Errorf("got %+v, expected: %+v", actual, expected)
	}
}
