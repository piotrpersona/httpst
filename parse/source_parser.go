package parse

import (
	"io"
	"log"
	"strings"

	"github.com/piotrpersona/httpst/schema"
	"golang.org/x/net/html"
)

func Source(sourceBody io.Reader) (data schema.HttpData) {
	tokenizer := html.NewTokenizer(sourceBody)

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
