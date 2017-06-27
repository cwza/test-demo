package main

import (
	"log"
	"net/http"

	"os"

	"github.com/cwza/test-demo/pkg/web"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	port = ":" + port
	if !ok {
		port = ":10020"
	}
	log.Fatal(http.ListenAndServe(port, web.Router))
}
