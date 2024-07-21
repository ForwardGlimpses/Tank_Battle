package json

import "encoding/json"

var (
	Marshal   = json.Marshal
	Unmarshal = json.Unmarshal
)

func MarshalToString(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}
