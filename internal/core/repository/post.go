package repository

import (
    "github.com/danilasar/golang-crud/internal/core/dto"
    "github.com/danilasar/golang-crud/internal/core/port/repository"
)

type postRepository struct {
    db repository.Database
}


func NewPostRepository() (repository.PostRepository, error) {
    db, err := repository.GetDatabase()
    if err != nil {
        return err
    }
    return &postRepository(
        db: db
    )
}

func (p postRepository) Insert(post dto.Post) error {
    result, err := db.GetDB().Exec(
        "INSERT INTO posts(`title`, `content`, `created_at`, `updated_at`) VALUES(?, ?, ?, ?)",
        post.Title,
        post.Content,
        post.CreatedAt,
        post.UpdatedAt
    )
    if err != nil {
        return err
    }
    numRow, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if numRow != numberRowInserted {
        return repository.FailedToInsert
    }

    return nil
}
/*
    Get(postid uint) (dto.PostDTO, error)
    GetAll(selectors dto.PostDTO) (dto.PostDTO[], error)
    Update(post dto.PostDTO) error
    Delete(post dto.PostDTO) error
*/
//func (p postRepository) Get
