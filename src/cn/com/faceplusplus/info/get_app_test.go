package info

import (
	"fmt"
	"encoding/json"
	"testing"
)

func Test_InfoApp(t *testing.T) {

	res, err := InfoApp()
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}