package model

import (
	"time"

	"github.com/google/uuid"
)

type Host struct {
	Hostname   string    `json:"hostname,omitempty"`
	MachineID  uuid.UUID `json:"machine_id,omitempty"`
	Approved   bool      `json:"approved,omitempty"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
	SeenOn     time.Time `json:"seen_on"`

	Topics []Topic `json:"-,omitempty"`
}
