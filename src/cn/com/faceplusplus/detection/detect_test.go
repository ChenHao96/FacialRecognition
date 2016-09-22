package detection

import (
	"testing"
	"encoding/json"
	"fmt"
)

func Test_DetectFaceImg(t *testing.T) {

	var param DetectRequestParam
	param.UPLOAD.URL = "http://b.hiphotos.baidu.com/baike/c0%3Dbaike80%2C5%2C5%2C80%2C26/sign=9db9758da6c27d1eb12b33967abcc60b/21a4462309f79052d1a480170ef3d7ca7bcbd564.jpg"

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