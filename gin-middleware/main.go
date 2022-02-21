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
	middleware.BindURLPrefix("/lark/lark")           // enable prefix binding if you need
	middleware.WithEncryption("<encrypt-key>")  // enable encryption if you need
	middleware.WithTokenVerification("<token>") // enable token verification if you need

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
