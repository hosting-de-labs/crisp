package model

var (
	_ TopicData = TopicDataInventory{}
)

type TopicDataInventory struct {
	Items []struct {
		Type         string
		Manufacturer string
		Model        string
		PartNumber   string
		SerialNumber string
	} `json:"items"`
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
