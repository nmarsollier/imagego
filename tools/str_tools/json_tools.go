package str_tools

import "encoding/json"

func ToJson(obj interface{}) string {
	jsonData, _ := json.Marshal(obj)
	return string(jsonData)
}
