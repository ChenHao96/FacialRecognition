/**
描述

返回该App中的所有Faceset。
*/
package info

import (
	"encoding/json"
	"net/url"
)

const faceSetListApi_url = API_URL + "/info/get_faceset_list"

func FaceSetListFaceImg() (responseValue FaceSetResponseValue, err error) {

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)

	apiUtl := faceSetListApi_url + "?" + reqParam.Encode()
	body, err := GetRequest(apiUtl)
	if nil != err {
		return
	}

	err = json.Unmarshal(body, &responseValue)
	return
}