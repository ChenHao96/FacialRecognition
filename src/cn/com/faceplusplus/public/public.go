package faceplusplus

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

const API_URL = "http://apicn.faceplusplus.com"

const API_KEY = "58e0813e6a9458268a47bd360c694b43"
const API_SECRET = "0vBdXZRFNDSyCWdnTAJlPnDLOcbho9sD"

type Coordinate struct {
	X float64 `json:"x,omitempty"` //横向坐标
	Y float64 `json:"y,omitempty"` //纵向坐标
}

type UploadRequestParam struct {
	URL string //待检测图片的URL与img二选一
	IMG string //通过POST方法上传的二进制数据，原始图片大小需要小于1M与url二选一
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

type SearchAndGroupRequestParam struct {
	FaceSet_ID   string //用于搜索的face组成的faceSetId
	FaceSet_NAME string //用于搜索的face组成的faceSetName
}

type ResponseValue_Candidate struct {
	SIMILARITY float64 `json:"similarity,omitempty"`
	ResponseValue_Faces
}

type ResponseValue_Face_Attribute struct {
	AGE     Attribute_Other                          `json:"age"`             //包含年龄分析结果，value的值为一个非负整数表示估计的年龄, range表示估计年龄的正负区间
	GENDER  ResponseValue_Face_Attribute_Confidence  `json:"gender"`          //包含性别分析结果，value的值为Male/Female, confidence表示置信度
	GLASS   *ResponseValue_Face_Attribute_Confidence `json:"glass,omitempty"` //包含眼镜佩戴分析结果，value的值为None/Dark/Normal, confidence表示置信度
	POSE    *ResponseValue_Face_Attribute_Pose       `json:"pose,omitempty"`  //包含脸部姿势分析结果，包括pitch_angle, roll_angle, yaw_angle，分别对应抬头，旋转（平面旋转），摇头。单位为角度。
	RACE    ResponseValue_Face_Attribute_Confidence  `json:"race"`            //包含人种分析结果，value的值为Asian/White/Black, confidence表示置信度
	SMILING Attribute_Other                          `json:"smiling"`         //包含微笑程度分析结果，value的值为0－100的实数，越大表示微笑程度越高
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

type ResponseValue_Faces struct {
	ATTRIBUTE *ResponseValue_Face_Attribute `json:"attribute,omitempty"` //人脸属性
	FACE_ID   string                       `json:"face_id"`              //被检测出的每一张人脸都在Face++系统中的标识符
	POSITION  *ResponseValue_Face_Position  `json:"position,omitempty"`  //面部属性坐标
	TAG       string                       `json:"tag,omitempty"`        //请求时传递的参数
}

type DetectResponseValue struct {
	FACE       []*ResponseValue_Faces `json:"face,omitempty"`      //被检测出的人脸的列表
	IMG_HEIGHT int                   `json:"img_height,omitempty"` //请求图片的高度
	IMG_ID     string               `json:"img_id"`                //Face++系统中的图片标识符，用于标识用户请求中的图片
	IMG_WIDTH  int                  `json:"img_width,omitempty"`   //请求图片的宽度
	SESSION_ID string               `json:"session_id,omitempty"`  //相应请求的session标识符，可用于结果查询
	URL        string               `json:"url,omitempty"`         //请求中图片的url
}

type FaceRequestParam struct {
	FACE_ID string //待verify的face_id
}

type FaceResponseValue_Info_Person struct {
	PERSON_ID   string `json:"person_id,omitempty"`
	PERSON_NAME string `json:"person_name,omitempty"`
	TAG         string `json:"tag,omitempty"`
}

type FaceResponseValue_Info_FaceSet struct {
	FaceSet_ID   string `json:"faceset_id,omitempty"`
	FaceSet_NAME string `json:"faceset_name,omitempty"`
	TAG          string `json:"tag,omitempty"`
}

type FaceResponseValue_Info_Group struct {
	GROUP_ID   string `json:"group_id,omitempty"`
	GROUP_NAME string `json:"group_name,omitempty"`
	TAG        string `json:"tag,omitempty"`
}

type PersonResponseValue struct {
	PERSON []*FaceResponseValue_Info_Person  `json:"person,omitempty"`
}

type FaceSetResponseValue struct {
	FaceSet []*FaceResponseValue_Info_FaceSet  `json:"faceset,omitempty"`
}

type GroupResponseValue struct {
	GROUP []*FaceResponseValue_Info_Group  `json:"group,omitempty"`
}

type GroupingPublicResponseValue struct {
	CREATE_TIME int                   `json:"create_time,omitempty"` //任务开始时间，单位：秒
	FINISH_TIME int                   `json:"finish_time,omitempty"` //任务结束时间，单位：秒
	SESSION_ID  string                `json:"session_id"`            //相应请求的session标识符，可用于结果查询
	STATUS      string                `json:"status,omitempty"`      //在/info/get_session中，相应session的与状态。取值为 SUCC(已完成) / FAILED(失败) / INQUEUE(队列中)
}

func GetRequest(apiUrl string) (body []byte, err error) {

	response, err := http.Get(apiUrl)
	defer response.Body.Close()
	if nil == err {
		body, err = ioutil.ReadAll(response.Body)
	}

	fmt.Println(string(body))

	return
}

func Upload(apiUrl string, param UploadRequestParam) (body []byte, err error) {

	if "" == param.URL && param.IMG != "" {

		body, err = upload("img", param.IMG, apiUrl)
		fmt.Println(string(body))
	} else if "" != param.URL && param.IMG != "" {

		apiUrl += "&url=" + param.URL
		body, err = GetRequest(apiUrl)
	}

	return
}

func upload(requestKey, fileName, url string) (body []byte, err error) {

	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	fw, err := w.CreateFormFile(requestKey, fileName)
	if err != nil {
		return
	}

	fd, err := os.Open(fileName)
	defer fd.Close()
	if err != nil {
		return
	}

	if _, err = io.Copy(fw, fd); err != nil {
		return
	}
	w.Close()

	request, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", w.FormDataContentType())

	response, err := http.DefaultClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(response.Body)
	return
}
