package response

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"messages"`
	Content interface{} `json:"content"`
	Time    int64       `json:"time"`
	Request Remote      `json:"request"`
}

type Remote struct {
	ReqHeaders http.Header   `json:"reqHeaders"`
	ReqParams  string        `json:"reqParams"`
	ReqBody    io.ReadCloser `json:"reqBody"`
	RemoteIp   string        `json:"remoteIp"`
	RemotePort string        `json:"remotePort"`
	ReqURL     string        `json:"reqUrl"`
}

func NewResponse(c *gin.Context, content interface{}) Response {
	remote := Remote{
		ReqHeaders: c.Request.Header,
		ReqParams:  c.Request.URL.Query().Encode(),
		ReqBody:    c.Request.Body,
		RemoteIp:   c.ClientIP(),
		RemotePort: c.Request.RemoteAddr,
		ReqURL:     c.Request.URL.Path,
	}

	return Response{
		Code:    200,
		Msg:     "success",
		Content: content,
		Time:    time.Now().Unix(),
		Request: remote,
	}
}
