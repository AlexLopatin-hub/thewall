package posts_transport_http

import (
	"net/http"

	"github.com/AlexLopatin-hub/thewall/internal/core/domain"
	core_http_request "github.com/AlexLopatin-hub/thewall/internal/core/transport/http/request"
	core_http_response "github.com/AlexLopatin-hub/thewall/internal/core/transport/http/response"
)

type CreatePostRequest struct {
	Text string `json:"text"`
}

type CreatePostResponse PostDTOResponse

func (h *PostsHTTPHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	responseHandler := core_http_response.NewHTTPResponseHandler(w)

	var request CreatePostRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request")
		return
	}

	postDomain := domainFromDTO(request)

	postDomain, err := h.postsService.CreatePost(ctx, postDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create post")
		return
	}

	response := CreatePostResponse(postDTOFromDomain(postDomain))

	responseHandler.JSONResponse(http.StatusOK, response)
}

func domainFromDTO(dto CreatePostRequest) domain.Post {
	return domain.NewPostUninitialized(
		dto.Text,
	)
}
