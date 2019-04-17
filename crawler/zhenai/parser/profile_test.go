package parser

import (
	"github.com/fisher310/goder/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	url := "http://m.zhenai.com/u/1320662004"
	result := ParseProfile(contents, "非诚勿扰", url)

	var expected = model.Profile{
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
	}

	if len(result.Items) < 1 {
		t.Errorf("解析错误, 期望解析出正确的Items, 但ParseResult.Items为空")
	}

	profile := result.Items[0].Payload.(model.Profile)

	if profile.Name != expected.Name {
		t.Errorf("name is error, expected %s but the result is %s", expected.Name, profile.Name)
	}

	if profile.Gender != expected.Gender {
		t.Errorf("gender is error, expected is %s but the result is %s", expected.Gender, profile.Gender)
	}
	if profile.Age != expected.Age {
		t.Errorf("age is error, expected is %d but the result is %d", expected.Age, profile.Age)
	}

	if profile.Height != expected.Height {
		t.Errorf("height is error, expected is %d but the result is %d", expected.Height, profile.Height)
	}

	if profile.Weight != expected.Weight {
		t.Errorf("weight is error, expected is %d but the result is %d", expected.Weight, profile.Weight)
	}

	if profile.Income != expected.Income {
		t.Errorf("income is error, expected is %s but the result is %s", expected.Income, profile.Income)
	}

	if profile.Hukou != expected.Hukou {
		t.Errorf("hukou is error, expected is %s but the result is %s", expected.Hukou, profile.Hukou)
	}

	if profile.Marriage != expected.Marriage {
		t.Errorf("marriage is error, expected is %s but the result is %s", expected.Marriage, profile.Marriage)
	}

	if profile.Xinzuo != expected.Xinzuo {
		t.Errorf("xinzuo is error, expected is %s but the result is %s", expected.Xinzuo, profile.Xinzuo)
	}

	if profile.Car != expected.Car {
		t.Errorf("car is error, expected is %s, but the result is %s", expected.Car, profile.Car)
	}

}
