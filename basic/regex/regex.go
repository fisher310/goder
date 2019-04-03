package main

import (
	"fmt"
	"regexp"
)

const text = "my email is ccmouse@gmail.com aa@mail.com asdf@mai.jj.cc"

func main() {

	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-z0-9]+)\.([a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)

	for _, m := range match {
		fmt.Println(m)
	}
}
