package utils

import (
	"encoding/json"
	"fmt"
)

func Lg(v any) {
	fmt.Println(v)
}

func Lj(v any) {
	result, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(result))
}
