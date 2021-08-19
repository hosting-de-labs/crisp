package model

import (
	"encoding/json"
	"time"
)

const (
	TopicAll              string = "*"
	TopicOsLinux          string = "os-linux"
	TopicInventory        string = "inventory"
	TopicNetworkInterface string = "network-interface"

	TopicDataValueUnknown string = "unknown"
)

type TopicData interface {
	Deserialize(string) error
	Serialize() (string, error)
	Valid() bool
}

type Topic struct {
	Name     string        `json:"name,omitempty"`
	Metadata TopicMetadata `json:"metadata"`
	Data     TopicData     `json:"data"`
}

type TopicMetadata struct {
	AddedOn time.Time `json:"added_on"`
}

func deserialize(d string, td interface{}) error {
	return json.Unmarshal([]byte(d), td)
}

func serialize(td interface{}) (string, error) {
	data, err := json.Marshal(td)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
