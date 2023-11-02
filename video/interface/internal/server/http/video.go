package http

import (
	"fmt"
	"github.com/Snaptime23/snaptime-server/v2/tools"
	"github.com/Snaptime23/snaptime-server/v2/video/interface/internal/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
)

type HttpServer struct {
	svr       *service.Service
	videoChan chan string
}

func NewServer(conn *grpc.ClientConn) *HttpServer {
	return &HttpServer{
		svr:       service.NewService(conn),
		videoChan: make(chan string, 20),
	}
}

func (s *HttpServer) UpLoadVideo(c *gin.Context) {
	file, err := c.FormFile("video")
	if tools.HandleError(c, err, "参数填写错误") {
		return
	}
	if file.Size > 100*1024*1024 {
		tools.SendResp(c, 404, "", "文件过大，请上传大小在 100MB 以内的文件")
	}
	path := "./tmp/" + file.Filename
	err = c.SaveUploadedFile(file, path)
	if tools.HandleError(c, err, "文件保存失败") {
		return
	}
	go func() {
		s.videoChan <- path
	}()
	tools.SendResp(c, http.StatusOK, "", "上传成功")
}

func (s *HttpServer) upLoad() {
	for name := range s.videoChan {
		fmt.Println(name)
	}
}
