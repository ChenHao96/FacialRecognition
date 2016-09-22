package grouping

import (
	. "cn/com/faceplusplus/public"
	"fmt"
	"encoding/json"
	"testing"
)

func Test_GroupingFaceImg(t *testing.T) {

	var param SearchAndGroupRequestParam
	param.FaceSet_ID = ""

	res, err := GroupingFaceImg(param)
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}
