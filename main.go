package main

import (
    "github.com/gin-gonic/gin"
    "time"
    "log"
)

func getLongSyncHandler(context *gin.Context) {
    time.Sleep(5 * time.Second)
    log.Println("Done! in path " + context.Request.URL.Path)
}

func getLongAsyncHandler(originalContext *gin.Context) {
    context := originalContext.Copy()
    go func() {
        time.Sleep(5 * time.Second)
        log.Println("Done! in path " + context.Request.URL.Path)
    }()
}

func main() {
    engine := gin.Default()
    engine.GET("/long_sync", getLongSyncHandler)
    engine.GET("/long_async", getLongAsyncHandler)
    engine.Run()
}
