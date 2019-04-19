package main

import (
	"fmt"
	"github.com/braintree/manners"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	handler := newHandler()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	log.Fatal(manners.ListenAndServe(":8080", handler))
}

func newHandler() *handler {
	return &handler{}
}

type handler struct {
}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	_, _ = fmt.Fprint(res, "Hello, my name is ", name)
}

func listenForShutdown(ch <-chan os.Signal) {
	<-ch
	log.Println("will be shutdown...")
	manners.Close()
}
