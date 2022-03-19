package main

import (
	"crawler/internal/api/httpserver"
	v1 "crawler/internal/api/v1"
	"crawler/internal/crawler"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

const (
	defaultPort = "9999"
	defaultHost = "0.0.0.0"
)

func main() {
	port, ok := os.LookupEnv("CRAWLER_PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("CRAWLER_HOST")
	if !ok {
		host = defaultHost
	}

	if err := execute(net.JoinHostPort(host, port)); err != nil {
		os.Exit(1)
	}
}

func execute(addr string) error {

	crawler := crawler.NewCrawlerService()
	for i := 0; i < 10; i++ {
			go crawler.CrawlerWorker()
	}
	crawlerController := v1.NewCrawlerController(crawler)

	router := httpserver.NewRouter(chi.NewRouter(), crawlerController)
	server := http.Server{
		Addr: addr,
		Handler: &router,
	}
	return server.ListenAndServe()
}
