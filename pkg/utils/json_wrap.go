package utils

import (
	. "docker-/internal/log"
	"encoding/json"
)

func Unmarshal(bytes []byte, v interface{}) {
	if err := json.Unmarshal(bytes, v); err != nil {
		Log.Warnf("unmarshal error %v", bytes)
	}
}
