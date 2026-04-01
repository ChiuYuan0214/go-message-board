package routes

import (
	"fmt"
	"net/http"
	"path/filepath"
	"security/utils"
	"strconv"
)

func (h *ProfileHandler) handleUploadImage(writer http.ResponseWriter, req *http.Request) {
	setHeader(writer, "json")
	res, status := h.uploadImageMap.useHandler(writer, req)
	DoResponse(res, status, writer)
}

func (h *ProfileHandler) upload(req *http.Request) (res interface{}, statusCode int) {
	userId := getUserIdFromContext(req)
	err := req.ParseMultipartForm(10 << 19)
	if isErr(err) {
		return newRes("fail").message("image over 5MB or error in parsing."), http.StatusBadRequest
	}
	if len(req.MultipartForm.File) < 1 {
		return newRes("fail").message("file cannot be empty."), http.StatusBadRequest
	}

	desc := req.FormValue("desc")
	file, handler, err := req.FormFile("file")
	if isErr(err) {
		return newRes("fail").message("error in parsing."), http.StatusBadRequest
	}
	defer file.Close()

	fileName := fmt.Sprintf("img_%s%s", strconv.FormatUint(userId, 10), filepath.Ext(handler.Filename))
	if !utils.UploadFile("images", fileName, file) {
		return newRes("fail").message("something went wrong."), http.StatusInternalServerError
	}

	if message, status := h.profileService.InsertProfileImageInfo(&userId, &fileName, &desc); status != 0 {
		return newRes("fail").message(message), status
	}

	return newRes("success"), http.StatusOK
}
