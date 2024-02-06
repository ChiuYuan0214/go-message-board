package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

func ParseBody(body io.ReadCloser, containerAddr interface{}) (string, int) {
	decoder := json.NewDecoder(body)
	if !decoder.More() {
		return "body was empty.", http.StatusBadRequest
	}

	err := decoder.Decode(containerAddr)
	if err != nil {
		return "body format was wrong.", http.StatusBadRequest
	}

	return "", 0
}

func GetQuery(req *http.Request) url.Values {
	return req.URL.Query()
}

func GetParam(req *http.Request, key string) string {
	return req.URL.Query().Get(key)
}

func ConstructParams(data *map[string]interface{}) (string, []interface{}) {
	var setValues []string
	var args []interface{}
	for key, val := range *data {
		setValues = append(setValues, fmt.Sprintf("%s = ?", key))
		args = append(args, val)
	}
	return strings.Join(setValues, ", "), args
}

func ConstructParamsFromStruct(data interface{}) (cols string, args []interface{}) {
	dataType := reflect.TypeOf(data)
	if dataType.Kind() != reflect.Struct {
		return
	}

	mapping := make(map[string]interface{})
	dataValue := reflect.ValueOf(data)

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i).Tag.Get("json")
		val := dataValue.Field(i).Interface()
		if val != "" && val != 0 {
			mapping[field] = val
		}
	}

	return ConstructParams(&mapping)
}
