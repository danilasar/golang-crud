package server

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/danilasar/golang-crud/internal/infra/config"
)

const defaultHost = "0.0.0.0"

type HttpServer interface {
    Start()
    Stop()
}

type httpServer struct {
    Port   uint
    server *http.Server
}

func NewHttpServer(router *gin.Engine, config config.HttpServerConfig) HttpServer {
    return &httpServer{
        Port: config.Port,
        server: &http.Server{
            Addr:    fmt.Sprintf("%s:%d", defaultHost, config.Port),
            Handler: router,
        },
    }
}

func (httpServer httpServer) Start() {
    go func() {
        if err := httpServer.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf(
                "failed to start HttpServer on port %d, err=%s",
                httpServer.Port, err.Error(),
            )
        }
    }()

    log.Printf("Service started on port %d", httpServer.Port)
}

func (httpServer httpServer) Stop() {
    ctx, cancel := context.WithTimeout(
        context.Background(), time.Duration(3)*time.Second,
    )
    defer cancel()

    if err := httpServer.server.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown, err=%s", err.Error())
    }
}
