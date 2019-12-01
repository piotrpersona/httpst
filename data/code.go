package data

import "fmt"

// HTTPCode describes particular http status code.
type HTTPCode struct {
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (hc *HTTPCode) Summary(long bool) (summary string) {
	basicInfo := fmt.Sprintf("%s %s\n", hc.Code, hc.Title)
	summary = basicInfo
	if long {
		summary += fmt.Sprintf("\n%s\n", hc.Description)
	}
	return
}

func (hc *HTTPCode) Group() (HTTPGroup, error) {
	groups := GetHttpInfo().Groups
	for _, group := range groups {
		codeStart := fmt.Sprintf("%c", hc.Code[0])
		groupStart := fmt.Sprintf("%c", group.Prefix[0])
		if codeStart == groupStart {
			return group, nil
		}
	}
	return HTTPGroup{}, fmt.Errorf("No such groups for code: %s", hc.Code)
}
