/**
描述

针对verify功能对一个person进行训练。请注意:

1.在一个person内进行verify之前，必须先对该person进行Train
2.当一个person内的数据被修改后(例如增删Person相关的Face等)，为使这些修改生效，person应当被重新Train
3.Train所花费的时间较长, 因此该调用是异步的，仅返回session_id。
4.训练的结果可以通过/info/get_session查询。当训练完成时，返回值中将包含{"success": true}
*/
package train

import (
	. "cn/com/faceplusplus/public"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const verifyApi_url = API_URL + "/train/verify"

type VerifyRequestParam struct {
	PERSON_ID   string //验证对象personId
	PERSON_NAME string //验证对象personName
}

func VerifyPersonImg(param VerifyRequestParam) (sessionId string, err error) {

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)
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

	var responseValue TrainResponseValue
	err = json.Unmarshal(body, &responseValue)
	if nil == err {
		sessionId = responseValue.SESSION_ID
	}

	return
}
