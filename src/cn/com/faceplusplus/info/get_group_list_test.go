package info

import (
	"fmt"
	"encoding/json"
	"testing"
)

func Test_GroupListFaceImg(t *testing.T) {

	res, err := GroupListFaceImg()
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}