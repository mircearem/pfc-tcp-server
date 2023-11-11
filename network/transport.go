package network

import (
	"io"
	"log"
	"net"

	"github.com/sirupsen/logrus"
)

// A peer in the network
type Peer struct {
	listenAddr string
	conn       net.Conn
}

func NewPeer(c net.Conn) *Peer {
	return &Peer{
		listenAddr: c.RemoteAddr().String(),
		conn:       c,
	}
}

func (p *Peer) Write(b []byte) error {
	_, err := p.conn.Write(b)
	return err
}

func (p *Peer) ReadLoop(msgch chan *Message, delPeer chan *Peer) {
	log.Println("Starting read loop")
	defer p.conn.Close()
	buf := make([]byte, 2048)

	for {
		n, err := p.conn.Read(buf)
		if err == io.EOF {
			logrus.Error(err)
			break
		}

		log.Println(buf[:n])
		msg := NewMessage(p, buf[:n])
		msgch <- msg
	}

	delPeer <- p
}

type Transport struct {
	listenAddr string
	ln         net.Listener
	AddPeer    chan *Peer
	DelPeer    chan *Peer
}

func NewTransport(addr string) *Transport {
	return &Transport{
		listenAddr: addr,
	}
}

func (t *Transport) ListenAndAccept() error {
	ln, err := net.Listen("tcp", t.listenAddr)
	if err != nil {
		return err
	}
	t.ln = ln
	defer t.ln.Close()

	// Accept incomming connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			logrus.Error(err)
			continue
		}

		peer := NewPeer(conn)
		t.AddPeer <- peer
	}
}
