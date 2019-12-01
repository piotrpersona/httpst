package data

import "fmt"

// HTTPGroup describes particular http group
// There are 5 types of groups:
// 1xx - Informational
// 2xx - Success
// 3xx - Redirection
// 4xx - Client Error
// 5xx - Server Error
type HTTPGroup struct {
	Prefix      string `json:"prefix"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (hg HTTPGroup) String() string {
	return fmt.Sprintf("%s %s\n%s", hg.Prefix, hg.Title, hg.Description)
}
