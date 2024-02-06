package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

type HandlerType func(req *http.Request) (res interface{}, statusCode int)
type MethodMapType map[string]HandlerType

func (this *MethodMapType) get(f HandlerType) *MethodMapType {
	(*this)["GET"] = f
	return this
}

func (this *MethodMapType) put(f HandlerType) *MethodMapType {
	(*this)["PUT"] = f
	return this
}

func (this *MethodMapType) post(f HandlerType) *MethodMapType {
	(*this)["POST"] = f
	return this
}

func (this *MethodMapType) delete(f HandlerType) *MethodMapType {
	(*this)["DELETE"] = f
	return this
}

func setContentType(writer http.ResponseWriter, contentType string) {
	switch contentType {
	default:
		writer.Header().Set("Content-Type", "application/json")
	}
}

func (methodMap *MethodMapType) useHandler(request *http.Request) (res interface{}, statusCode int) {
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

func (res *ResponseMessage) setList(key string, list []interface{}) *ResponseMessage {
	(*res)[key] = list
	return res
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

func getUserIdFromContext(req *http.Request) int64 {
	return req.Context().Value("userId").(int64)
}
