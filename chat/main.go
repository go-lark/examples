// Package main group examples
package main

import (
	"context"

	"github.com/go-lark/lark/v2"
)

// should set your own appID and appSecret
const (
	appID     = "cli_xxxxxx"
	appSecret = "el1Pxxxxxx"
)

func main() {
	ctx := context.Background()
	bot := lark.NewChatBot(appID, appSecret)
	userOpenID := "ou_xxxxx"
	// Create Chat
	req := lark.CreateChatRequest{
		Name:     "hello",
		ChatMode: "group",
		ChatType: "public",
	}
	chatInfo, _ := bot.CreateChat(ctx, req)
	chatID := chatInfo.Data.ChatID

	// Get Chat Info
	bot.GetChat(ctx, chatID)

	// Add Group Member
	bot.AddChatMember(ctx, chatID, []string{userOpenID})

	// Remove Chat Member
	bot.RemoveChatMember(ctx, chatID, []string{userOpenID})

	// List Chat
	bot.ListChat(ctx, "ByCreateTimeAsc", "", 10)
}
