package main

import (
	"fmt"
	"log"
	"os"

	"github.com/artiga033/a_play_list"
	"github.com/artiga033/a_play_list/alist"
	"github.com/artiga033/a_play_list/playlist"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, skipping..")
	}
	var conf a_play_list.Config
	a_play_list.LoadEnvConf(&conf)
	writer, err := playlist.NewM3U8Writer(os.Stdout)
	if err != nil {
		log.Fatal("writer not available")
	}
	gen := a_play_list.NewGenerator(alist.NewClient(
		alist.MakeClientOptions{
			Endpoint: conf.Endpoint,
			User:     conf.User,
			Pass:     conf.Pass}),
		writer)
	var path, password string
	if len(os.Args) >= 3 {
		path = os.Args[1]
		password = os.Args[2]
	} else if len(os.Args) >= 2 {
		path = os.Args[1]
	} else {
		fmt.Println("usage:\na_play_list <path> [password]\nenv vars:\n\tA_PLAY_LIST_ENDPOINT\n\tA_PLAY_LIST_USER\n\tA_PLAY_LIST_PASS")
		return
	}
	gen.GenerateFor(path, password)
}
