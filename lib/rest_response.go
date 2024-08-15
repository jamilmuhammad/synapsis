package lib

import (
	"encoding/json"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Error *Error      `json:"error"`
	Data  interface{} `json:"data"`
}

func failResponseWriter(sp opentracing.Span, w http.ResponseWriter, err error, errStatusCode int) {
	w.Header().Set("Content-Type", "application/json")

	var resp Response
	w.WriteHeader(errStatusCode)
	resp.Error = &Error{Code: errStatusCode, Message: err.Error()}
	resp.Data = nil

	responseBytes, _ := json.Marshal(resp)
	w.Write(responseBytes)

	sp.SetTag("error", true)
	LogResponse(sp, resp)
}

func successResponseWriter(sp opentracing.Span, w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	var resp Response
	w.WriteHeader(statusCode)
	resp.Error = nil
	resp.Data = data

	responseBytes, _ := json.Marshal(resp)
	w.Write(responseBytes)

	LogResponse(sp, resp)
}

func WriteResponse(sp opentracing.Span, w http.ResponseWriter, err error, data any) {
	switch err.(type) {
	case *ErrNotFound, ErrNotFound:
		failResponseWriter(sp, w, err, http.StatusNotFound)
	case *ErrBadRequest, ErrBadRequest:
		failResponseWriter(sp, w, err, http.StatusBadRequest)
		return
	case *ErrForbidden, ErrForbidden:
		failResponseWriter(sp, w, err, http.StatusForbidden)
		return
	case *ErrDuplicate, ErrDuplicate:
		failResponseWriter(sp, w, err, http.StatusConflict)
		return
	case nil:
		successResponseWriter(sp, w, data, http.StatusOK)
		return
	default:
		failResponseWriter(sp, w, err, http.StatusInternalServerError)
		return
	}
}
