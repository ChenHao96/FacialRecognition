/**
描述

计算两个Face的相似性以及五官相似度
*/
package recognition

import (
	. "cn/com/faceplusplus/public"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const compareApi_url = API_URL + "/recognition/compare"

type CompareRequestParam struct {
	FACE_ID1 string //第一个Face的face_id
	FACE_ID2 string //第二个Face的face_id
	ASYNC    bool   //如果置为true，该API将会以异步方式被调用；也就是立即返回一个session id，稍后可通过/info/get_session查询结果。默认值为false
}

type CompareResponseValue struct {
	SESSION_ID           string  `json:"session_id"`
	SIMILARITY           float64 `json:"similarity"`
	COMPONENT_SIMILARITY ResponseValue_Component_Similarity `json:"component_similarity"`
}

type ResponseValue_Component_Similarity struct {
	EYE     float64 `json:"eye"`
	MOUTH   float64 `json:"mouth"`
	NOSE    float64 `json:"nose"`
	EYEBROW float64 `json:"eyebrow"`
}

func CompareFaceImg(param CompareRequestParam) CompareResponseValue {

	var responseValue CompareResponseValue

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)
	reqParam.Set("async", strconv.FormatBool(param.ASYNC))
	if "" != param.FACE_ID1 {
		reqParam.Set("face_id1", param.FACE_ID1)
	}
	if "" != param.FACE_ID2 {
		reqParam.Set("face_id2", param.FACE_ID2)
	}

	response, err := http.Get(compareApi_url)
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