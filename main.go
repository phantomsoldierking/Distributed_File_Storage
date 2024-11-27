package main

import (
	"log"

	"github.com/anthdm/foreverstore/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	log.Fatal(tr.ListenAndAccept())

	select {}

}
