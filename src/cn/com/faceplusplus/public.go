package faceplusplus

const API_URL = "http://apicn.faceplusplus.com"

const API_KEY = "58e0813e6a9458268a47bd360c694b43"
const API_SECRET = "0vBdXZRFNDSyCWdnTAJlPnDLOcbho9sD"

type Coordinate struct {
	X float64 `json:"x,omitempty"` //横向坐标
	Y float64 `json:"y,omitempty"` //纵向坐标
}

/********************************************************************landmark*********************************************************************************************************/
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
/*************************************************************************************************************************************************************************************/

/********************************************************************detect*********************************************************************************************************/
type ResponseValue_Face struct {
	ATTRIBUTE ResponseValue_Face_Attribute `json:"attribute"`     //人脸属性
	FACE_ID   string                       `json:"face_id"`       //被检测出的每一张人脸都在Face++系统中的标识符
	POSITION  ResponseValue_Face_Position  `json:"position"`      //面部属性坐标
	TAG       string                       `json:"tag,omitempty"` //请求时传递的参数
}

type ResponseValue_Face_Attribute struct {
	AGE     Attribute_Other                         `json:"age"`              //包含年龄分析结果，value的值为一个非负整数表示估计的年龄, range表示估计年龄的正负区间
	GENDER  ResponseValue_Face_Attribute_Confidence `json:"gender"`           //包含性别分析结果，value的值为Male/Female, confidence表示置信度
	GLASS   *ResponseValue_Face_Attribute_Confidence `json:"glass,omitempty"` //包含眼镜佩戴分析结果，value的值为None/Dark/Normal, confidence表示置信度
	POSE    *ResponseValue_Face_Attribute_Pose       `json:"pose,omitempty"`  //包含脸部姿势分析结果，包括pitch_angle, roll_angle, yaw_angle，分别对应抬头，旋转（平面旋转），摇头。单位为角度。
	RACE    ResponseValue_Face_Attribute_Confidence `json:"race"`             //包含人种分析结果，value的值为Asian/White/Black, confidence表示置信度
	SMILING Attribute_Other                         `json:"smiling"`          //包含微笑程度分析结果，value的值为0－100的实数，越大表示微笑程度越高
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

type ResponseValue_Face_Attribute_Pose struct {
	PITCH_ANGLE *Attribute_Other `json:"pitch_angle,omitempty"` //抬头
	ROLL_ANGLE  *Attribute_Other `json:"roll_angle,omitempty"`  //旋转
	YAW_ANGLE   *Attribute_Other `json:"yaw_angle,omitempty"`   //摇头
}

type Attribute_Other struct {
	RANGE int     `json:"range,omitempty"` //用于年龄 range表示估计年龄的正负区间
	VALUE float64 `json:"value,omitempty"` //值
}

type ResponseValue_Face_Attribute_Confidence struct {
	CONFIDENCE float32 `json:"confidence,omitempty"` //confidence表示置信度
	VALUE      string  `json:"value,omitempty"`      //value的值为Male/Female
}
/***********************************************************************************************************************************************************************************/