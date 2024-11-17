package service

import (
    "github.com/danilasar/golang-crud/internal/core/model/request"
    "github.com/danilasar/golang-crud/internal/core/model/response"
)

type PostService interface {
    Create(request *request.CretePostRequest)

}


