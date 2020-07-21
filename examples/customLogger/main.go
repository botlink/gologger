package main

import (
	"github.com/syrinsecurity/gologger"
)

func main() {
	log := gologger.NewCustomLogger("./logs-", ".txt", 0)

	//This will make the filename update every month for example: logs-Jul-2020.txt
	gologger.SetNameConventionMonthYear(log)
	go log.Service()
	defer log.Close()

	log.Write("test")

}
