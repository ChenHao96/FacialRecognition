package info

import (
	"fmt"
	"encoding/json"
	"testing"
)

func Test_ImageFaceImg(t *testing.T) {

	var param ImageRequestParam
	param.IMG_ID = "b6ec860704c54d6bef4ccd4692a1a6b8"

	res, err := ImageFaceImg(param)
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}