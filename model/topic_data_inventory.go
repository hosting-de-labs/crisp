package model

var (
	_ TopicData = TopicDataInventory{}
)

type TopicDataInventory struct {
	Items []InventoryItem
}

func (td TopicDataInventory) Valid() bool {
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
