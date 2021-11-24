package main

import (
	"card-deck-api/api"
	"flag"
	"fmt"
)

func main() {
	portPtr := flag.Int("port", 13370, "specifies the port under which the API will be available")
	flag.Parse()

	router := api.RegisterGinRouter()

	address := fmt.Sprintf(":%d", *portPtr)
	router.Run(address)
}
