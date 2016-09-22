package faceplusplus

const API_URL = "http://apicn.faceplusplus.com"

const API_KEY = "58e0813e6a9458268a47bd360c694b43"
const API_SECRET = "0vBdXZRFNDSyCWdnTAJlPnDLOcbho9sD"

type Coordinate struct {
	X float64 `json:"x,omitempty"` //横向坐标
	Y float64 `json:"y,omitempty"` //纵向坐标
}

type TrainResponseValue struct {
	SESSION_ID string `json:"session_id"` //相应请求的session标识符，可用于结果查询
}

type ResponseValue_Face_Position struct {
	CENTER      Coordinate `json:"center"`      //检出的人脸框的中心点坐标, x & y 坐标分别表示在图片中的宽度和高度的百分比 (0~100之间的实数)
	EYE_LEFT    Coordinate `json:"eye_left"`    //相应人脸的左眼坐标，x & y 坐标分别表示在图片中的宽度和高度的百分比 (0~100之间的实数)
	EYE_RIGHT   Coordinate `json:"eye_right"`   //相应人脸的右眼坐标，x & y 坐标分别表示在图片中的宽度和高度的百分比 (0~100之间的实数)
	HEIGHT      float32    `json:"height"`      //0~100之间的实数，表示检出的脸的高度在图片中百分比
	MOUTH_LEFT  Coordinate `json:"mouth_left"`  //相应人脸的左侧嘴角坐标，x & y 坐标分别表示在图片中的宽度和高度的百分比 (0~100之间的实数)
	MOUTH_RIGHT Coordinate `json:"mouth_right"` //相应人脸的右侧嘴角坐标，x & y 坐标分别表示在图片中的宽度和高度的百分比 (0~100之间的实数
	NOSE        Coordinate `json:"nose"`        //相应人脸的鼻尖坐标，x & y 坐标分别表示在图片中的宽度和高度的百分比 (0~100之间的实数)
	WIDTH       float32    `json:"width"`       //0~100之间的实数，表示检出的脸的宽度在图片中百分比
}

type ResponseValue_Face_Attribute_Confidence struct {
	CONFIDENCE  float64 `json:"confidence,omitempty"` //confidence表示置信度
	VALUE       string  `json:"value,omitempty"`      //value的值为Male/Female
	PERSON_ID   string  `json:"person_id,omitempty"`
	PERSON_NAME string  `json:"person_name,omitempty"`
	TAG         string  `json:"tag,omitempty"`
}
