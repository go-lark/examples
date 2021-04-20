package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-lark/lark"
	larkgin "github.com/go-lark/lark-gin"
)

func main() {
	r := gin.Default()
	middleware := larkgin.NewLarkMiddleware()
	client := &http.Client{}
	r.Use(middleware.LarkMessageHandler())
	// e.g., redirect :9876 to :9875
	r.POST("/handle", func(c *gin.Context) {
		if msg, ok := middleware.GetMessage(c); ok && msg != nil {
			fmt.Println("Handle the msg:", msg)

			lark.PostEvent(client, "http://127.0.0.1:9875/newhandle", *msg)
		}
	})
	r.Run(":9876")
}
