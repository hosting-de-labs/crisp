package model

var (
	_ TopicData = &TopicDataNetworkInterface{}
)

// TopicDataNetworkInterface is a container for network interfaces being transmitted as topic data.
type TopicDataNetworkInterface struct {
	Interfaces []NetworkInterface `json:"interfaces,omitempty"`
}

// Deserialize calls internal helper to read data from json string.
func (td *TopicDataNetworkInterface) Deserialize(d string) error {
	return deserialize(d, td)
}

// Serialize returns a json string or reports an error.
func (td *TopicDataNetworkInterface) Serialize() (string, error) {
	return serialize(td)
}

// Valid verifies that every interface have required fields set.
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
