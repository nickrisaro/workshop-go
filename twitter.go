package main

import (
	"github.com/nickrisaro/workshop-go/rest"
	"github.com/nickrisaro/workshop-go/tweet"
)

func main() {

	tweetRepository := tweet.NewRepository()
	apiREST := rest.New(tweetRepository)

	tweet := tweet.New("@nickrisaro", "Acá codeando un poco en @golang")
	tweetRepository.Save(tweet)

	apiREST.Start()
}
