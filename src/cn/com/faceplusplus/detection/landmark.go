/**
描述

检测给定人脸(Face)相应的面部轮廓，五官等关键点的位置，包括25点和83点两种模式。
*/
package detection

import (
	. "cn/com/faceplusplus/public"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"net/url"
)

const landmarkApi_url = API_URL + "/detection/landmark"

type LandmarkRequestParam struct {
	FACE_ID string //待检测人脸的face_id
	TYPE    string //表示返回的关键点个数，目前支持83p或25p，默认为83p
}

type LandmarkResponseValue struct {
	SESSION_ID string `json:"session_id"`                    //相应请求的session标识符，可用于结果查询
	RESULT     []LandmarkResponseValue_Result`json:"result"` //识别结果
}

type ResponseValue_Face_Landmark struct {
	CONTOUR_CHIN                      *Coordinate `json:"contour_chin,omitempty"`
	CONTOUR_LEFT1                     *Coordinate `json:"contour_left1,omitempty"`
	CONTOUR_LEFT2                     *Coordinate `json:"contour_left2,omitempty"`
	CONTOUR_LEFT3                     *Coordinate `json:"contour_left3,omitempty"`
	CONTOUR_LEFT4                     *Coordinate `json:"contour_left4,omitempty"`
	CONTOUR_LEFT5                     *Coordinate `json:"contour_left5,omitempty"`
	CONTOUR_LEFT6                     *Coordinate `json:"contour_left6,omitempty"`
	CONTOUR_LEFT7                     *Coordinate `json:"contour_left7,omitempty"`
	CONTOUR_LEFT8                     *Coordinate `json:"contour_left8,omitempty"`
	CONTOUR_LEFT9                     *Coordinate `json:"contour_left9,omitempty"`
	CONTOUR_RIGHT1                    *Coordinate `json:"contour_right1,omitempty"`
	CONTOUR_RIGHT2                    *Coordinate `json:"contour_right2,omitempty"`
	CONTOUR_RIGHT3                    *Coordinate `json:"contour_right3,omitempty"`
	CONTOUR_RIGHT4                    *Coordinate `json:"contour_right4,omitempty"`
	CONTOUR_RIGHT5                    *Coordinate `json:"contour_right5,omitempty"`
	CONTOUR_RIGHT6                    *Coordinate `json:"contour_right6,omitempty"`
	CONTOUR_RIGHT7                    *Coordinate `json:"contour_right7,omitempty"`
	CONTOUR_RIGHT8                    *Coordinate `json:"contour_right8,omitempty"`
	CONTOUR_RIGHT9                    *Coordinate `json:"contour_right9,omitempty"`
	LEFT_EYE_BOTTOM                   Coordinate `json:"left_eye_bottom"`
	LEFT_EYE_CENTER                   Coordinate `json:"left_eye_center"`
	LEFT_EYE_LEFT_CORNER              Coordinate `json:"left_eye_left_corner"`
	LEFT_EYE_LOWER_LEFT_QUARTER       *Coordinate `json:"left_eye_lower_left_quarter,omitempty"`
	LEFT_EYE_PUPIL                    Coordinate `json:"left_eye_pupil"`
	LEFT_EYE_RIGHT_CORNER             Coordinate `json:"left_eye_right_corner"`
	LEFT_EYE_TOP                      Coordinate `json:"left_eye_top"`
	LEFT_EYE_UPPER_LEFT_QUARTER       *Coordinate `json:"left_eye_upper_left_quarter,omitempty"`
	LEFT_EYE_UPPER_RIGHT_QUARTER      *Coordinate `json:"left_eye_upper_right_quarter,omitempty"`
	LEFT_EYEBROW_LEFT_CORNER          Coordinate `json:"left_eyebrow_left_corner"`
	LEFT_EYEBROW_LOWER_LEFT_QUARTER   *Coordinate `json:"left_eyebrow_lower_left_quarter,omitempty"`
	LEFT_EYEBROW_LOWER_MIDDLE         *Coordinate `json:"left_eyebrow_lower_middle,omitempty"`
	LEFT_EYEBROW_LOWER_RIGHT_QUARTER  *Coordinate `json:"left_eyebrow_lower_right_quarter,omitempty"`
	LEFT_EYEBROW_RIGHT_CORNER         Coordinate `json:"left_eyebrow_right_corner"`
	LEFT_EYEBROW_UPPER_LEFT_QUARTER   *Coordinate `json:"left_eyebrow_upper_left_quarter,omitempty"`
	LEFT_EYEBROW_UPPER_MIDDLE         *Coordinate `json:"left_eyebrow_upper_middle,omitempty"`
	LEFT_EYEBROW_UPPER_RIGHT_QUARTER  *Coordinate `json:"left_eyebrow_upper_right_quarter,omitempty"`
	MOUTH_LEFT_CORNER                 Coordinate `json:"mouth_left_corner"`
	MOUTH_LOWER_LIP_BOTTOM            Coordinate `json:"mouth_lower_lip_bottom"`
	MOUTH_LOWER_LIP_LEFT_CONTOUR1     *Coordinate `json:"mouth_lower_lip_left_contour1,omitempty"`
	MOUTH_LOWER_LIP_LEFT_CONTOUR2     *Coordinate `json:"mouth_lower_lip_left_contour2,omitempty"`
	MOUTH_LOWER_LIP_LEFT_CONTOUR3     *Coordinate `json:"mouth_lower_lip_left_contour3,omitempty"`
	MOUTH_LOWER_LIP_RIGHT_CONTOUR1    *Coordinate `json:"mouth_lower_lip_right_contour1,omitempty"`
	MOUTH_LOWER_LIP_RIGHT_CONTOUR2    *Coordinate `json:"mouth_lower_lip_right_contour2,omitempty"`
	MOUTH_LOWER_LIP_RIGHT_CONTOUR3    *Coordinate `json:"mouth_lower_lip_right_contour3,omitempty"`
	MOUTH_LOWER_LIP_TOP               Coordinate `json:"mouth_lower_lip_top"`
	MOUTH_RIGHT_CORNER                Coordinate `json:"mouth_right_corner"`
	MOUTH_UPPER_LIP_BOTTOM            Coordinate `json:"mouth_upper_lip_bottom"`
	MOUTH_UPPER_LIP_LEFT_CONTOUR1     *Coordinate `json:"mouth_upper_lip_left_contour1,omitempty"`
	MOUTH_UPPER_LIP_LEFT_CONTOUR2     *Coordinate `json:"mouth_upper_lip_left_contour2,omitempty"`
	MOUTH_UPPER_LIP_LEFT_CONTOUR3     *Coordinate `json:"mouth_upper_lip_left_contour3,omitempty"`
	MOUTH_UPPER_LIP_RIGHT_CONTOUR1    *Coordinate `json:"mouth_upper_lip_right_contour1,omitempty"`
	MOUTH_UPPER_LIP_RIGHT_CONTOUR2    *Coordinate `json:"mouth_upper_lip_right_contour2,omitempty"`
	MOUTH_UPPER_LIP_RIGHT_CONTOUR3    *Coordinate `json:"mouth_upper_lip_right_contour3,omitempty"`
	MOUTH_UPPER_LIP_TOP               Coordinate `json:"mouth_upper_lip_top"`
	NOSE_CONTOUR_LEFT1                *Coordinate `json:"nose_contour_left1,omitempty"`
	NOSE_CONTOUR_LEFT2                *Coordinate `json:"nose_contour_left2,omitempty"`
	NOSE_CONTOUR_LEFT3                *Coordinate `json:"nose_contour_left3,omitempty"`
	NOSE_CONTOUR_LOWER_MIDDLE         *Coordinate `json:"nose_contour_lower_middle,omitempty"`
	NOSE_CONTOUR_RIGHT1               *Coordinate `json:"nose_contour_right1,omitempty"`
	NOSE_CONTOUR_RIGHT2               *Coordinate `json:"nose_contour_right2,omitempty"`
	NOSE_CONTOUR_RIGHT3               *Coordinate `json:"nose_contour_right3,omitempty"`
	NOSE_LEFT                         Coordinate `json:"nose_left"`
	NOSE_RIGHT                        Coordinate `json:"nose_right"`
	NOSE_TIP                          Coordinate `json:"nose_tip"`
	RIGHT_EYE_BOTTOM                  Coordinate `json:"right_eye_bottom"`
	RIGHT_EYE_CENTER                  Coordinate `json:"right_eye_center"`
	RIGHT_EYE_LEFT_CORNER             Coordinate `json:"right_eye_left_corner"`
	RIGHT_EYE_LOWER_LEFT_QUARTER      *Coordinate `json:"right_eye_lower_left_quarter,omitempty"`
	RIGHT_EYE_LOWER_RIGHT_QUARTER     *Coordinate `json:"right_eye_lower_right_quarter,omitempty"`
	RIGHT_EYE_PUPIL                   Coordinate `json:"right_eye_pupil"`
	RIGHT_EYE_RIGHT_CORNER            Coordinate `json:"right_eye_right_corner"`
	RIGHT_EYE_TOP                     Coordinate `json:"right_eye_top"`
	RIGHT_EYE_UPPER_LEFT_QUARTER      *Coordinate `json:"right_eye_upper_left_quarter,omitempty"`
	RIGHT_EYE_UPPER_RIGHT_QUARTER     *Coordinate `json:"right_eye_upper_right_quarter,omitempty"`
	RIGHT_EYEBROW_LEFT_CORNER         Coordinate `json:"right_eyebrow_left_corner"`
	RIGHT_EYEBROW_LOWER_LEFT_QUARTER  *Coordinate `json:"right_eyebrow_lower_left_quarter,omitempty"`
	RIGHT_EYEBROW_LOWER_MIDDLE        *Coordinate `json:"right_eyebrow_lower_middle,omitempty"`
	RIGHT_EYEBROW_LOWER_RIGHT_QUARTER *Coordinate `json:"right_eyebrow_lower_right_quarter,omitempty"`
	RIGHT_EYEBROW_RIGHT_CORNER        Coordinate `json:"right_eyebrow_right_corner"`
	RIGHT_EYEBROW_UPPER_LEFT_QUARTER  *Coordinate `json:"right_eyebrow_upper_left_quarter,omitempty"`
	RIGHT_EYEBROW_UPPER_MIDDLE        *Coordinate `json:"right_eyebrow_upper_middle,omitempty"`
	RIGHT_EYEBROW_UPPER_RIGHT_QUARTER *Coordinate `json:"right_eyebrow_upper_right_quarter,omitempty"`
}

type LandmarkResponseValue_Result struct {
	FACE_ID  string  `json:"face_id"`                      //人脸在Face++系统中的标识符
	LANDMARK ResponseValue_Face_Landmark `json:"landmark"` //包含详细关键点分析结果，包含多个关键点的坐标。
}

func LandmarkFaceImg(param LandmarkRequestParam) LandmarkResponseValue {

	var responseValue LandmarkResponseValue

	reqParam := url.Values{}
	reqParam.Set("api_key", API_KEY)
	reqParam.Set("api_secret", API_SECRET)
	reqParam.Set("face_id", param.FACE_ID)
	reqType := param.TYPE
	if "25p" != reqType && "83p" != reqType {
		reqType = "83p"
	}
	reqParam.Set("type", reqType)

	apiUrl := landmarkApi_url + "?" + reqParam.Encode()

	response, err := http.Get(apiUrl)
	defer response.Body.Close()
	if nil != err {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &responseValue)
	if nil != err {
		panic(err.Error())
	}

	return responseValue
}