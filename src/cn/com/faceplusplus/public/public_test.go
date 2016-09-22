package faceplusplus

import (
	"testing"
	"os"
)

func Test_GetRequest(t *testing.T) {

	_, err := GetRequest("https://www.baidu.com")
	if nil != err {
		panic(err.Error())
	}
}

func Test_Upload(t *testing.T) {

	var param UploadRequestParam
	param.IMG = os.Getenv("GOPATH") + "/resources/xijingpin.jpg"
	param.URL = "http://b.hiphotos.baidu.com/baike/c0%3Dbaike80%2C5%2C5%2C80%2C26/sign=9db9758da6c27d1eb12b33967abcc60b/21a4462309f79052d1a480170ef3d7ca7bcbd564.jpg"

	Upload(API_URL + "/detection/detect?api_key=" + API_KEY + "&api_secret=" + API_SECRET, param)
}
