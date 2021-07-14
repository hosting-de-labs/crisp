package model

import "github.com/hosting-de-labs/go-crisp/utils"

type VLANStatus int

const (
	VLANStatusUnknown = iota
	VLANStatusActive
	VLANStatusReserved
	VLANStatusDeprecated
)

type VLAN struct {
	ID          uint16     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Status      VLANStatus `json:"status,omitempty"`
	Description string     `json:"description,omitempty"`
	Tags        []string   `json:"tags,omitempty"`
}

func (v VLAN) Clone() (out VLAN) {
	out = VLAN{
		ID:          v.ID,
		Name:        v.Name,
		Status:      v.Status,
		Description: v.Description,
	}

	if len(v.Tags) > 0 {
		out.Tags = make([]string, len(v.Tags))
		copy(out.Tags, v.Tags)
	}

	return out
}

//IsEqual compares the current IPAddress object against another IPAddress object
func (v VLAN) IsEqual(v2 VLAN) bool {
	return utils.CompareStruct(v, v2, []string{}, []string{})
}
