package main

import (
	"fmt"

	"github.com/TafveezA/BlockchainNetwork/network"
)

func main() {
	trLocal := network.NewLocalTransport("Local")

opts := network.NewServer(opts)(
	Transports: []network.Transport{trLocal}
)
}