// Package main notification example
package main

import (
	"context"

	"github.com/go-lark/lark/v2"
)

func main() {
	ctx := context.Background()
	bot := lark.NewNotificationBot("https://open.feishu.cn/open-apis/bot/v2/hook/7b01451f-113b-4296-8f0d-9615499d6545")
	mbText := lark.NewMsgBuffer(lark.MsgText)
	mbText.Text("hello")
	bot.PostNotification(ctx, mbText.Build())

	mbPost := lark.NewMsgBuffer(lark.MsgPost)
	mbPost.Post(lark.NewPostBuilder().Title("hello").TextTag("world", 1, true).Render())
	bot.PostNotification(ctx, mbPost.Build())

	mbImg := lark.NewMsgBuffer(lark.MsgImage)
	mbImg.Image("img_a97c1597-9c0a-47c1-9fb4-dd3e5e37ac9g")
	bot.PostNotification(ctx, mbImg.Build())

	mbShareGroup := lark.NewMsgBuffer(lark.MsgShareCard)
	mbShareGroup.ShareChat("oc_1c09434c2264eb52dc7667895a7fac6d")
	bot.PostNotification(ctx, mbShareGroup.Build())
}
