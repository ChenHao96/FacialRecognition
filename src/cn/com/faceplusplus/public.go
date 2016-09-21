package faceplusplus

const API_URL = "http://apicn.faceplusplus.com"

const API_KEY = "58e0813e6a9458268a47bd360c694b43"
const API_SECRET = "0vBdXZRFNDSyCWdnTAJlPnDLOcbho9sD"

type ResponseValue_Face struct {
	ATTRIBUTE ResponseValue_Face_Attribute `json:"attribute"`     //人脸属性
	FACE_ID   string                       `json:"face_id"`       //被检测出的每一张人脸都在Face++系统中的标识符
	POSITION  ResponseValue_Face_Position  `json:"position"`      //面部属性坐标
	TAG       string                       `json:"tag,omitempty"` //请求时传递的参数
}

type ResponseValue_Face_Attribute struct {
	AGE     Attribute_Other                         `json:"age"`             //包含年龄分析结果，value的值为一个非负整数表示估计的年龄, range表示估计年龄的正负区间
	GENDER  ResponseValue_Face_Attribute_Confidence `json:"gender"`          //包含性别分析结果，value的值为Male/Female, confidence表示置信度
	GLASS   ResponseValue_Face_Attribute_Confidence `json:"glass,omitempty"` //包含眼镜佩戴分析结果，value的值为None/Dark/Normal, confidence表示置信度
	POSE    ResponseValue_Face_Attribute_Pose       `json:"pose,omitempty"`  //包含脸部姿势分析结果，包括pitch_angle, roll_angle, yaw_angle，分别对应抬头，旋转（平面旋转），摇头。单位为角度。
	RACE    ResponseValue_Face_Attribute_Confidence `json:"race"`            //包含人种分析结果，value的值为Asian/White/Black, confidence表示置信度
	SMILING Attribute_Other                         `json:"smiling"`         //包含微笑程度分析结果，value的值为0－100的实数，越大表示微笑程度越高
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

type Coordinate struct {
	X float64 `json:"x"` //横向坐标
	Y float64 `json:"y"` //纵向坐标
}

type ResponseValue_Face_Attribute_Pose struct {
	PITCH_ANGLE Attribute_Other `json:"pitch_angle,omitempty"` //抬头
	ROLL_ANGLE  Attribute_Other `json:"roll_angle,omitempty"`  //旋转
	YAW_ANGLE   Attribute_Other `json:"yaw_angle,omitempty"`   //摇头
}

type Attribute_Other struct {
	RANGE int     `json:"range,omitempty"` //用于年龄 range表示估计年龄的正负区间
	VALUE float64 `json:"value,omitempty"`           //值
}

type ResponseValue_Face_Attribute_Confidence struct {
	CONFIDENCE float32 `json:"confidence,omitempty"` //confidence表示置信度
	VALUE      string  `json:"value,omitempty"`      //value的值为Male/Female
}