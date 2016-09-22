/**
描述

检测给定图片(Image)中的所有人脸(Face)的位置和相应的面部属性

目前面部属性包括性别(gender), 年龄(age), 种族(race), 微笑程度(smiling), 眼镜(glass)和姿势(pose)
若结果的face_id没有被加入任何faceSet/person之中，则在72小时之后过期被自动清除.
*/
package detection

import (
	"bytes"
	. "cn/com/faceplusplus/public"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const detectApi_url = API_URL + "/detection/detect"

type DetectRequestParam struct {
	URL       string //待检测图片的URL与img二选一
	IMG       string //通过POST方法上传的二进制数据，原始图片大小需要小于1M与url二选一
	MODE      string //检测模式可以是normal(默认) 或者 oneFace 。在oneFace模式中，检测器仅找出图片中最大的一张脸
	ATTRIBUTE string //可以是none或者由逗号分割的属性列表。默认为gender, age, race, smiling。目前支持的属性包括：gender, age, race, smiling, glass, pose
	TAG       string //可以为图片中检测出的每一张Face指定一个不包含^@,&=*'"等非法字符且不超过255字节的字符串作为tag，tag信息可以通过 /info/get_face 查询
	ASYNC     bool   //如果置为true，该API将会以异步方式被调用；也就是立即返回一个session id，稍后可通过/info/get_session查询结果。默认值为false
}

type ResponseValue_Face struct {
	ATTRIBUTE ResponseValue_Face_Attribute `json:"attribute"`     //人脸属性
	FACE_ID   string                       `json:"face_id"`       //被检测出的每一张人脸都在Face++系统中的标识符
	POSITION  ResponseValue_Face_Position  `json:"position"`      //面部属性坐标
	TAG       string                       `json:"tag,omitempty"` //请求时传递的参数
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

type DetectResponseValue struct {
	FACE       []ResponseValue_Face `json:"face"`          //被检测出的人脸的列表
	IMG_HEIGHT int                  `json:"img_height"`    //请求图片的高度
	IMG_ID     string               `json:"img_id"`        //Face++系统中的图片标识符，用于标识用户请求中的图片
	IMG_WIDTH  int                  `json:"img_width"`     //请求图片的宽度
	SESSION_ID string               `json:"session_id"`    //相应请求的session标识符，可用于结果查询
	URL        string               `json:"url,omitempty"` //请求中图片的url
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

	var body []byte
	apiUrl := detectApi_url + "?" + reqParam.Encode()
	if "" == param.URL && param.IMG != "" {

		body, err = upload("img", param.IMG, apiUrl)
	} else if "" == param.IMG && param.URL != "" {

		apiUrl += "&url=" + param.URL
		response, err := http.Get(apiUrl)
		defer response.Body.Close()
		if nil == err {
			body, err = ioutil.ReadAll(response.Body)
		}
	}

	if nil != err {
		return
	}

	err = json.Unmarshal(body, &responseValue)
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
