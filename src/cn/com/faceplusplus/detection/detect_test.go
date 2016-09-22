package detection

import (
	"testing"
	"encoding/json"
	"fmt"
	"os"
)

func Test_DetectFaceImg(t *testing.T) {

	var param DetectRequestParam
	param.UPLOAD.IMG = os.Getenv("GOPATH") + "/resources/xijingpin.jpg"

	res, err := DetectFaceImg(param)
	if nil != err {
		panic(err.Error())
	}

	value, err := json.Marshal(res)
	if nil != err {
		panic(err.Error())
	}

	fmt.Println(string(value))
}