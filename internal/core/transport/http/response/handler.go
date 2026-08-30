package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/AlexLopatin-hub/thewall/internal/core/errors"
)

type HTTPResponseHandler struct {
	w http.ResponseWriter
}

func NewHTTPResponseHandler(
	w http.ResponseWriter,
) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		w: w,
	}
}

func (h *HTTPResponseHandler) JSONResponse(
	statusCode int,
	responseBody any,
) {
	h.w.WriteHeader(statusCode)

	if err := json.NewEncoder(h.w).Encode(responseBody); err != nil {
		fmt.Println("error writing HTTP response")
	}
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	var statusCode int

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
	default:
		statusCode = http.StatusInternalServerError
	}

	responseBody := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	h.JSONResponse(statusCode, responseBody)
}
