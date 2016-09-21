package main

import (
	"os"
	"fmt"
	"cn/com/faceplusplus/detection"
	"encoding/json"
)

func main() {

	resourcesPath := os.Getenv("GOPATH") + "/resources"
	imgPath := resourcesPath + "/xijingpin.jpg"

	value := detection.DetectFaceImg(detection.RequestParam{IMG:imgPath})

	body, _ := json.Marshal(value)

	fmt.Println(string(body))
}