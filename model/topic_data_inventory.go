package model

var (
	_ TopicData = TopicDataInventory{}
)

type InventoryItem struct {
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	PartNumber   string `json:"part_number,omitempty"`
	SerialNumber string `json:"serial_number,omitempty"`
}

type TopicDataInventory struct {
	Items []InventoryItem
}

func (td TopicDataInventory) Valid() bool {
	if len(td.Items) == 0 {
		return false
	}

	for _, item := range td.Items {
		if item.Type == "" ||
			item.Manufacturer == "" ||
			item.Model == "" {
			return false
		}
	}

	return true
}
