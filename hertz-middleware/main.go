// Package main hertz middleware demo
package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/go-lark/lark-hertz"
)

func main() {
	r := server.Default()
	middleware := larkhertz.NewLarkMiddleware()
	r.Use(middleware.LarkChallengeHandler())
	r.Use(middleware.LarkEventHandler())

	// all supported events
	eventGroup := r.Group("/event")
	{
		eventGroup.Use(middleware.LarkEventHandler())
		eventGroup.POST("/", func(c context.Context, ctx *app.RequestContext) {
			if event, ok := middleware.GetEvent(ctx); ok { // => returns `*lark.EventV2`
				fmt.Println(event)
			}
		})
	}

	// card callback only
	cardGroup := r.Group("/card")
	{
		cardGroup.Use(middleware.LarkCardHandler())
		cardGroup.POST("/callback", func(c context.Context, ctx *app.RequestContext) {
			if card, ok := middleware.GetCardCallback(ctx); ok { // => returns `*lark.EventCardCallback`
				fmt.Println(card)
			}
		})
	}

	r.Spin()
}
