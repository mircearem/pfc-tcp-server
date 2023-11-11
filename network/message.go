package network

const (
	PROTOCOL_VERSION        = 1
	HANDSHAKE_REQUEST_BYTE  = 5
	HANDSHAKE_RESPONSE_BYTE = 6
)

// MessageHandshakeRequest contains the data sent to a
// newly connected Peer to establish a handshake
type MessageHandshakeRequest struct {
	ProtocolVersion      uint16
	HandshakeRequestByte uint8
}

// MessageHandshakeResponse contains the data sent to
// the server by the newly connected Peer as a response
// to the MessageHandshakeRequest
type MessageHandshakeResponse struct {
	ProtocolVersion       uint16
	HandshakeResponseByte uint8
}

type Message struct {
	From    *Peer
	Payload any
}

// Decode the message body
func NewMessage(from *Peer, payload any) *Message {
	return &Message{
		From:    from,
		Payload: payload,
	}
}
