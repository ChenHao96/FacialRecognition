/**
描述

给定一个Face和一个FaceSet，在该FaceSet内搜索最相似的Face。
提示：若搜索集合需要包含超过10000张人脸，可以分成多个FaceSet分别调用search功能再将结果按confidence顺序合并即可。

注意，当FaceSet中的信息被修改之后（增加，删除了Face等），为了保证结果与最新数据一致，FaceSet应当被重新train.
见/train/search。否则调用此API时将使用最后一次train时的数据。
*/
package recognition

import (
	. "cn/com/faceplusplus/public"
	"encoding/json"
	"strconv"
	"net/url"
)

const searchApi_url = API_URL + "/recognition/search"

type SearchRequestParam struct {
	FaceSet_ID   string   //指定搜索范围为此FaceSetId
	FaceSet_NAME string   //指定搜索范围为此FaceSetName
	COUNT        int      //表示一共获取不超过count个搜索结果。默认count=3
	ASYNC        bool     //如果置为true，该API将会以异步方式被调用；也就是立即返回一个session id，稍后可通过/info/get_session查询结果。默认值为false
	KEY_FACE_ID  []string //开发者也可以指定一个face_id的列表来表明对这些face进行识别。可以设置此参数key_face_id为一个逗号隔开的face_id列表。
}

type SearchResponseValue struct {
	CANDIDATE  *[]ResponseValue_Candidate `json:"candidate,omitempty"` //搜索结果，包含相应face信息与相应的置信度
	SESSION_ID string                     `json:"session_id"`          //相应请求的session标识符，可用于结果查询
}

type ResponseValue_Candidate struct {
	FACE_ID    string  `json:"face_id,omitempty"`
	SIMILARITY float64 `json:"similarity,omitempty"`
	TAG        string  `json:"tag,omitempty"`
}

func SearchFaceImg(param SearchRequestParam) (responseValue SearchResponseValue, err error) {

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)
	reqParam.Set("async", strconv.FormatBool(param.ASYNC))
	if 0 < param.COUNT {
		reqParam.Set("count", strconv.Itoa(param.COUNT))
	}
	if "" != param.FaceSet_ID {
		reqParam.Set("faceset_id", param.FaceSet_ID)
	}
	if "" != param.FaceSet_NAME {
		reqParam.Set("faceset_name", param.FaceSet_NAME)
	}
	var key_face_id string
	if len(param.KEY_FACE_ID) > 0 {
		for index, faceId := range param.KEY_FACE_ID {

			if index != 0 && "" != key_face_id {
				key_face_id += ","
			}

			if "" != faceId {
				key_face_id += faceId
			}
		}
		if "" != key_face_id {
			reqParam.Set("key_face_id", key_face_id)
		}
	}

	apiUtl := searchApi_url + "?" + reqParam.Encode()
	body, err := GetRequest(apiUtl)
	if nil != err {
		return
	}

	err = json.Unmarshal(body, &responseValue)
	return
}