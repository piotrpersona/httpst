package schema

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// ReadSchema will unmarshal json file to HTTPData object.
func ReadSchema(filename string) (hd HTTPData) {
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
