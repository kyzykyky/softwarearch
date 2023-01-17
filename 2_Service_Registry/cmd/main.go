package main

import (
	"fmt"

	productsvc "github.com/kyzykyky/softwarearch/svcreg/productsvc/server"
	stocksvc "github.com/kyzykyky/softwarearch/svcreg/stocksvc/server"
)

func main() {
	productserver := productsvc.Server{
		Host: "localhost",
		Port: 8000,
	}
	go productserver.Start()

	stockserver := stocksvc.Server{
		Host: "localhost",
		Port: 8001,
	}
	go stockserver.Start()

	// input for keeping the program running
	var input string
	_, _ = fmt.Scanln(&input)
}
