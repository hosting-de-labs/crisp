package model

var (
	_ TopicData = &TopicDataInventory{}
)

// TopicDataInventory is a container for inventory items being transmitted as topic data.
type TopicDataInventory struct {
	Items []InventoryItem `json:"items,omitempty"`
}

// Deserialize calls internal helper to read data from json string.
func (td *TopicDataInventory) Deserialize(d string) error {
	return deserialize(d, td)
}

// Serialize returns a json string or reports an error.
func (td *TopicDataInventory) Serialize() (string, error) {
	return serialize(td)
}

// Valid verifies that every inventory item have required fields set.
func (td *TopicDataInventory) Valid() bool {
	if len(td.Items) == 0 {
		return false
	}

	for _, item := range td.Items {
		if item.Type == 0 ||
			item.Manufacturer == "" ||
			item.Model == "" {
			return false
		}
	}

	return true
}
