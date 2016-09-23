/**
描述

获取该App相关的信息(该接口在访问时提示弃用)
*/
package info

import (
	"encoding/json"
	"net/url"
)

const AppApi_url = API_URL + "/info/get_app"

type AppResponseValue struct {
	NAME        string `json:"name,omitempty"`
	INFO        string`json:"info,omitempty"`
	DESCRIPTION string `json:"description,omitempty"`
}

func InfoApp() (responseValue AppResponseValue, err error) {

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)

	apiUtl := AppApi_url + "?" + reqParam.Encode()
	body, err := GetRequest(apiUtl)
	if nil != err {
		return
	}

	err = json.Unmarshal(body, &responseValue)
	return
}