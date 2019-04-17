package parser

import (
	"github.com/fisher310/goder/crawler/engine"
	"github.com/fisher310/goder/crawler/model"
	"regexp"
	"strconv"
)

var (
	ageRe      = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([0-9]+)岁</div>`)
	genderRe   = regexp.MustCompile(`"genderString":"([^"]+)"`)
	marriageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([已未]婚|离异)</div>`)
	xinzuoRe   = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(.{2}座)\([0-9.-]+\)</div>`)
	heightRe   = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([0-9]+)cm</div>`)
	weightRe   = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([0-9]+)kg</div>`)
	incomeRe   = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>月收入:([0-9.\-]+万?)</div>`)
	hukouRe    = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>籍贯:([^<]+)</div>`)
	houseRe    = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>([^<]+房)</div>`)
	carRe      = regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>([^<]+车)</div>`)
	idUrlRe    = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
)

func ParseProfile(contents []byte, name string, url string) engine.ParseResult {

	profile := model.Profile{}
	profile.Name = name
	profile.Gender = extractString(contents, genderRe)
	if age, err := strconv.Atoi(extractString(contents, ageRe)); err == nil {
		profile.Age = age
	}

	profile.Marriage = extractString(contents, marriageRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)

	if height, err := strconv.Atoi(extractString(contents, heightRe)); err == nil {
		profile.Height = height
	}

	if weightRe, err := strconv.Atoi(extractString(contents, weightRe)); err == nil {
		profile.Weight = weightRe
	}

	profile.Income = extractString(contents, incomeRe)
	profile.Hukou = extractString(contents, hukouRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{Url: url, Type: "zhenai", Id: extractString([]byte(url), idUrlRe), Payload: profile},
		},
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
