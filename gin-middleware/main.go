// Package main .
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	larkgin "github.com/go-lark/lark-gin/v2"
	"github.com/go-lark/lark/v2"
)

func main() {
	r := gin.Default()
	middleware := larkgin.NewLarkMiddleware()
	middleware.BindURLPrefix("/lark/lark") // enable prefix binding if you need
	// middleware.WithEncryption("<encrypt-key>")  // enable encryption if you need
	// middleware.WithTokenVerification("<token>") // enable token verification if you need

	baseGroup := r.Group("/lark")
	baseGroup.Use(middleware.LarkChallengeHandler())

	// Events
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

	r.Run(":8044")
}
