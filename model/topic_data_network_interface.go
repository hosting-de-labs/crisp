package model

import (
	"net"
)

var (
	_ TopicData = TopicDataNetworkInterface{}
)

type TopicDataNetworkInterface struct {
	Name     string
	Type     string
	Enabled  bool
	MAC      net.HardwareAddr
	IPs      []string
	Children []string
}

func (td TopicDataNetworkInterface) Valid() bool {
	return td.Name != "" &&
		td.Type != "" &&
		td.MAC.String() != ""
}
