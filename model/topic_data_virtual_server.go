package model

import (
	"encoding/json"
)

var (
	_ TopicData = &TopicDataVirtualServer{}
)

type TopicDataVirtualServer struct {
	VirtualServer *VirtualServer
}

func (td *TopicDataVirtualServer) Deserialize(d string) error {
	return json.Unmarshal([]byte(d), td)
}

func (td *TopicDataVirtualServer) Serialize() (string, error) {
	data, err := json.Marshal(td)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (td *TopicDataVirtualServer) Valid() bool {
	return false
}
