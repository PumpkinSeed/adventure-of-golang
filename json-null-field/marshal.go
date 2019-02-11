package main

import (
	"encoding/json"
	"fmt"

	"github.com/volatiletech/null"
)

type Test struct {
	TestString      null.String `json:"test_string"`
	TestStringValue null.String `json:"tsv,omitempty"`
	TestUin64       null.Uint64
	TestU64         uint64
	TestU32         null.Uint32 `json:"testu32,omitempty"`
}

func main() {
	var t Test
	t.TestString.SetValid("test")
	t.TestUin64.SetValid(1234)
	t.TestU64 = 1333

	res, _ := json.Marshal(t)
	fmt.Println(string(res))

	var t2 Test

	json.Unmarshal(res, &t2)

	fmt.Println(t2)

	fmt.Println(t2.TestStringValue.String)
}
