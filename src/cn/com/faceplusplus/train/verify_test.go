package train

import (
	"fmt"
	"encoding/json"
	"testing"
)

func Test_VerifyPersonImg(t *testing.T) {

	var param VerifyRequestParam
	param.PERSON_ID = ""
	param.PERSON_NAME = ""

	res, err := VerifyPersonImg(param)
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}
