// Package main at message examples
package main

import (
	"context"
	"fmt"

	"github.com/go-lark/lark/v2"
)

func at(ctx context.Context, bot *lark.Bot) {
	openID := "ou_xxxxxxx" // you may get this from incoming message or other API
	mb := lark.NewMsgBuffer(lark.MsgText)
	// Method 1
	text := fmt.Sprintf("AT Message 1 <at user_id=\"%s\">%s</at>", openID, "name")
	msg1 := mb.BindOpenChatID("oc_xxxxxx").Text(text).Build()
	bot.PostMessage(ctx, msg1)
	mb.Clear()
	// Method 2
	tb2 := lark.NewTextBuilder().Text("AT Message 2").Mention(openID)
	msg2 := mb.BindOpenChatID("oc_xxxxxx").Text(tb2.Render()).Build()
	bot.PostMessage(ctx, msg2)
	mb.Clear()
	// AT ALL, should use cautiously to prevent disturbing
	tb3 := lark.NewTextBuilder().Text("AT ALL").MentionAll()
	msgAll := mb.BindOpenChatID("oc_xxxxxx").Text(tb3.Render()).Build()
	bot.PostMessage(ctx, msgAll)
}

func image(ctx context.Context, bot *lark.Bot) {
	// Upload to Lark at first
	resp, err := bot.UploadImage(ctx, "/path/to/your.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Send with ImageKey
	mb := lark.NewMsgBuffer(lark.MsgImage)
	msg := mb.BindEmail("xxxxxx@example.com").Image(resp.Data.ImageKey).Build()
	bot.PostMessage(ctx, msg)
}

func post(ctx context.Context, bot *lark.Bot) {
	msg := lark.NewMsgBuffer(lark.MsgPost)
	postContent := lark.NewPostBuilder().
		Title("post title").
		TextTag("hello, world", 1, true).
		LinkTag("Google", "https://google.com/").
		AtTag("www", "oc_xxxxxx").
		ImageTag("d9f7d37e-c47c-411b-8ec6-9861132e6986", 300, 300).
		Render()
	om := msg.BindOpenChatID("oc_xxxxxx").Post(postContent).Build()
	bot.PostMessage(ctx, om)
}

func shareChat(ctx context.Context, bot *lark.Bot) {
	mb := lark.NewMsgBuffer(lark.MsgShareCard)
	msg := mb.BindEmail("youremail@example.com").ShareChat("oc_xxxxx").Build()
	bot.PostMessage(ctx, msg)
}
