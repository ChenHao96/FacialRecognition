package detection

import (
	. "cn/com/faceplusplus"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"net/url"
)

const landmarkApi_url = API_URL + "/detection/landmark"

type LandmarkRequestParam struct {
	FACE_ID string //待检测人脸的face_id
	TYPE    string //表示返回的关键点个数，目前支持83p或25p，默认为83p
}

type LandmarkResponseValue struct {
	SESSION_ID string `json:"session_id"`                    //相应请求的session标识符，可用于结果查询
	RESULT     []LandmarkResponseValue_Result`json:"result"` //识别结果
}

type LandmarkResponseValue_Result struct {
	FACE_ID  string  `json:"face_id"`                      //人脸在Face++系统中的标识符
	LANDMARK ResponseValue_Face_Landmark `json:"landmark"` //包含详细关键点分析结果，包含多个关键点的坐标。
}

func LandmarkFaceImg(param LandmarkRequestParam) LandmarkResponseValue {

	var responseValue LandmarkResponseValue

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)
	reqParam.Set("face_id", param.FACE_ID)
	reqType := param.TYPE
	if "25p" != reqType && "83p" != reqType {
		reqType = "83p"
	}
	reqParam.Set("type", reqType)

	apiUrl := landmarkApi_url + "?" + reqParam.Encode()

	response, err := http.Get(apiUrl)
	defer response.Body.Close()
	if nil != err {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &responseValue)
	if nil != err {
		panic(err.Error())
	}

	return responseValue
}