package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer func(r io.ReadCloser) {
		err := r.Close()
		if err != nil {
			log.Fatalf("error when close the response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	reader, e := determineEncoding(resp.Body)

	utf8Reader := transform.NewReader(reader, e.NewEncoder())

	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r io.Reader) (*bufio.Reader, encoding.Encoding) {
	reader := bufio.NewReader(r)
	bytes, err := reader.Peek(1024)
	if err != nil {
		log.Printf("Fetch error, %s", err)
		return reader, unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return reader, e
}
