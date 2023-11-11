package network

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"sync"
)

type Server struct {
	transport *Transport
	mu        sync.RWMutex
	peers     map[net.Addr]*Peer
	addPeerCh chan *Peer
	delPeerCh chan *Peer
	msgCh     chan *Message
}

func NewServer(addr string) *Server {
	s := &Server{
		peers:     make(map[net.Addr]*Peer),
		addPeerCh: make(chan *Peer, 10),
		delPeerCh: make(chan *Peer),
		msgCh:     make(chan *Message, 1024),
	}

	tr := NewTransport(addr)

	s.transport = tr
	tr.AddPeer = s.addPeerCh
	tr.DelPeer = s.delPeerCh

	return s
}

func (s *Server) Start() {

}

func (s *Server) addPeer(p *Peer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.peers[p.conn.RemoteAddr()] = p
}

func (s *Server) delPeer(p *Peer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p.conn.Close()
	delete(s.peers, p.conn.RemoteAddr())
}

func (s *Server) handshake(p *Peer) error {
	reqmsg := &MessageHandshakeRequest{
		ProtocolVersion:      PROTOCOL_VERSION,
		HandshakeRequestByte: HANDSHAKE_REQUEST_BYTE,
	}
	reqbuf := new(bytes.Buffer)
	binary.Write(reqbuf, binary.LittleEndian, reqmsg)

	// Send the handshake request
	if _, err := p.conn.Write(reqbuf.Bytes()); err != nil {
		return err
	}

	// Handle the handshake response
	buf := make([]byte, 1024)
	n, err := p.conn.Read(buf)
	if err != nil {
		return err
	}

	respmsg := &MessageHandshakeResponse{}

	respbuf := bytes.NewBuffer(buf[:n])
	if err := binary.Read(respbuf, binary.LittleEndian, respbuf); err != nil {
		return err
	}

	// Check for matches for protocol version and response byte
	if respmsg.ProtocolVersion != PROTOCOL_VERSION {
		errStr := fmt.Sprintf("handshake error, expected protocol version: %v", PROTOCOL_VERSION)
		return errors.New(errStr)
	}

	if respmsg.HandshakeResponseByte != HANDSHAKE_REQUEST_BYTE {
		return errors.New("handshake error, invalid request byte")
	}

	return nil
}
