package domain

import (
	"time"
)

type Stringer interface {
	String() string
}

type Tweet interface {
	GetUser() string
	GetText() string
	GetDate() *time.Time
}

type TextTweet struct {
	User string
	Text string
	Date *time.Time
}

func NewTweet(user string, text string) TextTweet {
	var time = time.Now()
	textTweet := TextTweet{user, text, &time}
	return textTweet
}

func (textTweet *TextTweet) PrintableTweet() string {
	return "@" + textTweet.User + ": " + textTweet.Text
}

func (textTweet *TextTweet) String() string {
	return "@" + textTweet.User + ": " + textTweet.Text
}

func (tweet TextTweet) GetUser() string {
	return tweet.User
}

func (tweet TextTweet) GetText() string {
	return tweet.Text
}

func (tweet TextTweet) GetDate() *time.Time {
	return tweet.Date
}
