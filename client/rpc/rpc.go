package rpc

import (
	"github.com/elthariel/go-micro/client"
)

func NewClient(opts ...client.Option) client.Client {
	return client.NewClient(opts...)
}
