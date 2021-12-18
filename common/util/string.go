package util

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(o interface{}) {
	s, _ := json.Marshal(o)
	fmt.Println(string(s))
}

func PrintJSONIndent(o interface{}) {
	s, _ := json.MarshalIndent(o, "", "  ")
	fmt.Println(string(s))
}
