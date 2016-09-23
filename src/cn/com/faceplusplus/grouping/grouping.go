/**
描述

给出一个FaceSet，尝试将其分类，使得来自同一个人的Face被放在同一类中。
Grouping所花费的时间较长, 因此该调用是异步的，仅返回session_id。
*/
package grouping

import (
	"net/url"
	"encoding/json"
)

const groupingApi_url = API_URL + "/grouping/grouping"

type GroupingResponseValue struct {
	RESULT *ResponseValue_Result `json:"result,omitempty"` //分类结果
	GroupingPublicResponseValue
}

type ResponseValue_Result struct {
	GROUP     [][]*ResponseValue_Candidate `json:"group,omitempty"`     //分类结果，来自不同人的face被归到不同的类中
	UnGrouped []*ResponseValue_Candidate   `json:"ungrouped,omitempty"` //未归类face对象，无法被归类的face集合
}

func GroupingFaceImg(param SearchAndGroupRequestParam) (responseValue GroupingResponseValue, err error) {

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)
	if "" != param.FaceSet_ID {
		reqParam.Set("faceset_id", param.FaceSet_ID)
	}
	if "" != param.FaceSet_NAME {
		reqParam.Set("faceset_name", param.FaceSet_NAME)
	}

	apiUtl := groupingApi_url + "?" + reqParam.Encode()
	body, err := GetRequest(apiUtl)
	if nil != err {
		return
	}

	err = json.Unmarshal(body, &responseValue)
	return
}