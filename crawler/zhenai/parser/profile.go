package parser

import (
	"github.com/fisher310/goder/crawler/engine"
	"github.com/fisher310/goder/crawler/model"
	"regexp"
	"strconv"
)

var (
	ageRe      = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
	marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
)

func ParseProfile(contents []byte) engine.ParseResult {

	profile := model.Profile{}
	if age, err := strconv.Atoi(extractString(contents, ageRe)); err == nil {
		profile.Age = age
	}

	profile.Marriage = extractString(contents, marriageRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
