package worker

import (
	"errors"
	"fmt"
	"github.com/fisher310/goder/crawler/engine"
	"github.com/fisher310/goder/crawler/zhenai/parser"
	"github.com/fisher310/goder/crawler_distributed/config"
	"log"
)

type SerializedParser struct {
	Name string
	ARgs interface{}
}

//

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			ARgs: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	p, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: p,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.CreateFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.CreateFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if userName, ok := p.ARgs.(string); ok {
			return parser.CreateProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.ARgs)
		}
	default:
		return nil, errors.New("unknown parser name " + p.Name)
	}
}
