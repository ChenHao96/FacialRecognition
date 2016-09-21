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