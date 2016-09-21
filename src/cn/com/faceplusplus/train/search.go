/**
描述

针对search功能对一个faceSet进行训练。请注意:

1.在一个faceSet内进行search之前，必须先对该faceSet进行Train
2.当一个faceSet内的数据被修改后(例如增删Face等)，为使这些修改生效，faceSet应当被重新Train
3.Train所花费的时间较长, 因此该调用是异步的，仅返回session_id。
4.训练的结果可以通过/info/get_session查询。当训练完成时，返回值中将包含{"success": true}
*/
package train

import (
	. "cn/com/faceplusplus/public"
	"encoding/json"
	"net/http"
	"net/url"
	"io/ioutil"
)

const searchApi_url = API_URL + "/train/search"

type SearchRequestParam struct {
	FaceSet_ID   string //用于搜索的face组成的faceSetId
	FaceSet_NAME string //用于搜索的face组成的faceSetName
}

func SearchFacesImg(param SearchRequestParam) string {

	var responseValue TrainResponseValue

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)
	if "" != param.FaceSet_ID {
		reqParam.Set("faceset_id", param.FaceSet_ID)
	}
	if "" != param.FaceSet_NAME {
		reqParam.Set("faceset_name", param.FaceSet_NAME)
	}

	response, err := http.Get(searchApi_url)
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