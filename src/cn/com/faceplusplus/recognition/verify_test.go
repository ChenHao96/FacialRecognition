package recognition

import (
	"fmt"
	"encoding/json"
	"testing"
)

func Test_VerifyFaceImg(t *testing.T) {

	var param VerifyRequestParam
	param.FACE_ID = ""
	param.PERSON_ID = ""

	res, err := VerifyFaceImg(param)
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}