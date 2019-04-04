package main

import (
	"github.com/fisher310/goder/crawler/engine"
	"github.com/fisher310/goder/crawler/zhenai/parser"
)

const (
	url = "http://www.zhenai.com/zhenghun"
)

func main() {

	engine.Run(engine.Request{Url: url, ParserFunc: parser.ParseCityList})

}
