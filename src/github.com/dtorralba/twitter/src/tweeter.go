package main

import (
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/dtorralba/twitter/src/domain"
	"github.com/dtorralba/twitter/src/service"
)

func main() {
	var tweet *domain.Tweet
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your user: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			tweet = domain.NewTweet(user, text)

			service.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showAllTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweets()

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "findByID",
		Help: "Find a Tweet by ID",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your ID: ")

			id := c.ReadLine()
			idNew, _ := strconv.Atoi(id)

			tweet, _ := service.GetTweetById(idNew)

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countTweetsByUser",
		Help: "Count tweets made by a user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write user: ")

			user := c.ReadLine()

			tweet := service.CountTweetsByUser(user)

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "TweetsByUser",
		Help: "Retrieve The Tweets Sent By An User",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write user: ")

			user := c.ReadLine()

			tweet := service.GetTweetsByUser(user)

			c.Println(tweet)

			return
		},
	})

	shell.Run()

}
