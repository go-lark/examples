package main

import (
	"github.com/go-lark/lark"
)

// should set your own appID and appSecret
const (
	appID     = "cli_xxxxxx"
	appSecret = "el1Pxxxxxx"
)

// 卡片构造工具 https://open.feishu.cn/tool/cardbuilder
// 消息卡片元素介绍 https://open.feishu.cn/document/ukTMukTMukTM/uAzMxEjLwMTMx4CMzETM
// 卡片内的元素可以使用 lark.CardBuilder 的方法创建，通过链式调用设置元素属性
//
// You can use https://open.feishu.cn/tool/cardbuilder to build a card body json
// Refer to https://open.feishu.cn/document/ukTMukTMukTM/uAzMxEjLwMTMx4CMzETM for card elements
// To create a card/element, use methods of lark.CardBuilder;
// Use chained call to set properties of elements, see examples below.
func main() {
	bot := lark.NewChatBot(appID, appSecret)
	bot.GetTenantAccessTokenInternal(true)
	b := lark.NewCardBuilder()
	card := b.Card(
		b.Div(
			b.Field(b.Text("左侧内容")).Short(),
			b.Field(b.Text("Short content field")).Short(),
			b.Field(b.Text("整排内容")),
			b.Field(b.Text("Full-width **Markdown** content").LarkMd()),
		),
		b.Div().
			Text(b.Text("Text Content with extra img")).
			Extra(
				b.Img("img_a7c6aa35-382a-48ad-839d-d0182a69b4dg"),
			),
		b.Action(
			b.Button(b.Text("**Primary**").LarkMd()).Primary(),
			b.Button(b.Text("Confirm")).
				Confirm("Confirm", "Are you sure?"),
			b.Overflow(
				b.Option("1").Text("Option 1"),
				b.Option("2").Text("选项2"),
			).
				Value(map[string]interface{}{"k": "v"}),
		).
			TrisectionLayout(),
		b.Note().
			AddText(b.Text("Note **Text**").LarkMd()).
			AddImage(b.Img("img_a7c6aa35-382a-48ad-839d-d0182a69b4dg")),
	).
		Wathet().
		Title("卡片标题 Card Title")
	msg := lark.NewMsgBuffer(lark.MsgInteractive)
	// card.String() 会将卡片内容渲染成字符串，如构建 CardBuilder 不支持的卡片结构也可以直接传入json串
	// card.String() will render card content into a json string. You can also use raw json content for unsupported card structures.
	om := msg.BindEmail("youremail@example.com").Card(card.String()).Build()
	bot.PostMessage(om)
}
