package util

import (
	"encoding/json"
	"fmt"
)

func Print(os ...interface{}) {
	for _, o := range os {
		printJSONIndent(o)
	}
}

func PrintJSON(o interface{}) {
	s, _ := json.Marshal(o)
	fmt.Println(string(s))
}

func printJSONIndent(o interface{}) {
	s, _ := json.MarshalIndent(o, "", "  ")
	fmt.Println(string(s))
}
