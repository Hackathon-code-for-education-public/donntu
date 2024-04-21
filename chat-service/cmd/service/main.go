package main

import "chat-service/internal/app"

func main() {
	a, cleanup, err := app.Init()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	a.Run()
}
