# Examples of go-lark (v2)

> Notice: The `appID`, `appSecret`, `email` and other identities are not open to public. To make examples work, you should set them on your own.

## Notification Bot

- [Notification](/notification)

## Chat Bot

- [Authentication](/auth)
- [Messaging](/message)
    - [Basic Message](/message/main.go): Text and some shortcuts that wrapped with Message Buffer.
    - [Message Buffer](/message/message_buffer.go): @at, Image, Share Chat, Rich Text (a.k.a Post).
    - [Interactive Message](/message/card.go)

### Chat

- [Chat CRUD](/chat): Create, Add, Remove, and List Chat.

### Event Handling

- [Gin Middleware](/gin-middleware)
- [Hertz Middleware](/hertz-middleware)
- [Event Forwarding](/event-forward)
