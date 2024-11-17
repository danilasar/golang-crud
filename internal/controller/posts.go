package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/danilasar/golang-crud/internal/core/common/router"
    "github.com/danilasar/golang-crud/internal/core/model/request"
    "github.com/danilasar/golang-crud/internal/core/model/response"
    "github.com/danilasar/golang-crud/internal/core/port/service"
)

type PostController struct {
    gin *gin.Engine
    postService service.PostService
}

func NewPostController(
    gin *gin.Engine
    postService service.PostService,
) PostController {
    return PostController(
        gin: gin,
        postService: postService
    )
}

func (p PostController) InitRouter() {
    api := p.gin.Group("/posts")
    router.Put(api, "/", p.create)
}

func (p PostController) create(c *gin.Context) {
    var req request.CreatePostRequest
    err := c.ShouldBindJSON(&req)
    if err != nil {
        c.AbortWithCode(http.StatusBadRequest)
        return
    }
    resp := p.postService.Create(req)
    c.JSON(resp.StatusCode, resp.Content)
}
