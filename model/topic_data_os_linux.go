package model

import (
	"fmt"
)

var (
	_ TopicData = &TopicDataOsLinux{}
)

type OsLinux struct {
	Arch              string `json:"arch,omitempty"`
	OsName            string `json:"os_name,omitempty"`
	OsVersionId       string `json:"os_version_id,omitempty"`
	OsVersionCodename string `json:"os_version_codename,omitempty"`
	Version           string `json:"version,omitempty"`
	Virtualization    string `json:"virtualization,omitempty"`
}

// TopicDataOsLinux is a container for OsLinux data being transmitted as topic data.
type TopicDataOsLinux struct {
	OsLinux
}

// Deserialize calls internal helper to read data from json string.
func (td *TopicDataOsLinux) Deserialize(d string) error {
	return deserialize(d, td)
}

// Serialize returns a json string or reports an error.
func (td *TopicDataOsLinux) Serialize() (string, error) {
	return serialize(td)
}

// Valid verifies required os linux data is set.
func (td TopicDataOsLinux) Valid() bool {
	return td.Arch != "" &&
		td.OsName != "" &&
		td.OsVersionId != "" &&
		td.OsVersionCodename != "" &&
		td.Version != ""
}

func (td TopicDataOsLinux) String() string {
	return fmt.Sprintf("%s %s (%s, %s)", td.OsName, td.OsVersionId, td.OsVersionCodename, td.Arch)
}
