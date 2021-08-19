package model

var (
	_ TopicData = &TopicDataNetworkInterface{}
)

type TopicDataNetworkInterface struct {
	Interfaces []NetworkInterface `json:"interfaces,omitempty"`
}

func (td *TopicDataNetworkInterface) Deserialize(d string) error {
	return deserialize(d, td)
}

func (td *TopicDataNetworkInterface) Serialize() (string, error) {
	return serialize(td)
}

func (td TopicDataNetworkInterface) Valid() bool {
	if len(td.Interfaces) == 0 {
		return false
	}

	for _, intf := range td.Interfaces {
		if intf.Name == "" ||
			intf.Type == "" ||
			intf.MACAddress.String() == "" {
			return false
		}
	}

	return true
}
