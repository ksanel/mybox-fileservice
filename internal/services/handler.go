package services

import (
	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Router *gin.Engine
	Bucket *storage.BucketHandle
}

func Run(router *gin.Engine, bkt *storage.BucketHandle, addr string) error {
	h := Handler{
		Router: router,
		Bucket: bkt,
	}

	router.GET("/v1/API/folders", h.GetFolders)

	err := router.Run(addr)
	if err != nil {
		return err
	}
	return nil
}
