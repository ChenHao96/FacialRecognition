/**
描述

给定一组Face，返回相应的信息(包括源图片, 相关的person等等)。
*/
package info

import (
	"net/url"
	"encoding/json"
)

const faceApi_url = API_URL + "/info/get_face"

type FaceResponseValue struct {
	FACE_INFO []*FaceResponseValue_Info `json:"face_info,omitempty"`
}

type FaceResponseValue_Info struct {
	FaceSet []*FaceResponseValue_Info_FaceSet `json:"faceset,omitempty"`
	IMG_ID  string                            `json:"img_id,omitempty"`
	URL     string                            `json:"url,omitempty"`
	PersonResponseValue
	ResponseValue_Faces
}

func FaceFaceImg(param FaceRequestParam) (responseValue FaceResponseValue, err error) {

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)
	if "" != param.FACE_ID {
		reqParam.Set("face_id", param.FACE_ID)
	}

	apiUtl := faceApi_url + "?" + reqParam.Encode()
	body, err := GetRequest(apiUtl)
	if nil != err {
		return
	}

	err = json.Unmarshal(body, &responseValue)
	return
}