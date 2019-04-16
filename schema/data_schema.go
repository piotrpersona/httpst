package schema

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type HttpCode struct {
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type HttpGroup struct {
	Prefix      string `json:"prefix"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type HttpData struct {
	Groups []HttpGroup `json:"groups"`
	Codes  []HttpCode  `json:"codes"`
}

func (hd *HttpData) Dump(filename string) {
	jsonBytes, err := json.Marshal(hd)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(filename, jsonBytes, 0644)
}
