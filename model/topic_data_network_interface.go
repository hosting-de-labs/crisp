package model

var (
	_ TopicData = TopicDataNetworkInterface{}
)

type TopicDataNetworkInterface struct {
	NetworkInterface
}

func (td TopicDataNetworkInterface) Valid() bool {
	return td.Name != "" &&
		td.Type != "" &&
		td.MACAddress.String() != ""
}
