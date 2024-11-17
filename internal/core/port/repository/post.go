package repository

import (
    "errors"
    //"github.com/danilasar/golang-crud/internal/port/repository"
    "github.com/danilasar/golang-crud/internal/core/dto"
)

type PostRepository interface {
    //db repository.Database
    Insert(post dto.PostDTO) error
    Get(postid uint) (dto.PostDTO, error)
    GetAll(selectors dto.PostDTO) (dto.PostDTO[], error)
    Update(post dto.PostDTO) error
    Delete(post dto.PostDTO) error
}

var (
    FailedToInsert = errors.New("failed to create post")
)
