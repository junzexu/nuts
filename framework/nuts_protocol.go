package nuts

import (
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// NutProtocol hook
type NutProtocol struct {
	beginTime *time.Time
	thrift.TProtocol
}

// WriteMessageBegin .
// log message call
func (p *NutProtocol) WriteMessageBegin(name string, typeID thrift.TMessageType, seqid int32) error {
	writeBegin := time.Now()
	err := p.TProtocol.WriteMessageBegin(name, typeID, seqid)
	writeEnd := time.Now()

	if p.beginTime == nil {
		writeCost := writeEnd.Sub(writeBegin).Nanoseconds()
		logger.Info("client nut rpc begin method: %s, messageType: %d, seqid: %d, write cost: %d nanoseconds", name, typeID, seqid, writeCost)
		p.beginTime = &writeBegin
	} else {
		cost := writeEnd.Sub(*p.beginTime).Nanoseconds()
		logger.Info("server rpc end method: %s, messageType: %d, seqid: %d, cost=%d nanoseconds", name, typeID, seqid, cost)
		p.beginTime = nil
	}
	return err
}

// ReadMessageBegin .
// log message called
func (p *NutProtocol) ReadMessageBegin() (name string, typeId thrift.TMessageType, seqid int32, err error) {
	readBegin := time.Now()
	name, typeId, seqid, err = p.TProtocol.ReadMessageBegin()
	readEnd := time.Now()

	if p.beginTime == nil {
		// rpc call begin
		readCost := readEnd.Sub(readBegin).Nanoseconds()
		logger.Info("server nut rpc begin method: %s, messageType: %d, seqid: %d, readCost: %d nanoseconds", name, typeId, seqid, readCost)
		p.beginTime = &readBegin
	} else {
		cost := readEnd.Sub(*p.beginTime).Nanoseconds()
		logger.Info("client rpc end method: %s, messageType: %d, seqid: %d, cost=%d nanoseconds", name, typeId, seqid, cost)
		p.beginTime = nil
	}
	return
}

// NutProtocolFactory .
type NutProtocolFactory struct {
	thrift.TProtocolFactory
}

// GetProtocol .
func (f *NutProtocolFactory) GetProtocol(trans thrift.TTransport) thrift.TProtocol {
	logger.Info("nut protocol Factory GetProtocol")
	protocol := f.TProtocolFactory.GetProtocol(trans)
	return &NutProtocol{TProtocol: protocol}
}

// NewNutProtocolFactory .
func NewNutProtocolFactory(lowLayer thrift.TProtocolFactory) *NutProtocolFactory {
	return &NutProtocolFactory{lowLayer}
}
