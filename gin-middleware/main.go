// Package main .
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
	middleware.BindURLPrefix("/lark/lark") // enable prefix binding if you need
	// middleware.WithEncryption("<encrypt-key>")  // enable encryption if you need
	// middleware.WithTokenVerification("<token>") // enable token verification if you need

	baseGroup := r.Group("/lark")
	baseGroup.Use(middleware.LarkChallengeHandler())

	// Events v2
	v2Group := baseGroup.Group("/v2")
	v2Group.Use(middleware.LarkEventHandler())
	{
		r.POST("/", func(c *gin.Context) {
			if event, ok := middleware.GetEvent(c); ok && event != nil {
				if event.Header.EventType == lark.EventTypeMessageReceived {
					if messageReceived, err := event.GetMessageReceived(); err == nil {
						fmt.Println(messageReceived)
					}
				} else if event.Header.EventType == lark.EventTypeUserAdded {
					if userAdded, err := event.GetUserAdded(); err == nil {
						fmt.Println(userAdded)
					}
				}
				// ... and other events
			}
		})

	}

	// Card callback
	cardGroup := baseGroup.Group("/card")
	cardGroup.Use(middleware.LarkCardHandler())
	{
		cardGroup.POST("/", func(c *gin.Context) {
			if event, ok := middleware.GetCardCallback(c); ok && event != nil {
				fmt.Println("Handle the event", event)
			}
		})

	}

	// Legacy event v1
	v1Group := baseGroup.Group("/v1")
	v1Group.Use(middleware.LarkMessageHandler())
	{
		v1Group.POST("/lark", func(c *gin.Context) {
			if msg, ok := middleware.GetMessage(c); ok && msg != nil {
				fmt.Println("Handle the msg:", msg)
			}
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}

	r.Run(":8044")
}
