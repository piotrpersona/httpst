package data

import (
	"fmt"
)

// HTTPData contains set of all available Http codes and groups.
type HTTPData struct {
	Groups []HTTPGroup `json:"groups"`
	Codes  []HTTPCode  `json:"codes"`
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
