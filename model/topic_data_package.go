package model

var (
	_ TopicData = &TopicDataPackage{}
)

type TopicDataPackage struct {
	Packages []Package
}

// Deserialize calls internal helper to read data from json string.
func (td *TopicDataPackage) Deserialize(s string) error {
	return deserialize(s, td)
}

// Serialize returns a json string or reports an error.
func (td *TopicDataPackage) Serialize() (string, error) {
	return serialize(td)
}

// Valid verifies that every package have required fields set.
func (td *TopicDataPackage) Valid() bool {
	if len(td.Packages) == 0 {
		return false
	}

	for _, p := range td.Packages {
		if p.Name == "" ||
			p.Version == "" ||
			p.State == 0 {
			return false
		}
	}

	return true
}
