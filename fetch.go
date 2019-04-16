package main

import (
	"io"
	"log"
	"net/http"

	"github.com/piotrpersona/httpst/parse"
)

func fetchStatusesSource() io.Reader {
	const httpStatusCodeSourceURL = "https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html"
	response, err := http.Get(httpStatusCodeSourceURL)
	if err != nil {
		log.Fatal(err)
	}
	return response.Body
}

func main() {
	responseBody := fetchStatusesSource()
	data := parse.Source(responseBody)
	data.Dump("http.json")
}
