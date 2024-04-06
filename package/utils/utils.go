package utils

import (
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
)
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
