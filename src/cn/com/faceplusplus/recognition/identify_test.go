package recognition

import (
	"fmt"
	"encoding/json"
	"testing"
	"os"
)

func Test_IdentifyFacesImg(t *testing.T) {

	var param IdentifyRequestParam
	param.GROUP_ID = ""
	param.GROUP_NAME = ""
	param.UPLOAD.IMG = os.Getenv("GOPATH") + "/resources/xijingpin.jpg"

	res, err := IdentifyFacesImg(param)
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}
