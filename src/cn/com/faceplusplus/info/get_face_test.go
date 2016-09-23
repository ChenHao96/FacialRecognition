package info

import (
	. "cn/com/faceplusplus/public"
	"testing"
	"encoding/json"
	"fmt"
)

func Test_FaceFaceImg(t *testing.T) {

	var param FaceRequestParam
	param.FACE_ID = "909eab4924aab2b6dbf417d1dee57e61"

	res, err := FaceFaceImg(param)
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}