package main

import (
	"flag"
	"fmt"
	fshare "go-fshare"
)

var (
	username  string
	password  string
	fshareUrl string
)

func main() {
	flag.StringVar(&username, "u", "", "username")
	flag.StringVar(&password, "p", "", "password")
	flag.StringVar(&fshareUrl, "url", "", "fshare url")
	flag.Parse()

	client := fshare.NewClient(fshare.NewConfig(username, password))
	if err := client.Login(); err != nil {
		panic(err)
	}
	if url, err := client.GetDownloadURL(fshareUrl); err != nil {
		panic(err)
	} else {
		fmt.Println(url)
	}
}
