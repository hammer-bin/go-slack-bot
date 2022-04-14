package main

import (
	"fmt"
	"go-slack-bot/app"
)

func main() {
	fmt.Println("Start process!")
	app.Lottery()
	fmt.Println("End process!")
}
