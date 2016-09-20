package detection

import (
	. "../"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

const api_url = API_URL + "/detection/detect"

type RequestParam struct {
	API_KEY    string //App的Face++ API Key
	API_SECRET string //APP的Face++ API Secret
	URL        string //待检测图片的URL与img二选一
	IMG        []byte //通过POST方法上传的二进制数据，原始图片大小需要小于1M与url二选一
	MODE       string //检测模式可以是normal(默认) 或者 oneface 。在oneface模式中，检测器仅找出图片中最大的一张脸
	ATTRIBUTE  string //可以是none或者由逗号分割的属性列表。默认为gender, age, race, smiling。目前支持的属性包括：gender, age, race, smiling, glass, pose
	TAG        string //可以为图片中检测出的每一张Face指定一个不包含^@,&=*'"等非法字符且不超过255字节的字符串作为tag，tag信息可以通过 /info/get_face 查询
	ASYNC      bool   //如果置为true，该API将会以异步方式被调用；也就是立即返回一个session id，稍后可通过/info/get_session查询结果。默认值为false
}

type ResponseValue struct {
	FACE       []ResponseValue_Face `json:"face"`       //被检测出的人脸的列表
	IMG_HEIGHT int                  `json:"img_height"` //请求图片的高度
	IMG_ID     string               `json:"img_id"`     //Face++系统中的图片标识符，用于标识用户请求中的图片
	IMG_WIDTH  int                  `json:"img_width"`  //请求图片的宽度
	SESSION_ID string               `json:"session_id"` //相应请求的session标识符，可用于结果查询
	URL        string               `json:"url"`        //请求中图片的url
}

func DetectFaceImg(param RequestParam) ResponseValue {

	var response ResponseValue
	if (nil != param) {

		//TODO:未完成
		response, err := http.Post("", "", bytes.NewReader(param.IMG))
		defer response.Body.Close()
		if err != nil {
			panic(err.Error())
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err.Error())
		}

		err = json.Unmarshal(body, &response)
		if nil != err {
			panic(err.Error())
		}
	}

	return response
}