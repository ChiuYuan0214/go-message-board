package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type HandlerType func(req *http.Request) (res interface{}, statusCode int)
type MethodMapType map[string]HandlerType

func (methodMap *MethodMapType) get(f HandlerType) *MethodMapType {
	(*methodMap)["GET"] = f
	return methodMap
}

func (methodMap *MethodMapType) put(f HandlerType) *MethodMapType {
	(*methodMap)["PUT"] = f
	return methodMap
}

func (methodMap *MethodMapType) post(f HandlerType) *MethodMapType {
	(*methodMap)["POST"] = f
	return methodMap
}

func (methodMap *MethodMapType) delete(f HandlerType) *MethodMapType {
	(*methodMap)["DELETE"] = f
	return methodMap
}

func trigger(handler HandlerType, request *http.Request) (res interface{}, statusCode int) {
	if reflect.TypeOf(handler).Kind() != reflect.Func {
		return nil, http.StatusInternalServerError
	}
	return handler(request)
}

func setContentType(writer http.ResponseWriter, contentType string) {
	switch contentType {
	default:
		writer.Header().Set("Content-Type", "application/json")
	}
}

func (this *MethodMapType) useHandler(request *http.Request) (res interface{}, statusCode int) {
	var method string
	switch request.Method {
	case http.MethodGet:
		method = "GET"
	case http.MethodPut:
		method = "PUT"
	case http.MethodPost:
		method = "POST"
	case http.MethodDelete:
		method = "DELETE"
	}
	handler, exist := (*this)[method]
	if !exist {
		return newRes("fail").message("Method not allowed."), http.StatusBadRequest
	}
	return trigger(handler, request)
}

func getQuery(req *http.Request) url.Values {
	return req.URL.Query()
}

func getParam(req *http.Request, key string) string {
	return req.URL.Query().Get(key)
}

func getUserIdFromContext(req *http.Request) int64 {
	return req.Context().Value("userId").(int64)
}

func getUserIdFromQuery(req *http.Request) int64 {
	userIdStr := getParam(req, "userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return 0
	}
	return userId
}

func getPageSize(req *http.Request) (int64, int64) {
	page := getParam(req, "page")
	size := getParam(req, "size")
	pageNum, pageErr := strconv.ParseInt(page, 10, 64)
	if pageErr != nil {
		pageNum = 1
	}
	sizeNum, sizeErr := strconv.ParseInt(size, 10, 64)
	if sizeErr != nil {
		sizeNum = 10
	}
	return pageNum, sizeNum
}

func constructParams(data *map[string]interface{}) (string, []interface{}) {
	var setValues []string
	var args []interface{}
	for key, val := range *data {
		setValues = append(setValues, fmt.Sprintf("%s = ?", key))
		args = append(args, val)
	}
	return strings.Join(setValues, ", "), args
}

type ResponseMessage map[string]interface{}

func newRes(status string) *ResponseMessage {
	res := ResponseMessage{}
	res["status"] = status
	return &res
}

func (res *ResponseMessage) message(mes string) *ResponseMessage {
	(*res)["message"] = mes
	return res
}

func (res *ResponseMessage) setId(id int64) *ResponseMessage {
	(*res)["id"] = id
	return res
}

func (res *ResponseMessage) setItem(key string, item interface{}) *ResponseMessage {
	(*res)[key] = item
	return res
}

func (res *ResponseMessage) setList(key string, list interface{}) *ResponseMessage {
	(*res)[key] = list
	return res
}

func DoResponse(res interface{}, statusCode int, writer http.ResponseWriter) {
	writer.WriteHeader(statusCode)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(res)
}

func isErr(err error) bool {
	if err != nil {
		log.Println("error:", err)
		return true
	}
	return false
}
