package recognition

import (
	"fmt"
	"encoding/json"
	"testing"
)

func Test_SearchFaceImg(t *testing.T) {

	var param SearchRequestParam
	param.FaceSet_ID = ""
	param.KEY_FACE_ID = []string{}

	res, err := SearchFaceImg(param)
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}