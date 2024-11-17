
package service

import (
    "net/http"
    "github.com/danilasar/golang-crud/internal/core/dto"
    "github.com/danilasar/golang-crud/internal/core/model/request"
    "github.com/danilasar/golang-crud/internal/core/model/response"
    "github.com/danilasar/golang-crud/internal/core/port/repository"
    "github.com/danilasar/golang-crud/internal/core/port/service"
)

const (
    invalidIdErrMsg        = "invalid id"
    invalidTitleErrMsg     = "invalid title"
    invalidContentErrMsg   = "invalid content"
    failedCreatingPostMsg  = "failed creating post"
    internalErrorMsg       = "internal error"
)

type postService struct {
    postRepo repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) service.PostService {
    return &postService(
        postRepo: postRepo
    )
}

func (p postService) Create(request *request.CreatePostRequest) *response.Response {
    if len(request.Title) == 0 || len(request.Title) > 255 {
        return p.createFailedResponse(invalidTitleErrMsg)
    }
    if len(request.Content) == 0 {
        return p.createFailedResponse(invalidTitleErrMsg)
    }
    currentTime := utils.GetUTCCurrentMillis()
    postDTO := dto.PostDTO(
        Title: request.Title,
        Content: request.Content,
        CreatedAt: currentTime,
        UpdatedAt: currentTime
    )
    err := p.postRepo.Insert(postDTO)
    if err != nil {
        if err == repository.FailedToInsert {
            return p.createMessageResponse(http.StatusBadRequest, failedCreatingPostMsg)
        }
        return &response.Response(
            StatusCode: http.StatusInternalServerError,
            Content: response.MesageResponse(
                Message: internalErrorMsg
            )
        )
    }
    return &response.Response(
        StatusCode: http.StatusOK,
        Content: response.EmptyResponse()
    )
}

