package nuts

// 负责请求处理循环

import (
	"net"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// NutServer implement thrift.TServer interface
type NutServer struct {
	serverTransport thrift.TServerTransport
	processor       thrift.TProcessorFactory
}

// NewNutServer .
func NewNutServer(processor thrift.TProcessorFactory, addr net.Addr, clientTimeout time.Duration) *NutServer {
	return &NutServer{
		serverTransport: thrift.NewTServerSocketFromAddrTimeout(addr, clientTimeout),
		processor:       processor,
	}
}

// ProcessorFactory .
func (svr *NutServer) ProcessorFactory() thrift.TProcessorFactory {
	return svr.processor
}

// ServerTransport .
func (svr *NutServer) ServerTransport() thrift.TServerTransport {
	return svr.serverTransport
}

// InputTransportFactory .
func (svr *NutServer) InputTransportFactory() thrift.TTransportFactory {
	return thrift.NewTBufferedTransportFactory(4096)
}

// OutputTransportFactory .
func (svr *NutServer) OutputTransportFactory() thrift.TTransportFactory {
	return thrift.NewTBufferedTransportFactory(4096)
}

// InputProtocolFactory .
func (svr *NutServer) InputProtocolFactory() thrift.TProtocolFactory {
	return NewNutProtocolFactory(thrift.NewTBinaryProtocolFactoryDefault())
}

// OutputProtocolFactory .
func (svr *NutServer) OutputProtocolFactory() thrift.TProtocolFactory {
	return NewNutProtocolFactory(thrift.NewTBinaryProtocolFactoryDefault())
}

// Serve Starts the server
func (svr *NutServer) Serve() error {
	itranFactory := svr.InputTransportFactory()
	iprotFactory := svr.InputProtocolFactory()
	processFactory := svr.ProcessorFactory()

	svr.serverTransport.Listen()
	for {
		conn, err := svr.serverTransport.Accept()
		if err == nil {
			go func(transport thrift.TTransport) {
				trans, _ := itranFactory.GetTransport(transport)
				protocol := iprotFactory.GetProtocol(trans)

				processor := processFactory.GetProcessor(nil)

				counter := 0
				for {
					ok, error := processor.Process(nil, protocol, protocol)
					if !ok {
						logger.Error("Proccess get false, exit!")
						break
					}

					if error != nil {
						logger.Error("Process get error: %+v, exit", err)
						break
					}
					counter++
				}
				logger.Info("connection %+v do %d request.", conn, counter)
			}(conn)
		}
	}
}

// Stop the server. This is optional on a per-implementation basis. Not
// all servers are required to be cleanly stoppable.
func (svr *NutServer) Stop() error {
	return svr.serverTransport.Close()
}
