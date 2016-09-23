/**
描述

检测给定图片(Image)中的所有人脸(Face)的位置和相应的面部属性

目前面部属性包括性别(gender), 年龄(age), 种族(race), 微笑程度(smiling), 眼镜(glass)和姿势(pose)
若结果的face_id没有被加入任何faceSet/person之中，则在72小时之后过期被自动清除.
*/
package detection

import (
	. "cn/com/faceplusplus/public"
	"encoding/json"
	"net/url"
	"strconv"
)

const detectApi_url = API_URL + "/detection/detect"

type DetectRequestParam struct {
	UPLOAD    UploadRequestParam
	MODE      string //检测模式可以是normal(默认) 或者 oneFace 。在oneFace模式中，检测器仅找出图片中最大的一张脸
	ATTRIBUTE string //可以是none或者由逗号分割的属性列表。默认为gender, age, race, smiling。目前支持的属性包括：gender, age, race, smiling, glass, pose
	TAG       string //可以为图片中检测出的每一张Face指定一个不包含^@,&=*'"等非法字符且不超过255字节的字符串作为tag，tag信息可以通过 /info/get_face 查询
	ASYNC     bool   //如果置为true，该API将会以异步方式被调用；也就是立即返回一个session id，稍后可通过/info/get_session查询结果。默认值为false
}

func DetectFaceImg(param DetectRequestParam) (responseValue DetectResponseValue, err error) {

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)
	reqParam.Set("async", strconv.FormatBool(param.ASYNC))
	if "" != param.TAG {
		reqParam.Set("tag", param.TAG)
	}
	if "" != param.ATTRIBUTE {
		reqParam.Set("attribute", param.ATTRIBUTE)
	}
	if "" != param.MODE {
		reqParam.Set("mode", param.MODE)
	}

	apiUrl := detectApi_url + "?" + reqParam.Encode()
	body, err := Upload(apiUrl, param.UPLOAD)
	if nil != err {
		return
	}

	err = json.Unmarshal(body, &responseValue)
	return
}