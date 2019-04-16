package parse

import (
	"io"
	"log"
	"strings"

	"github.com/piotrpersona/httpst/schema"
	"golang.org/x/net/html"
)

// Source will parse provided Reader interface and convert it to schema.HTTPData.
func Source(sourceBody io.Reader) (data schema.HTTPData) {
	tokenizer := html.NewTokenizer(sourceBody)

	data = schema.HTTPData{
		Groups: []schema.HTTPGroup{},
		Codes:  []schema.HTTPCode{},
	}

	var latestHTTPCode *schema.HTTPCode
	var latestHTTPpGroup *schema.HTTPGroup
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
					if latestHTTPpGroup != nil {
						data.Groups = append(data.Groups, *latestHTTPpGroup)
					}
					groupTitleSplitted := strings.Split(sectionTitle, " ")
					groupPrefix := groupTitleSplitted[len(groupTitleSplitted)-1]
					latestHTTPpGroup = &schema.HTTPGroup{
						Prefix:      groupPrefix,
						Title:       sectionTitle,
						Description: "",
					}
					latestDescription = &latestHTTPpGroup.Description
				case 2: // status code
					if latestHTTPCode != nil {
						data.Codes = append(data.Codes, *latestHTTPCode)
					}
					splittedTitle := strings.SplitAfterN(sectionTitle, " ", 2)
					latestHTTPCode = &schema.HTTPCode{
						Code:        strings.TrimSpace(splittedTitle[0]),
						Title:       splittedTitle[1],
						Description: "",
					}
					latestDescription = &latestHTTPCode.Description
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
