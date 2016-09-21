package detection

import (
	. "cn/com/faceplusplus"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"io"
	"os"
	"mime/multipart"
	"net/url"
	"strconv"
)

const detectApi_url = API_URL + "/detection/detect"

type DetectRequestParam struct {
	URL       string //待检测图片的URL与img二选一
	IMG       string //通过POST方法上传的二进制数据，原始图片大小需要小于1M与url二选一
	MODE      string //检测模式可以是normal(默认) 或者 oneface 。在oneface模式中，检测器仅找出图片中最大的一张脸
	ATTRIBUTE string //可以是none或者由逗号分割的属性列表。默认为gender, age, race, smiling。目前支持的属性包括：gender, age, race, smiling, glass, pose
	TAG       string //可以为图片中检测出的每一张Face指定一个不包含^@,&=*'"等非法字符且不超过255字节的字符串作为tag，tag信息可以通过 /info/get_face 查询
	ASYNC     bool   //如果置为true，该API将会以异步方式被调用；也就是立即返回一个session id，稍后可通过/info/get_session查询结果。默认值为false
}

type DetectResponseValue struct {
	FACE       []ResponseValue_Face `json:"face"`          //被检测出的人脸的列表
	IMG_HEIGHT int                  `json:"img_height"`    //请求图片的高度
	IMG_ID     string               `json:"img_id"`        //Face++系统中的图片标识符，用于标识用户请求中的图片
	IMG_WIDTH  int                  `json:"img_width"`     //请求图片的宽度
	SESSION_ID string               `json:"session_id"`    //相应请求的session标识符，可用于结果查询
	URL        string               `json:"url,omitempty"` //请求中图片的url
}

func DetectFaceImg(param DetectRequestParam) DetectResponseValue {

	var responseValue DetectResponseValue
	var body []byte
	var err error

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

	apiUrl := detectApi_url + "?" + reqParam.Encode()
	if "" == param.URL && param.IMG != "" {

		body, err = upload("img", param.IMG, apiUrl)
	} else if "" == param.IMG && param.URL != "" {

		apiUrl += "&url=" + param.URL
		response, err := http.Get(apiUrl)
		if nil != err {
			panic(err.Error())
		}
		body, err = ioutil.ReadAll(response.Body)
		response.Body.Close()
	}

	if nil != err {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &responseValue)
	if nil != err {
		panic(err.Error())
	}

	return responseValue
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