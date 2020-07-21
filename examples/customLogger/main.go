package main

import (
	"time"

	"github.com/syrinsecurity/gologger"
)

func main() {
	log := gologger.NewCustomLogger("./logs-", ".txt", 200)
	go log.Service()

	log.Write("test")

	time.Sleep(time.Second * 1)
	log.Close()
}
