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

type TopicDataOsLinux struct {
	OsLinux
}

func (td *TopicDataOsLinux) Deserialize(d string) error {
	return deserialize(d, td)
}

func (td *TopicDataOsLinux) Serialize() (string, error) {
	return serialize(td)
}

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
