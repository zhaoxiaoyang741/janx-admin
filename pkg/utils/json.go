package utils

import (
	"encoding/json"
)

// Struct2Json 将结构体转换为JSON字符串
func Struct2Json(obj interface{}) string {
	str, err := json.Marshal(obj)
	if err != nil {
		return "change to json error"
	}
	return string(str)
}
