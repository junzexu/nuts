package nuts

// impletement thrift.TClient

import (
	"context"
	"fmt"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// NewNutClient .
func NewNutClient(conf string) *NutClient {
	client := &NutClient{}

	cliTransport := NewNutTransport()
	itranFactory, _ := client.itranFactory.GetTransport(cliTransport)
	otranFactory, _ := client.otranFactory.GetTransport(cliTransport)
	iprotFactory := client.iprotFactory.GetProtocol(itranFactory)
	oprotFactory := client.oprotFactory.GetProtocol(otranFactory)

	client.TClient = thrift.NewTStandardClient(iprotFactory, oprotFactory)

	return client
}

// NutClient .
type NutClient struct {
	thrift.TClient

	cliTransport thrift.TTransport
	itranFactory thrift.TTransportFactory
	otranFactory thrift.TTransportFactory
	iprotFactory thrift.TProtocolFactory
	oprotFactory thrift.TProtocolFactory
}

// Call .
func (cli *NutClient) Call(ctx context.Context, method string, args, result thrift.TStruct) error {
	begin := time.Now()
	err := cli.TClient.Call(ctx, method, args, result)
	cost := time.Now().Sub(begin).Nanoseconds()

	// just like a middleware
	fmt.Printf("call %s cost %d nanoseconds", method, cost)

	return err
}
