package service

import "github.com/ncassia/Tweeter/domain"

var Tweet domain.Tweet

func PublishTweet(tweet *domain.Tweet) {
	Tweet.User = tweet.User
	Tweet.Text = tweet.Text
	Tweet.Date = tweet.Date
}

func GetTweet() domain.Tweet {
	return Tweet
}
