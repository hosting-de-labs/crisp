package model

import "fmt"

type PackageState int

const (
	PackageStateUnknown PackageState = iota
	PackageStateInstalled
	PackageStateRemoved
)

var (
	_ fmt.Stringer = PackageState(0)
)

func (p PackageState) String() string {
	switch p {
	case PackageStateInstalled:
		return "installed"
	case PackageStateRemoved:
		return "removed"
	default:
		return "unknown"
	}
}

var (
	_ fmt.Stringer = &Package{}
)

type Package struct {
	Name    string       `json:"name,omitempty"`
	Version string       `json:"version,omitempty"`
	State   PackageState `json:"state,omitempty"`
}

func (p *Package) String() string {
	return fmt.Sprintf("%s: %s (%s)", p.Name, p.Version, p.State)
}
