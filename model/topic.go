package model

import (
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
