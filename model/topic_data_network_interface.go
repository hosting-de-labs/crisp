package model

import "encoding/json"

var (
	_ TopicData = &TopicDataNetworkInterface{}
)

type TopicDataNetworkInterface struct {
	NetworkInterface
}

func (td *TopicDataNetworkInterface) Deserialize(d string) error {
	return json.Unmarshal([]byte(d), td)
}

func (td *TopicDataNetworkInterface) Serialize() (string, error) {
	data, err := json.Marshal(td)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (td TopicDataNetworkInterface) Valid() bool {
	return td.Name != "" &&
		td.Type != "" &&
		td.MACAddress.String() != ""
}
