/**
描述

对于一个待查询的Face列表（或者对于给定的Image中所有的Face），在一个Group中查询最相似的Person。

注意，当Group中的信息被修改之后（增加了Person, Face等），为了保证结果与最新数据一致，Group应当被重新train.
见/train/identify 。否则调用此API时将使用最后一次train时的数据。
*/
package recognition

import (
	. "cn/com/faceplusplus/public"
	"encoding/json"
	"net/url"
	"strconv"
)

const identifyApi_url = API_URL + "/recognition/identify"

type IdentifyRequestParam struct {
	UPLOAD      UploadRequestParam
	GROUP_ID    string   //识别候选人组成的GroupId
	GROUP_NAME  string   //识别候选人组成的GroupName
	MODE        string   //检测模式可以是normal(默认) 或者 oneFace 。在oneFace模式中，检测器仅找出图片中最大的一张脸
	ASYNC       bool     //如果置为true，该API将会以异步方式被调用；也就是立即返回一个session id，稍后可通过/info/get_session查询结果。默认值为false
	KEY_FACE_ID []string //开发者也可以指定一个face_id的列表来表明对这些face进行识别。可以设置此参数key_face_id为一个逗号隔开的face_id列表。
}

type IdentifyResponseValue struct {
	FACE       *[]IdentifyResponseValue_Faces `json:"face,omitempty"` //人脸的列表
	SESSION_ID string                         `json:"session_id"`     //相应请求的session标识符，可用于结果查询
}

type IdentifyResponseValue_Faces struct {
	CANDIDATE *[]ResponseValue_Face_Attribute_Confidence `json:"candidate,omitempty"` //识别结果。candidate包含不超过3个人，包含相应person信息与相应的置信度
	FACE_ID   string                                     `json:"face_id"`             //被检测出的每一张人脸都在Face++系统中的标识符
	POSITION  *ResponseValue_Face_Position               `json:"p·osition,omitempty"` //面部属性坐标
}

func IdentifyFacesImg(param IdentifyRequestParam) (responseValue IdentifyResponseValue, err error) {

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)
	reqParam.Set("async", strconv.FormatBool(param.ASYNC))
	if "" != param.GROUP_ID {
		reqParam.Set("group_id", param.GROUP_ID)
	}
	if "" != param.GROUP_NAME {
		reqParam.Set("group_name", param.GROUP_NAME)
	}
	if "" != param.MODE {
		reqParam.Set("mode", param.MODE)
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

	apiUrl := identifyApi_url + "?" + reqParam.Encode()
	body, err := Upload(apiUrl, param.UPLOAD)
	if nil != err {
		return
	}

	err = json.Unmarshal(body, &responseValue)
	return
}