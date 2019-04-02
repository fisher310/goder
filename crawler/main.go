package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error")
		return
	}

	reader, e := determineEncoding(resp.Body)

	utf8Reader := transform.NewReader(reader, e.NewEncoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", all)
}


func determineEncoding(r io.Reader) (*bufio.Reader, encoding.Encoding) {
	reader := bufio.NewReader(r)
	bytes, err := reader.Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _:= charset.DetermineEncoding(bytes, "")
	return reader, e
}