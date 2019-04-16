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
	var latestHTTPGroup *schema.HTTPGroup
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
				var sectionHeader = Section(tokenizer)
				var numberOfPeriods = strings.Count(sectionHeader.Number, ".")
				switch numberOfPeriods {
				case 1: // group
					if latestHTTPGroup != nil {
						data.Groups = append(data.Groups, *latestHTTPGroup)
					}
					var groupTitleSplitted = strings.Split(sectionHeader.Title, " ")
					var groupPrefix = groupTitleSplitted[len(groupTitleSplitted)-1]
					latestHTTPGroup = &schema.HTTPGroup{
						Prefix:      groupPrefix,
						Title:       sectionHeader.Title,
						Description: "",
					}
					latestDescription = &latestHTTPGroup.Description
				case 2: // status code
					if latestHTTPCode != nil {
						data.Codes = append(data.Codes, *latestHTTPCode)
					}
					var splittedTitle = strings.SplitAfterN(sectionHeader.Title, " ", 2)
					latestHTTPCode = &schema.HTTPCode{
						Code:        strings.TrimSpace(splittedTitle[0]),
						Title:       splittedTitle[1],
						Description: "",
					}
					latestDescription = &latestHTTPCode.Description
				}
			case "p":
				tokenizer.Next()
				var description = tokenizer.Token().Data
				if latestDescription != nil {
					*latestDescription += description
				}
			}
		}
	}
	return
}
