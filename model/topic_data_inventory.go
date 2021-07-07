package model

import "encoding/json"

var (
	_ TopicData = &TopicDataInventory{}
)

type TopicDataInventory struct {
	Items []InventoryItem `json:"items,omitempty"`
}

func (td *TopicDataInventory) Deserialize(d string) error {
	return json.Unmarshal([]byte(d), td)
}

func (td *TopicDataInventory) Serialize() (string, error) {
	data, err := json.Marshal(td)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

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
