package model

import "sort"

//DedicatedServer represents a dedicated server
type DedicatedServer struct {
	Host

	AssetTag     string `json:"asset_tag,omitempty"`
	SerialNumber string `json:"serial_number,omitempty"`

	Inventory []InventoryItem `json:"inventory,omitempty"`
}

func NewDedicatedServer() *DedicatedServer {
	return &DedicatedServer{
		Host:      *NewHost(),
		Inventory: []InventoryItem{},
	}
}

//Copy creates a deep copy of the given host
func (d DedicatedServer) Copy() DedicatedServer {
	out := DedicatedServer{}

	//Copy Host
	out.Host = d.Host.Copy()

	out.AssetTag = d.AssetTag

	//Copy Inventory
	if len(d.Inventory) > 0 {
		out.Inventory = make([]InventoryItem, 0, len(d.Inventory))
		for _, item := range d.Inventory {
			newItem := item.Copy()
			out.Inventory = append(out.Inventory, newItem)
		}
	}

	return out
}

//IsEqual compares the current object with another VirtualServer object
func (d DedicatedServer) IsEqual(d2 DedicatedServer) bool {
	//compare Host struct
	if !d.Host.IsEqual(d2.Host) {
		return false
	}

	//compare asset tag
	if d.AssetTag != d2.AssetTag {
		return false
	}

	return compareInventoryItems(d, d2)
}

func compareInventoryItems(d1 DedicatedServer, d2 DedicatedServer) bool {
	//compare length of inventory items
	if len(d1.Inventory) != len(d2.Inventory) {
		return false
	}

	//sort inventory items
	sort.Slice(d1.Inventory, func(i, j int) bool { return d1.Inventory[i].GetHashableString() < d1.Inventory[j].GetHashableString() })
	sort.Slice(d2.Inventory, func(i, j int) bool { return d2.Inventory[i].GetHashableString() < d2.Inventory[j].GetHashableString() })

	//iterate through inventory items and compare each item using IsEqual
	for key, item1 := range d1.Inventory {
		if !item1.IsEqual(d2.Inventory[key]) {
			return false
		}
	}

	return true
}
