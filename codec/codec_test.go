package codec_test

import (
	"io"
	"testing"

	"github.com/nguyencatpham/go-micro/codec"
	"github.com/nguyencatpham/go-micro/codec/bytes"
	"github.com/nguyencatpham/go-micro/codec/grpc"
	"github.com/nguyencatpham/go-micro/codec/json"
	"github.com/nguyencatpham/go-micro/codec/jsonrpc"
	"github.com/nguyencatpham/go-micro/codec/proto"
	"github.com/nguyencatpham/go-micro/codec/protorpc"
	"github.com/nguyencatpham/go-micro/codec/text"
)

type testRWC struct{}

func (rwc *testRWC) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (rwc *testRWC) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (rwc *testRWC) Close() error {
	return nil
}

func getCodecs(c io.ReadWriteCloser) map[string]codec.Codec {
	return map[string]codec.Codec{
		"bytes":    bytes.NewCodec(c),
		"grpc":     grpc.NewCodec(c),
		"json":     json.NewCodec(c),
		"jsonrpc":  jsonrpc.NewCodec(c),
		"proto":    proto.NewCodec(c),
		"protorpc": protorpc.NewCodec(c),
		"text":     text.NewCodec(c),
	}
}

func Test_WriteEmptyBody(t *testing.T) {
	for name, c := range getCodecs(&testRWC{}) {
		err := c.Write(&codec.Message{
			Type:   codec.Error,
			Header: map[string]string{},
		}, nil)
		if err != nil {
			t.Fatalf("codec %s - expected no error when writing empty/nil body: %s", name, err)
		}
	}
}
