package main

import (
	"cn/com/faceplusplus/detection"
	"encoding/json"
	"fmt"
)

func main() {

	//resourcesPath := os.Getenv("GOPATH") + "/resources"
	//imgPath := resourcesPath + "/xijingpin.jpg"
	//
	//value, _ := detection.DetectFaceImg(detection.DetectRequestParam{IMG:imgPath})

	value, _ := detection.LandmarkFaceImg(detection.LandmarkRequestParam{FACE_ID: "0814532b8ca9afd3341dce872750792e", TYPE: "25p"})

	body, _ := json.Marshal(value)
	fmt.Println(string(body))
}
