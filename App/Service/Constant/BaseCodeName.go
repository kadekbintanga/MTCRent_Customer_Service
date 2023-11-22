package Constant

import (
	"strings"
)

type CodeNameInterface interface {
	OptionCodeNames() []string
}

type CodeName struct{}

func (cn CodeName) Get(cni CodeNameInterface) []map[string]interface{} {
	var results []map[string]interface{}
	for _, code := range cni.OptionCodeNames() {
		results = append(results, cn.CodeAndName(code))
	}

	return results
}

func (cn CodeName) Display(code string) string {
	display := strings.Replace(code, "_", " ", -1)
	display = strings.Replace(display, "-", " ", -1)

	return strings.Title(display)
}

func (cn CodeName) CodeAndName(code string) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"display": cn.Display(code),
	}
}
