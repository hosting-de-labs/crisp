package model

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/hosting-de-labs/go-crisp/utils"
)

//InventoryItem represents a device / module that is attached to / built into a dedicated server
type InventoryItem struct {
	Type         InventoryItemType
	Manufacturer string
	Model        string
	PartNumber   string
	SerialNumber string
	AssetTag     string
	Details      map[string]string
	Tags         []string
}

func NewInventoryItem() *InventoryItem {
	return &InventoryItem{
		Details: make(map[string]string),
	}
}

func (i InventoryItem) GetHashableString() string {
	str := fmt.Sprintf("%s:%s:%s:%s:%s", i.Manufacturer, i.Model, i.PartNumber, i.AssetTag, i.SerialNumber)

	if len(i.Details) > 0 {
		var items []string

		for key, val := range i.Details {
			items = append(items, fmt.Sprintf("%s:%s", key, val))
		}

		sort.Strings(items)

		str = fmt.Sprintf("%s:details{%s}", str, strings.Join(items, ","))
	}

	str = strings.Replace(str, " ", "", -1)

	return str
}

func (i InventoryItem) Copy() (out InventoryItem) {
	out = *NewInventoryItem()
	out.Type = i.Type
	out.Manufacturer = i.Manufacturer
	out.Model = i.Model
	out.PartNumber = i.PartNumber
	out.SerialNumber = i.SerialNumber
	out.AssetTag = i.AssetTag

	for key, val := range i.Details {
		out.Details[key] = val
	}

	return out
}

//IsEqual compares an InventoryItem with another one
func (i InventoryItem) IsEqual(i2 InventoryItem) bool {
	if !utils.CompareStruct(i, i2, []string{}, []string{"Details"}) {
		return false
	}

	for key, val1 := range i.Details {
		if val2, ok := i2.Details[key]; !ok || val1 != val2 {
			return false
		}
	}

	return true
}

func (i InventoryItem) String() string {
	res, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(res)
}
