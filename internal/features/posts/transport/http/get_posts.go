package posts_transport_http

import (
	"net/http"

	"github.com/AlexLopatin-hub/thewall/internal/core/domain"
	core_http_response "github.com/AlexLopatin-hub/thewall/internal/core/transport/http/response"
)

type GetPostResponse PostDTOResponse

func (h *PostsHTTPHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	responseHandler := core_http_response.NewHTTPResponseHandler(w)

	posts, err := h.postsService.GetPosts(ctx)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"could not get posts",
		)
		return
	}

	response := postDTOFromDomains(posts)

	responseHandler.JSONResponse(http.StatusOK, response)
}

func postDTOFromDomains(posts []domain.Post) []GetPostResponse {
	dtoPosts := make([]GetPostResponse, len(posts))

	for index, post := range posts {
		dtoPosts[index] = GetPostResponse(postDTOFromDomain(post))
	}

	return dtoPosts
}
