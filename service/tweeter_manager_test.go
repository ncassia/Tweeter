package service

import "testing"

func TestPublishedTweetIsSaved(t *testing.T) {
	var tweet string = "This is my first tweet"

	PublishTweet(tweet)

	if GetTweet() != tweet {
		t.Error("Expected tweet is", tweet)
	}

}
