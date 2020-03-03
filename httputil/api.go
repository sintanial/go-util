package httputil

import (
	"github.com/ansel1/merry"
	"github.com/sintanial/go-lerry"
	"net/http"
)

type ApiResponse struct {
	Status bool        `json:"status"`
	Reason string      `json:"reason"`
	Data   interface{} `json:"data"`
}

func SuccessApiResponse(data interface{}) ApiResponse {
	return ApiResponse{true, "", data}
}

func WriteSuccessApiResponse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return WriteJson(SuccessApiResponse(data), w, r)
}

func FailedApiResponse(reason string) ApiResponse {
	return ApiResponse{false, reason, nil}
}

func WriteFailedApiResponse(w http.ResponseWriter, r *http.Request, reason string) error {
	return WriteJson(FailedApiResponse(reason), w, r)
}

type ApiHandler func(w http.ResponseWriter, r *http.Request) (interface{}, error)

func WrapApiHandler(handler ApiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := handler(w, r)
		if err == nil {
			WriteSuccessApiResponse(w, r, data)
			return
		}

		msg := merry.UserMessage(err)
		if msg == "" {
			msg = merry.Message(err)
		}

		WriteFailedApiResponse(w, r, msg)

		lerry.Log(err)
	}
}
