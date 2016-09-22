package main

import (
	"cn/com/faceplusplus/detection"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	resourcesPath := os.Getenv("GOPATH") + "/resources"
	imgPath := resourcesPath + "/xijingpin.jpg"

	var param detection.DetectRequestParam
	param.UPLOAD.IMG = imgPath
	value, _ := detection.DetectFaceImg(param)

	//value, _ := detection.LandmarkFaceImg(detection.LandmarkRequestParam{FACE_ID: "0814532b8ca9afd3341dce872750792e", TYPE: "25p"})

	body, _ := json.Marshal(value)
	fmt.Println(string(body))
}
