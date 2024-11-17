package request

type CreatePostRequest struct {
    Title string `json:"title"`
    Content string `json:"content"`
}

type GetPostRequest struct {
    Id uint64 `json:"id"`
}

type GetPostsRequest struct {
}

type UpdatePostRequest struct {
    Title string `json:"title"`
    Content string `json:"content"`
}

type DeletePostRequest struct {
    Id uint64 `json:"id"`
}
