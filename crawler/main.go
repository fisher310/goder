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
	"regexp"
)

const (
	url = "http://www.zhenai.com/zhenghun"
)

func main() {
	//resp, err := http.Get("http://www.zhenai.com/zhenghun")

	request, err := http.NewRequest(http.MethodGet, url, nil)

	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.162 Safari/537.36")
	request.Header.Add("Host", "www.zhenai.com")

	resp, err := http.DefaultClient.Do(request)

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
	fmt.Println("----------------------------------------------")
	printCityList(all)
}

func determineEncoding(r io.Reader) (*bufio.Reader, encoding.Encoding) {
	reader := bufio.NewReader(r)
	bytes, err := reader.Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return reader, e
}

func printCityList(contents []byte) {
	//re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenhun/[0-9a-z]+" [^>]*>[^<]+</a>`)
	re := regexp.MustCompile(`{linkContent:"([^,]+)",linkURL:"(http://m.zhenai.com/zhenghun/[a-z]+)"}`)
	matches := re.FindAllSubmatch(contents, -1)
	//matches := re.FindAll(contents, -1)
	fmt.Printf("%s\n", matches)

	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s\n", m[1], m[2])
	}

}
