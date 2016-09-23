package info

import (
	. "cn/com/faceplusplus/public"
	"net/url"
	"encoding/json"
)

const imageApi_url = API_URL + "/info/get_image"

type ImageRequestParam struct {
	IMG_ID string //目标图片的img_id
}

func ImageFaceImg(param ImageRequestParam) (responseValue DetectResponseValue, err error) {

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)
	if "" != param.IMG_ID {
		reqParam.Set("img_id", param.IMG_ID)
	}

	apiUtl := imageApi_url + "?" + reqParam.Encode()
	body, err := GetRequest(apiUtl)
	if nil != err {
		return
	}

	err = json.Unmarshal(body, &responseValue)
	return
}