package main

// curl -d "from=2022-10-06&to=2022-10-13" -X POST http://129.200.0.18/mantis/mantisreport.php

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/text/encoding/charmap"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DecodeWindows1251(ba []uint8) []uint8 {
	dec := charmap.Windows1251.NewDecoder()
	out, _ := dec.Bytes(ba)
	return out
}

func getHtml() string {
	data := url.Values{
		"from":   {"2022-10-06"},
		"to":     {"2022-10-13"},
		"report": {"%CF%EE%EA%E0%E7%E0%F2%FC"},
	}

	// resp, err := http.PostForm("http://s122-webdev.solar.com/mantis/mantisreport.php", data)
	resp, err := http.PostForm("http://129.200.0.18/mantis/mantisreport.php", data)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	html := DecodeWindows1251(body)

	return string(html)
}

func main() {
	bot, err := tgbotapi.NewBotAPI("5788367491:AAHSc2ARG3m4QwNnAWKdJd9FfcMg51beFwY")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 60

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message == nil {
			continue
		}

		var text = getHtml()
		var msgText = parseTable(text)

		// Now that we know we've gotten a new message, we can construct a
		// reply! We'll take the Chat ID and Text from the incoming message
		// and use it to create a new message.
		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)

		// We'll also say that this message is a reply to the previous message.
		// For any other specifications than Chat ID or Text, you'll need to
		// set fields on the `MessageConfig`.
		msg.ReplyToMessageID = update.Message.MessageID

		// Okay, we're sending our message off! We don't care about the message
		// we just sent, so we'll discard it.
		if _, err := bot.Send(msg); err != nil {
			// Note that panics are a bad way to handle errors. Telegram can
			// have service outages or network errors, you should retry sending
			// messages or more gracefully handle failures.
			panic(err)
		}
	}

}
