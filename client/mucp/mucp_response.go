package mucp

import (
	"github.com/nguyencatpham/go-micro/v3/codec"
	"github.com/nguyencatpham/go-micro/v3/network/transport"
)

type rpcResponse struct {
	header map[string]string
	body   []byte
	socket transport.Socket
	codec  codec.Codec
}

func (r *rpcResponse) Codec() codec.Reader {
	return r.codec
}

func (r *rpcResponse) Header() map[string]string {
	return r.header
}

func (r *rpcResponse) Read() ([]byte, error) {
	var msg transport.Message

	if err := r.socket.Recv(&msg); err != nil {
		return nil, err
	}

	// set internals
	r.header = msg.Header
	r.body = msg.Body

	return msg.Body, nil
}
