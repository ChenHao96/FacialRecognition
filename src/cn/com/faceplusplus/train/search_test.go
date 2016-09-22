package train

import (
	"fmt"
	"encoding/json"
	"testing"
)

func Test_SearchFacesImg(t *testing.T) {

	var param SearchRequestParam
	param.FaceSet_ID = ""
	param.FaceSet_NAME = ""

	res, err := SearchFacesImg(param)
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}