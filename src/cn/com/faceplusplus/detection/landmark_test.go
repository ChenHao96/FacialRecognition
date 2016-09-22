package detection

import (
	"fmt"
	"encoding/json"
	"testing"
)

func Test_LandmarkFaceImg(t *testing.T) {

	var param LandmarkRequestParam
	param.FACE_ID = "0814532b8ca9afd3341dce872750792e"

	res, err := LandmarkFaceImg(param)
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}
