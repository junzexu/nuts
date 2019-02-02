package nuts

import (
	"context"
	"net"
)

// DialFunc can use as net.Resolver.Dial member to change the ddefault action.
// to implement this just like a middleware, must wrap the default action when
// use this to change net.DefaultResolver
type DialFunc func(ctx context.Context, network, address string) (net.Conn, error)

// DialLogger .
func DialLogger(ctx context.Context, network, address string) (net.Conn, error) {
	logger.Info("dialing %s:%s", network, address)

	// net.DefaultResolver

	return nil, nil
}

// NutDialer .
type NutDialer struct {
	net.Dialer
	net.Resolver
	net.Conn
	net.Listener
}
