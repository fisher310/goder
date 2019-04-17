package frontpage

import (
	"github.com/fisher310/goder/crawler/engine"
	pageModel "github.com/fisher310/goder/crawler/frontpage/model"
	"github.com/fisher310/goder/crawler/model"
	"html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	tmp, err := template.ParseFiles("template.html")

	if err != nil {
		panic(err)
	}

	out, err := os.Create("template.test.html")

	page := pageModel.SearchResult{}
	page.Hits = 123
	page.Start = 1

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
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}
	err = tmp.Execute(out, page)
	if err != nil {
		panic(err)
	}
}
