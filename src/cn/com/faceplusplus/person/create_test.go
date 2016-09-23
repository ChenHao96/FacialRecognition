package person

import (
	"fmt"
	"encoding/json"
	"testing"
)

func Test_PersonCreateFaceImg(t *testing.T) {

	var param PersonCreateRequestParam
	res, err := PersonCreateFaceImg(param)
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}
