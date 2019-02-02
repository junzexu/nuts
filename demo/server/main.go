package main

import (
	"net"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/junzexu/nuts/demo/gen-go/multiple"
	"github.com/junzexu/nuts/framework"
	"github.com/junzexu/nuts/logging"
)

var logger = logging.GetLogger("nuts")

func main() {
	processor := multiple.NewMultiplicationServiceProcessor(&MultiplicationServiceHandler{})
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	logger.Info("server runing on: %+v", addr)
	timeout := time.Millisecond * 50
	server := nuts.NewNutServer(thrift.NewTProcessorFactory(processor), addr, timeout)
	server.Serve()
}
