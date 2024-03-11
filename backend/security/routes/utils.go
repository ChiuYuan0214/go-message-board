package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

type HandlerType func(req *http.Request) (res interface{}, statusCode int)
type MethodMapType map[string]HandlerType

func (mm *MethodMapType) get(f HandlerType) *MethodMapType {
	(*mm)["GET"] = f
	return mm
}

func (mm *MethodMapType) put(f HandlerType) *MethodMapType {
	(*mm)["PUT"] = f
	return mm
}

func (mm *MethodMapType) post(f HandlerType) *MethodMapType {
	(*mm)["POST"] = f
	return mm
}

func (mm *MethodMapType) delete(f HandlerType) *MethodMapType {
	(*mm)["DELETE"] = f
	return mm
}

func setHeader(writer http.ResponseWriter, contentType string) {
	SetCORS(writer)
	switch contentType {
	default:
		writer.Header().Set("Content-Type", "application/json")
	}
}

func (methodMap *MethodMapType) useHandler(writer http.ResponseWriter, request *http.Request) (res interface{}, statusCode int) {
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
	case http.MethodOptions:
		return newRes("success"), http.StatusOK
	}
	handler, exist := (*methodMap)[method]
	if !exist {
		return newRes("fail").message("Method not allowed."), http.StatusBadRequest
	}
	return trigger(handler, request)
}

func trigger(handler HandlerType, request *http.Request) (res interface{}, statusCode int) {
	if reflect.TypeOf(handler).Kind() != reflect.Func {
		return nil, http.StatusInternalServerError
	}
	return handler(request)
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

func SetCORS(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
	writer.Header().Add("Access-Control-Allow-Credentials", "true")
}

func DoResponse(res interface{}, statusCode int, writer http.ResponseWriter) {
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(res)
}

func isErr(err error) bool {
	if err != nil {
		log.Println("error:", err)
		return true
	}
	return false
}

func getUserIdFromContext(req *http.Request) uint64 {
	return req.Context().Value("userId").(uint64)
}

func getParam(req *http.Request, key string) string {
	return req.URL.Query().Get(key)
}
