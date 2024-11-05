package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	con net.Conn

	// если мы устанвливаем и забираем соединение - истина
	// если принмаем и забираем соединение - ложь
	outbound bool
}

func NewTCPPeer(con net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		con:      con,
		outbound: outbound,
	}
}

type TCPtransport struct {
	listenAddress string
	listener      net.Listener
	mu            sync.RWMutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPtransport {
	return &TCPtransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPtransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()

	return nil
}

func (t *TCPtransport) startAcceptLoop() error {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		go t.handleConn(conn)
	}
}

func (t *TCPtransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	fmt.Printf("new incoming connection from %+v\n", peer)
}
