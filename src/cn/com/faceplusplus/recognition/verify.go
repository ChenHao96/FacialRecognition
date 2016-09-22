/**
描述

给定一个Face和一个Person，返回是否是同一个人的判断以及置信度。

注意，当Person中的信息被修改之后（增加，删除了Face等），为了保证结果与最新数据一致，Person应当被重新train.
见/train/verify 。否则调用此API时将使用最后一次train时的数据。
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

const verifyApi_url = API_URL + "/recognition/verify"

type VerifyRequestParam struct {
	FACE_ID     string //待verify的face_id
	PERSON_ID   string //对应的PersonId
	PERSON_NAME string //对应的PersonName
	ASYNC       bool   //如果置为true，该API将会以异步方式被调用；也就是立即返回一个session id，稍后可通过/info/get_session查询结果。默认值为false
}

type VerifyResponseValue struct {
	IS_SAME_PERSON bool    `json:"is_same_person"` //两个输入是否为同一人的判断
	CONFIDENCE     float64 `json:"confidence"`     //系统对这个判断的置信度
	SESSION_ID     string  `json:"session_id"`     //相应请求的session标识符，可用于结果查询
}

func VerifyFaceImg(param VerifyRequestParam) (responseValue VerifyResponseValue, err error) {

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)
	reqParam.Set("async", strconv.FormatBool(param.ASYNC))
	if "" != param.FACE_ID {
		reqParam.Set("face_id", param.FACE_ID)
	}
	if "" != param.PERSON_ID {
		reqParam.Set("person_id", param.PERSON_ID)
	}
	if "" != param.PERSON_NAME {
		reqParam.Set("person_name", param.PERSON_NAME)
	}

	apiUtl := verifyApi_url + "?" + reqParam.Encode()
	response, err := http.Get(apiUtl)
	defer response.Body.Close()
	if nil != err {
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		return
	}

	err = json.Unmarshal(body, &responseValue)
	return
}
