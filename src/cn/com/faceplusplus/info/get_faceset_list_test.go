package info

import (
	"fmt"
	"encoding/json"
	"testing"
)

func Test_FaceSetListFaceImg(t *testing.T) {

	res, err := FaceSetListFaceImg()
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}