package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *RPC) error
}

type GOBDecoder struct{}

func (dec GOBDecoder) Decode(r io.Reader, msg *RPC) error {
	return gob.NewDecoder(r).Decode(msg)
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(r io.Reader, msg *RPC) error {
	peekBuf := make([]byte, 1)

	if _, err := r.Read(peekBuf); err != nil {
		return nil
	}

	// In case of a stream we are not decoding what is being sent over net
	// just setting stream true se we can handle that in our logick
	stream := peekBuf[0] == IncommingStream

	if stream {
		msg.Stream = true // this is go so 1 and true is not same
		return nil
	}

	buf := make([]byte, 1028) /// ?
	n, err := r.Read(buf)
	if err != nil {
		return err
	}

	msg.Payload = buf[:n]

	return nil
}
