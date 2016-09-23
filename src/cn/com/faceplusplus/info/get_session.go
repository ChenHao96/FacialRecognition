/**
描述

获取session相关状态和结果

可能的status：INQUEUE(队列中), SUCC(成功) 和FAILED(失败)
当status是SUCC时，返回结果中还包含session对应的结果
所有session都将在计算完成72小时之后过期，并被自动清除。
status返回值为SUCC仅表示成功取得运行结果，实际任务成功与否请根据result内容判断
*/
package info

import (
	"encoding/json"
	"net/url"
)

const sessionApi_url = API_URL + "/info/get_session"

type SessionRequestParam struct {
	SESSION_ID string
}

type SessionResponseValue struct {
	SESSION_ID  string                       `json:"session_id"`       //相应请求的session标识符，可用于结果查询
	CREATE_TIME int                          `json:"create_time"`      //任务开始时间，单位：秒
	FINISH_TIME int                          `json:"finish_time"`      //任务结束时间，单位：秒
	STATUS      string                       `json:"status"`           //可能取值有：INQUEUE(队列中), SUCC(成功) 和FAILED(失败)
	RESULT      *SessionResponseValue_Result `json:"result,omitempty"` //返回session_id对应的结果内容
}

type SessionResponseValue_Result struct {
	URL        string                              `json:"url,omitempty"`  //使用图片加载识别的url
	IMG_ID     string                              `json:"img_id"`         //图片id
	IMG_HEIGHT int                                 `json:"img_height"`     //图片高
	IMG_WIDTH  int                                 `json:"img_width"`      //图片宽
	FACE       []*ResponseValue_Faces `json:"face,omitempty"` //识别的属性
}

func SessionFaceImg(param SessionRequestParam) (responseValue SessionResponseValue, err error) {

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)
	reqParam.Set("session_id", param.SESSION_ID)

	apiUtl := sessionApi_url + "?" + reqParam.Encode()
	body, err := GetRequest(apiUtl)
	if nil != err {
		return
	}

	err = json.Unmarshal(body, &responseValue)
	return
}

