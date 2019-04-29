package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	pr := newPathResolver()

	pr.Add("GET /hello", hello)
	pr.Add("(GET|HEAD) /goodbye(/?[A-Za-z0-9]*)?", goodbye)
	pr.Add("GET /", homePage)

	log.Fatal(http.ListenAndServe(":8080", pr))
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
	cache    map[string]*regexp.Regexp
}

func newPathResolver() *pathResolver {
	return &pathResolver{
		handlers: make(map[string]http.HandlerFunc),
		cache:    make(map[string]*regexp.Regexp),
	}
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
	p.cache[path] = regexp.MustCompile(path)
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path
	for pattern, handlerFunc := range p.handlers {
		if p.cache[pattern].MatchString(check) {
			handlerFunc(res, req)
			return
		}
	}
	http.NotFound(res, req)
}

func hello(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	_, _ = fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
	p := req.URL.Path
	parts := strings.Split(p, "/")
	name := ""
	if len(parts) > 2 {
		name = parts[2]
	}
	if name == "" {
		name = "Inigo Montoya"
	}
	_, _ = fmt.Fprint(res, "Goodbye ", name)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	_, _ = fmt.Fprint(res, "The HomePage.")
}
