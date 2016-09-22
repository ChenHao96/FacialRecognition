package recognition

import (
	"fmt"
	"encoding/json"
	"testing"
)

func Test_CompareFaceImg(t *testing.T) {

	var param CompareRequestParam
	param.FACE_ID1 = "0814532b8ca9afd3341dce872750792e"
	param.FACE_ID2 = "2fb9fc3cc18c001fb01c13242218cb81"

	res, err := CompareFaceImg(param)
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}