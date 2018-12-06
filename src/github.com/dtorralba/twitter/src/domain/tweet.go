package domain

import (
	"time"
)

type Stringer interface {
	String() string
}

type Tweet struct {
	User string
	Text string
	Date *time.Time
}

func NewTweet(user string, text string) *Tweet {
	var time = time.Now()
	tweet := Tweet{user, text, &time}
	return &tweet
}

func (tweet *Tweet) PrintableTweet() string {
	return "@" + tweet.User + ": " + tweet.Text
}

func (tweet *Tweet) String() string {
	return "@" + tweet.User + ": " + tweet.Text
}
