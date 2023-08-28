// Package main notification example
package main

import (
	"github.com/go-lark/lark"
)

func main() {
	bot := lark.NewNotificationBot("https://open.feishu.cn/open-apis/bot/hook/c6c2eae4856a4c66866d25780fa87c30")
	bot.PostNotification("go-lark", "example")
	// Title could be empty
	bot.PostNotification("", "content only")

	botV2 := lark.NewNotificationBot("https://open.feishu.cn/open-apis/bot/v2/hook/7b01451f-113b-4296-8f0d-9615499d6545")
	mbText := lark.NewMsgBuffer(lark.MsgText)
	mbText.Text("hello")
	botV2.PostNotificationV2(mbText.Build())

	mbPost := lark.NewMsgBuffer(lark.MsgPost)
	mbPost.Post(lark.NewPostBuilder().Title("hello").TextTag("world", 1, true).Render())
	botV2.PostNotificationV2(mbPost.Build())

	mbImg := lark.NewMsgBuffer(lark.MsgImage)
	mbImg.Image("img_a97c1597-9c0a-47c1-9fb4-dd3e5e37ac9g")
	botV2.PostNotificationV2(mbImg.Build())

	mbShareGroup := lark.NewMsgBuffer(lark.MsgShareCard)
	mbShareGroup.ShareChat("oc_1c09434c2264eb52dc7667895a7fac6d")
	botV2.PostNotificationV2(mbShareGroup.Build())
}
