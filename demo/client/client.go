package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/junzexu/nuts/demo/gen-go/multiple"
	nuts "github.com/junzexu/nuts/framework"

	"git.apache.org/thrift.git/lib/go/thrift"
)

var (
	multipleClient *multiple.MultiplicationServiceClient
)

func init() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	transport := thrift.NewTSocketFromAddrTimeout(addr, time.Millisecond*50)
	transport.Open()
	protocol := nuts.NewNutProtocolFactory(thrift.NewTBinaryProtocolFactoryDefault()).GetProtocol(transport)
	client := thrift.NewTStandardClient(protocol, protocol)
	multipleClient = multiple.NewMultiplicationServiceClient(client)
}

func main() {
	for i := 0; i < 3; i++ {
		a, b := multiple.Int(rand.Int()), multiple.Int(rand.Int())
		rpcResult, err := multipleClient.Multiply(nil, a, b)
		fmt.Printf("rpc result: %+v , %+v => %+v, %+v\n", a, b, rpcResult, err)
		time.Sleep(time.Millisecond)
	}
}
