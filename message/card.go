// Package main card example
package main

import (
	"context"

	"github.com/go-lark/lark/v2"
)

// You can use https://open.feishu.cn/tool/cardbuilder to build a card body json
// Refer to https://open.feishu.cn/document/ukTMukTMukTM/uAzMxEjLwMTMx4CMzETM for card elements
// To create a card/element, use methods of lark.CardBuilder;
// Use chained call to set properties of elements, see examples below.
func card(ctx context.Context, bot *lark.Bot) {
	b := lark.NewCardBuilder()

	// 普通卡片
	// Normal card
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
	// card.String() will render card content into a json string. You can also use raw json content for unsupported card structures.
	om := msg.BindEmail("youremail@example.com").Card(card.String()).Build()
	bot.PostMessage(ctx, om)

	// Card with I18N
	i18nCard := b.I18N.
		Card(
			b.I18N.WithLocale(
				lark.LocaleEnUS,
				b.Div(
					b.Field(b.Text("English Content")),
				),
			),
			b.I18N.WithLocale(
				lark.LocaleZhCN,
				b.Div(
					b.Field(b.Text("中文内容")),
				),
			),
			b.I18N.WithLocale(
				lark.LocaleJaJP,
				b.Div(
					b.Field(b.Text("日本語コンテンツ")),
				),
			),
		).
		Title(
			b.I18N.LocalizedText(lark.LocaleEnUS, "English Title"),
			b.I18N.LocalizedText(lark.LocaleZhCN, "中文标题"),
			b.I18N.LocalizedText(lark.LocaleJaJP, "日本語タイトル"),
		).
		Red()
	om = msg.BindEmail("youremail@example.com").Card(i18nCard.String()).Build()
	bot.PostMessage(ctx, om)
}
