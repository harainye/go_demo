package gintest

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Post(c *gin.Context) {
	id := c.PostForm("id")
	page := c.PostForm("page")
	name := c.PostForm("name")
	message := c.PostForm("message")

	log.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
}
