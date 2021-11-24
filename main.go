package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"card-deck-api/api"
)

func main() {
	portPtr := flag.Int("port", 13370, "specifies the port under which the API will be available")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	router := api.RegisterGinRouter()

	address := fmt.Sprintf(":%d", *portPtr)
	router.Run(address)
}
