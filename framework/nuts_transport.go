package nuts

import (
	"git.apache.org/thrift.git/lib/go/thrift"
)

// NutTransport manage sockets
// just like connection pool
type NutTransport struct {
	thrift.TTransport
}

// NewNutTransport .
func NewNutTransport() *NutTransport {
	return &NutTransport{}
}
