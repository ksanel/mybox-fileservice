package services

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"mybox.com/services/fileservice/pkg/tools/response"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

// GetFolders
func (h Handler) GetFolders(c *gin.Context) {
	// Get body request
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		res := response.NewResponse(c, err)
		c.JSON(http.StatusBadRequest, res)
	}
	if body == nil {
		res := response.NewResponse(c, "body is empty")
		c.JSON(http.StatusBadRequest, res)
	}

	// binary to json object
	var jsonBody map[string]interface{}
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		res := response.NewResponse(c, "body is empty")
		c.JSON(http.StatusBadRequest, res)
	}

	// Declare bucket query
	query := &storage.Query{
		Prefix:    jsonBody["folder"].(string),
		Delimiter: "/",
	}
	ctx := context.Background()
	it := h.Bucket.Objects(ctx, query)
	folders := []string{}

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if attrs.Prefix != "" {
			folders = append(folders, attrs.Prefix)
			log.Println(attrs.Prefix)
		}
		if attrs.Prefix == "" {
			log.Println(attrs.Name)
			folders = append(folders, attrs.Name)
			log.Println(attrs.Name)
		}
	}

	res := response.NewResponse(c, folders)
	c.JSON(http.StatusOK, res)

}
