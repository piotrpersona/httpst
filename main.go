package main

import (
	"fmt"
	"os"

	"github.com/piotrpersona/httpst/schema"
)

func displayCodeInfo(codes []string, data schema.HTTPData) {
	for _, code := range codes {
		httpCode, err := data.RetrieveCode(code)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s %s\n", httpCode.Code, httpCode.Title)
		fmt.Printf("Description:%s\n", httpCode.Description)
	}
}

func main() {
	var dataPath = "/var/lib/httpst/http.json"
	var data = schema.ReadSchema(dataPath)
	var codes = os.Args[1:]
	displayCodeInfo(codes, data)
}
