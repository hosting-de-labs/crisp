package model

var (
	_ TopicData = &TopicDataPackage{}
)

type TopicDataPackage struct {
	Packages []Package
}

func (td *TopicDataPackage) Deserialize(s string) error {
	return deserialize(s, td)
}

func (td *TopicDataPackage) Serialize() (string, error) {
	return serialize(td)
}

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
