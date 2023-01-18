package main

import (
	"fmt"

	productsvc "github.com/kyzykyky/softwarearch/svcreg/productsvc/server"
	stocksvc "github.com/kyzykyky/softwarearch/svcreg/stocksvc/server"
)

func main() {
	productserver := productsvc.Server{
		ServiceId: "productsvc",
		Host:      "localhost",
		Port:      8000,
	}
	go productserver.Start()

	stockserver1 := stocksvc.Server{
		ServiceId: "stocksvc1",
		Host:      "localhost",
		Port:      8001,
	}
	go stockserver1.Start()
	stockserver2 := stocksvc.Server{
		ServiceId: "stocksvc2",
		Host:      "localhost",
		Port:      8002,
	}
	go stockserver2.Start()

	// input for keeping the program running
	var input string
	_, _ = fmt.Scanln(&input)
}
