package response

type Response struct {
    Content interface{}
    StatusCode uint16
}

type CreatePostResponse struct {
    Id uint64 `json:"id"`
}

type GetPostResponse struct {
    Id uint64 `json:"id"`
    Title string `json:"title"`
    Content string `json:"content"`
}

type GetPostsResponse struct {
    Posts GetPostResponse[] `json:"posts"`
}

type MessageResponse struct {
    Message string `json:"message"`
}

type EmptyResponse struct {

}
