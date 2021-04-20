package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	larkgin "github.com/go-lark/lark-gin"
)

func main() {
	r := gin.Default()
	middleware := larkgin.NewLarkMiddleware()
	larkGroup := r.Group("/lark")
	larkGroup.Use(middleware.LarkChallengeHandler())
	larkGroup.Use(middleware.LarkMessageHandler())
	larkGroup.POST("/lark", func(c *gin.Context) {
		if msg, ok := middleware.GetMessage(c); ok && msg != nil {
			fmt.Println("Handle the msg:", msg)
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run(":8044")
}
