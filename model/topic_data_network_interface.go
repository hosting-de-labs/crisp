package model

import (
	"net"
)

var (
	_ TopicData = TopicDataNetworkInterface{}
)

type NetworkInterface struct {
	Name     string           `json:"name"`
	Type     string           `json:"type"`
	Enabled  bool             `json:"enabled"`
	MAC      net.HardwareAddr `json:"mac"`
	IPs      []string         `json:"ips,omitempty"`
	Children []string         `json:"children,omitempty"`
}

type TopicDataNetworkInterface struct {
	NetworkInterface
}

func (td TopicDataNetworkInterface) Valid() bool {
	return td.Name != "" &&
		td.Type != "" &&
		td.MAC.String() != ""
}
