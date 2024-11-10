package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":3000"
	tcpOptc := TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}
	tr := NewTCPTransport(tcpOptc)
	assert.Equal(t, tr.ListenAddr, listenAddr)

	//server
	assert.Nil(t, tr.ListenAndAccept())

}
