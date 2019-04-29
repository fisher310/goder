package parser

import (
	"github.com/fisher310/goder/crawler/engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

//var moreCity = regexp.MustCompile(`<a target="_blank" href=("http://www.zhenai.com/zhenghun/guangxi/[a-z]+")>`)

func ParseCity(contents []byte, _ string) engine.ParseResult {

	matches := cityRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		//result.Items = append(result.Items, "User: "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: CreateProfileParser(string(m[2])),
		})
	}

	return result
}

type ProfileParser struct {
	username string
}

func (p *ProfileParser) Parse(content []byte, url string) engine.ParseResult {
	return ParseProfile(content, url, p.username)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.username
}

func CreateProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		username: name,
	}
}
