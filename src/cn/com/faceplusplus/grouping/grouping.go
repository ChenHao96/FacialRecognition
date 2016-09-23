/**
描述

给出一个FaceSet，尝试将其分类，使得来自同一个人的Face被放在同一类中。
Grouping所花费的时间较长, 因此该调用是异步的，仅返回session_id。
*/
package grouping

import (
	. "cn/com/faceplusplus/public"
	"net/url"
	"encoding/json"
)

const groupingApi_url = API_URL + "/grouping/grouping"

type GroupingResponseValue struct {
	CREATE_TIME int                   `json:"create_time,omitempty"` //任务开始时间，单位：秒
	FINISH_TIME int                   `json:"finish_time,omitempty"` //任务结束时间，单位：秒
	RESULT      *ResponseValue_Result `json:"x,omitempty"`           //分类结果
	SESSION_ID  string                `json:"session_id"`            //相应请求的session标识符，可用于结果查询
	STATUS      string                `json:"status,omitempty"`      //在/info/get_session中，相应session的与状态。取值为 SUCC(已完成) / FAILED(失败) / INQUEUE(队列中)
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