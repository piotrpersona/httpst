package schema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// HTTPCode describes particular http status code.
type HTTPCode struct {
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

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

// HTTPData contains set of all available Http codes and groups.
type HTTPData struct {
	Groups []HTTPGroup `json:"groups"`
	Codes  []HTTPCode  `json:"codes"`
}

// Dump will marshal data to json file.
func (hd *HTTPData) Dump(filename string) {
	jsonBytes, err := json.Marshal(hd)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(filename, jsonBytes, 0644)
}

// RetrieveCode will fetch information about particular code.
func (hd *HTTPData) RetrieveCode(code string) (hc HTTPCode, err error) {
	for _, httpCode := range hd.Codes {
		if httpCode.Code == code {
			hc = httpCode
			return
		}
	}
	err = fmt.Errorf("No such code: %s", code)
	return
}
