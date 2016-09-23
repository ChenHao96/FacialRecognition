package info

import (
	. "cn/com/faceplusplus/public"
	"encoding/json"
	"net/url"
)

const groupListApi_url = API_URL + "/info/get_group_list"

func GroupListFaceImg() (responseValue GroupResponseValue, err error) {

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)

	apiUtl := groupListApi_url + "?" + reqParam.Encode()
	body, err := GetRequest(apiUtl)
	if nil != err {
		return
	}

	err = json.Unmarshal(body, &responseValue)
	return
}