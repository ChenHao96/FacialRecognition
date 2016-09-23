package info

import (
	"fmt"
	"encoding/json"
	"testing"
)

func Test_SessionFaceImg(t *testing.T) {

	res, err := SessionFaceImg(SessionRequestParam{SESSION_ID:"9e2bfdae6a594057b84b9fde525148de"})
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}