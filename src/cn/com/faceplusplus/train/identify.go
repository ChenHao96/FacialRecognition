/**
描述

针对identify功能对一个Group进行训练。请注意:

1.在一个Group内进行identify之前，必须先对该Group进行Train
2.当一个Group内的数据被修改后(例如增删Person, 增删Person相关的Face等)，为使这些修改生效，Group应当被重新Train
3.Train所花费的时间较长, 因此该调用是异步的，仅返回session_id。
4.Train时需要保证group内的所有person均非空。
5.训练的结果可以通过/info/get_session查询。当训练完成时，返回值中将包含{"success": true}
*/
package train

import (
	. "cn/com/faceplusplus/public"
	"encoding/json"
	"net/http"
	"net/url"
	"io/ioutil"
)

const identifyApi_url = API_URL + "/train/identify"

type IdentifyRequestParam struct {
	GROUP_ID   string //识别候选人组成的GroupId
	GROUP_NAME string //识别候选人组成的GroupName
}

func IdentifyFacesImg(param IdentifyRequestParam) string {

	var responseValue TrainResponseValue

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)
	if "" != param.GROUP_ID {
		reqParam.Set("group_id", param.GROUP_ID)
	}
	if "" != param.GROUP_NAME {
		reqParam.Set("group_name", param.GROUP_NAME)
	}

	response, err := http.Get(identifyApi_url)
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

	return responseValue.SESSION_ID
}