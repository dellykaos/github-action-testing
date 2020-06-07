package main

import (
	"log"
	"net/http"

	"github.com/dellykaos/github-action-testing/handler"
	"github.com/dellykaos/github-action-testing/handler/healthz"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	healthz := healthz.New()
	h := handler.New()

	router.GET("/healthz", healthz.Healthz)
	router.GET("/", h.Hello)
	router.GET("/hello", h.Hello)
	router.GET("/hello/:name", h.Hello)

	// Start server
	log.Println("Listening at port", 3200)
	log.Fatal(http.ListenAndServe(":3200", router))
}
