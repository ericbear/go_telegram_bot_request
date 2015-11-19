### How to install ###

```
go get github.com/ericbear/go_telegram_bot_request
```

### Request Protocol for Telegram Bot ###

- Import

```
import (
	"github.com/ericbear/go_telegram_bot_request/message"
)
```

- Get Update - [api](https://core.telegram.org/bots/api#getupdates)

```
	token := "xxxxx.xxxxxxxxx"
	lastUpdateId := 0
	msg := message.GetUpdates("https://api.telegram.org/bot"+token+"/", lastUpdateId)
```

- Send Message - [api](https://core.telegram.org/bots/api#sendmessage)

```
	token := "xxxxx.xxxxxxxx"
	chatId := 1
	replyChatId := 0
	msg := message.SendMessage("https://api.telegram.org/bot"+token+"/", chatId, "hello world", replyChatId)
```

### How to create Bot ###

- [Telegram Bots](https://core.telegram.org/bots)
> add "BotFather" on telegram to create Bot
