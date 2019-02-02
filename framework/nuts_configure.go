package nuts

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

// TransportConf .
type TransportConf struct {
	Name     string
	Protocal string
	Host     string
	Port     string
}

// TransportFactoryConf .
type TransportFactoryConf struct {
	Name       string
	PoolSize   uint
	BufferSize uint
}

// NutServerConf .
type NutServerConf struct {
	ServerTransports  []TransportConf        `xml:">Transport"`
	TransportFactorys []TransportFactoryConf `xml:">TransportFactory"`
}

// NutClientConf .
type NutClientConf struct {
	Transport         TransportConf
	TransportFactorys []TransportFactoryConf `xml:">TransportFactory"`
}

// NutConf .
type NutConf struct {
	NutServer NutServerConf
	Protocols []string `xml:"Protocols>Protocol"`
	NutClient NutClientConf
}

// NewNutConf .
func NewNutConf(confFilename string) *NutConf {
	// TODO: add readFile error handle
	data, err := ioutil.ReadFile(confFilename)
	if err != nil {
		log.Panicf("read file error: param %+v, err %+v", confFilename, err)
		return nil
	}
	conf := NutConf{}

	// TODO: add unmarshal error handle
	err = xml.Unmarshal(data, &conf)
	if err != nil {
		log.Panicf("unmarshal error: param %+v, err %+v", confFilename, err)
	}

	return &conf
}
