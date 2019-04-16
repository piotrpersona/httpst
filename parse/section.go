package parse

import (
	"strings"

	"golang.org/x/net/html"
)

// SectionHeader describes HTTP Group/Code HTML Section.
type SectionHeader struct {
	Number, Title string
}

// Section retrieves SectionHeader from tokenizer.
// The assumption is that tokenizer stores <h3> object
// Tokens structure:
// <h3>  <-- start
// <a>
// Section Number
// </a>
// Section Title
func Section(tokenizer *html.Tokenizer) (section SectionHeader) {
	var sectionNumber, sectionTitle string
	tokenizer.Next() // <a>
	tokenizer.Next() // Number
	sectionNumber = tokenizer.Token().Data
	tokenizer.Next() // </a>
	tokenizer.Next() // Title
	sectionTitle = tokenizer.Token().Data

	sectionNumber = strings.TrimSpace(sectionNumber)
	sectionTitle = strings.TrimSpace(sectionTitle)

	section = SectionHeader{
		Number: sectionNumber,
		Title:  sectionTitle,
	}
	return
}
