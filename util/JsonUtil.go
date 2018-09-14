package util

import (
	"encoding/json"
	log "github.com/jeanphorn/log4go"
)

type JsonUtil struct {
}

// 将对象转换成json字符串
// obj 实体对象
func (*JsonUtil) ToJson(obj interface{}) (result string) {

	if obj == nil {
		return ""
	}
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Error("json convert error:", err)
	}
	if bytes != nil {
		return string(bytes)
	}
	return ""
}

// 将json字符串转换成对象
// jsonString json字符串
// 需要转换对的对象，必须先实例化
func (*JsonUtil) ToObject(jsonString string, obj interface{}) {

	if obj == nil {
		log.Error("json object is null")
	}
	err := json.Unmarshal([]byte(jsonString), &obj)
	if err != nil {
		log.Error("json convert error:", err)
	}
}
