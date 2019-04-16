package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/piotrpersona/httpst/parse"
	"github.com/piotrpersona/httpst/schema"

	"golang.org/x/net/html"
)

func Read(filename string) (hd schema.HttpData) {
	jsonBytes, err := os.Open(filename)
	defer jsonBytes.Close()
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := ioutil.ReadAll(jsonBytes)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonData, &hd)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func fetchStatusesSource() io.Reader {
	const httpStatusCodeSourceURL = "https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html"
	response, err := http.Get(httpStatusCodeSourceURL)
	if err != nil {
		log.Fatal(err)
	}
	return response.Body
}

func fetchData() (data schema.HttpData) {
	responseBody := fetchStatusesSource()
	tokenizer := html.NewTokenizer(responseBody)

	data = schema.HttpData{
		Groups: []schema.HttpGroup{},
		Codes:  []schema.HttpCode{},
	}

	var latestHttpCode *schema.HttpCode
	var latestHttpGroup *schema.HttpGroup
	var latestDescription *string

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			log.Fatalf("Malformed HTML")
		}
		if tokenType == html.StartTagToken {
			// var description string
			switch token := tokenizer.Token(); token.Data {
			case "h3":
				tokenizer.Next()
				tokenizer.Next()
				sectionNumber := tokenizer.Token().Data
				tokenizer.Next()
				tokenizer.Next()
				sectionTitle := tokenizer.Token().Data
				sectionTitle = strings.TrimSpace(sectionTitle)
				numberOfPeriods := strings.Count(sectionNumber, ".")
				switch numberOfPeriods {
				case 1: // group
					if latestHttpGroup != nil {
						data.Groups = append(data.Groups, *latestHttpGroup)
					}
					groupTitleSplitted := strings.Split(sectionTitle, " ")
					groupPrefix := groupTitleSplitted[len(groupTitleSplitted)-1]
					latestHttpGroup = &schema.HttpGroup{
						Prefix:      groupPrefix,
						Title:       sectionTitle,
						Description: "",
					}
					latestDescription = &latestHttpGroup.Description
				case 2: // status code
					if latestHttpCode != nil {
						data.Codes = append(data.Codes, *latestHttpCode)
					}
					splittedTitle := strings.SplitAfterN(sectionTitle, " ", 2)
					latestHttpCode = &schema.HttpCode{
						Code:        strings.TrimSpace(splittedTitle[0]),
						Title:       splittedTitle[1],
						Description: "",
					}
					latestDescription = &latestHttpCode.Description
				}
			case "p":
				tokenizer.Next()
				description := tokenizer.Token().Data
				if latestDescription != nil {
					*latestDescription += description
				}
			}
		}
	}
	return
}

func retrieveCode(requestedCode string, data schema.HttpData) (httpCode schema.HttpCode, err error) {
	for _, code := range data.Codes {
		if code.Code == requestedCode {
			httpCode = code
			return
		}
	}
	err = fmt.Errorf("No such code: %s", requestedCode)
	return
}

func displayCodeInfo(codes []string, data schema.HttpData) {
	for _, code := range codes {
		httpCode, err := retrieveCode(code, data)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %s\n", httpCode.Code, httpCode.Title)
		fmt.Printf("%s\n", httpCode.Description)
	}
}

func main() {
	responseBody := fetchStatusesSource()
	data := parse.Source(responseBody)
	data.Dump("http.json")
	// data := Read("http.json")
	// codes := os.Args[1:]
	// displayCodeInfo(codes, data)
}
