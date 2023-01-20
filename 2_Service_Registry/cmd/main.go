package main

import (
	"fmt"

	productsvc "github.com/kyzykyky/softwarearch/svcreg/productsvc/server"
	stocksvc "github.com/kyzykyky/softwarearch/svcreg/stocksvc/server"
)

func main() {
	go productsvc.Server{
		ServiceId: "productsvc",
		Host:      "localhost",
		Port:      8000,
	}.Start()

	go stocksvc.Server{
		ServiceId: "stocksvc1",
		Host:      "localhost",
		Port:      8001,
	}.Start()
	go stocksvc.Server{
		ServiceId: "stocksvc2",
		Host:      "localhost",
		Port:      8002,
	}.Start()

	// input for keeping the program running
	var input string
	_, _ = fmt.Scanln(&input)
}
