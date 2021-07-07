package model

import (
	"sort"

	"github.com/hosting-de-labs/go-crisp/utils"
)

//Host represents a host
type Host struct {
	Platform string

	Hostname  string
	IsManaged bool
	Tags      []string
	Comments  []string
}

//NewHost returns a new instance of an Host with Metadata initialized
func NewHost() *Host {
	return &Host{}
}

//HasTag checks for a specific tag being assigned to this host
func (h *Host) HasTag(tag string) bool {
	for _, existingTag := range h.Tags {
		if existingTag == tag {
			return true
		}
	}

	return false
}

//AddTag is a helper method to allow adding a number of tags to a host.
func (h *Host) AddTag(tags ...string) {
	for _, newTag := range tags {
		tagFound := false
		for _, existingTag := range h.Tags {
			if existingTag == newTag {
				tagFound = true
				break
			}
		}

		if !tagFound {
			h.Tags = append(h.Tags, newTag)
		}
	}
}

//Copy creates a deep copy of the given host
func (h Host) Copy() Host {
	out := NewHost()
	out.Hostname = h.Hostname
	out.Platform = h.Platform
	out.IsManaged = h.IsManaged

	//copy comments
	if len(h.Comments) > 0 {
		out.Comments = make([]string, len(h.Comments))
		copy(out.Comments, h.Comments)
	}

	//copy tags
	if len(h.Tags) > 0 {
		out.Tags = make([]string, len(h.Tags))
		copy(out.Tags, h.Tags)
	}

	return *out
}

//IsEqual compares the current object against another Host object
func (h Host) IsEqual(h2 Host) bool {

	if !utils.CompareStruct(h, h2, []string{}, []string{"Meta", "NetworkInterfaces", "Tags"}) {
		return false
	}

	//tags
	if len(h.Tags) != len(h2.Tags) {
		return false
	}

	sort.Strings(h.Tags)
	sort.Strings(h2.Tags)

	for i := 0; i < len(h.Tags); i++ {
		if h.Tags[i] != h2.Tags[i] {
			return false
		}
	}

	return true
}
