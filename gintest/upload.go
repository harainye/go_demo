package gintest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Upload(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	dst := "./" + file.Filename
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
