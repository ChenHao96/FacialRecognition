/**
描述

返回该App中的所有Person
*/
package info

import (
	"encoding/json"
	"net/url"
)

const personListApi_url = API_URL + "/info/get_person_list"

func PersonListFaceImg() (responseValue PersonResponseValue, err error) {

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)

	apiUtl := personListApi_url + "?" + reqParam.Encode()
	body, err := GetRequest(apiUtl)
	if nil != err {
		return
	}

	err = json.Unmarshal(body, &responseValue)
	return
}