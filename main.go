package main

import (
	"ForeverStore/p2p"
	"fmt"
	"log"
)

func OnPeer(peer p2p.Peer) error {
	//fmt.Println("doing some logic with the peer outside of TCPTransport")
	peer.Close()
	return nil
}

func main() {

	tcpOptc := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOptc)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v \n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
