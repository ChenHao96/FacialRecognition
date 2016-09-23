/**
描述

创建一个Person

一个Person最多允许包含10000个Face。
开发版最多允许创建100个person。
*/
package person

import (
	"encoding/json"
	"net/url"
)

const personCreateApi_url = API_URL + "/person/create"

type PersonCreateRequestParam struct {
	PERSON_NAME string   //Person的Name信息，必须在App中全局唯一。Name不能包含^@,&=*'"等非法字符，且长度不得超过255。Name也可以不指定，此时系统将产生一个随机的name。
	FACE_ID     []string //一组用逗号分隔的face_id, 表示将这些Face加入到该Person中
	TAG         string   //	Person相关的tag，不需要全局唯一，不能包含^@,&=*'"等非法字符，长度不能超过255。
	GROUP_ID    []string //一组用逗号分割的group id列表。如果该参数被指定，该Person被create之后就会被加入到这些组中。
	GROUP_NAME  []string //一组用逗号分割的或者group name列表。如果该参数被指定，该Person被create之后就会被加入到这些组中。
}

type PersonCreateResponseValue struct {
	ADDED_GROUP int    `json:"added_group"`   //成功被加入的group数量
	ADDED_FACE  int    `json:"added_face"`    //成功加入的face数量
	TAG         string `json:"tag,omitempty"` //person相关的tag
	PERSON_NAME string `json:"person_name"`   //相应person的name
	PERSON_ID   string `json:"person_id"`     //相应person的id
}

func PersonCreateFaceImg(param PersonCreateRequestParam) (responseValue PersonCreateResponseValue, err error) {

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)
	if "" != param.PERSON_NAME {
		reqParam.Set("person_name", param.PERSON_NAME)
	}
	if "" != param.TAG {
		reqParam.Set("tag", param.TAG)
	}
	var key_face_id string
	if len(param.FACE_ID) > 0 {
		for index, faceId := range param.FACE_ID {

			if index != 0 && "" != key_face_id {
				key_face_id += ","
			}

			if "" != faceId {
				key_face_id += faceId
			}
		}
		if "" != key_face_id {
			reqParam.Set("face_id", key_face_id)
		}
	}
	var key_group_id string
	if len(param.GROUP_ID) > 0 {
		for index, groupId := range param.GROUP_ID {

			if index != 0 && "" != key_group_id {
				key_group_id += ","
			}

			if "" != groupId {
				key_group_id += groupId
			}
		}
		if "" != key_group_id {
			reqParam.Set("group_id", key_group_id)
		}
	}
	var key_group_Name string
	if len(param.GROUP_NAME) > 0 {
		for index, groupName := range param.GROUP_NAME {

			if index != 0 && "" != key_group_Name {
				key_group_Name += ","
			}

			if "" != groupName {
				key_group_Name += groupName
			}
		}
		if "" != key_group_Name {
			reqParam.Set("group_name", key_group_Name)
		}
	}

	apiUtl := personCreateApi_url + "?" + reqParam.Encode()
	body, err := GetRequest(apiUtl)
	if nil != err {
		return
	}

	err = json.Unmarshal(body, &responseValue)
	return
}