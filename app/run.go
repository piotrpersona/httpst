package app

import (
	"fmt"

	"github.com/piotrpersona/httpst/data"
)

func displayCodeInfo(codes []string, data data.HTTPData, long, group bool) {
	for _, codeString := range codes {
		code, _ := data.RetrieveCode(codeString)
		fmt.Print(code.Summary(long))
		if group {
			codeGroup, _ := code.Group()
			fmt.Printf("\n%s", codeGroup)
		}
	}
}

func Run(codes []string, applicationConfig Config) {
	data := data.GetHttpInfo()
	group := applicationConfig.Group
	long := applicationConfig.Long
	displayCodeInfo(codes, data, long, group)
}
