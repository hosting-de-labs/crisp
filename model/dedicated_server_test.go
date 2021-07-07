package model_test

import (
	"strconv"
	"testing"

	"github.com/hosting-de-labs/go-crisp/model"
	"github.com/stretchr/testify/assert"
)

func mockDedicatedServer() model.DedicatedServer {
	d := model.NewDedicatedServer()
	d.Hostname = "host1"
	d.IsManaged = false
	d.Inventory = []model.InventoryItem{
		{
			Type:         model.InventoryItemTypeProcessor,
			Manufacturer: "Intel",
			Model:        "Xeon X5660",
		},
	}

	return *d
}

func TestDedicatedServer_Copy(t *testing.T) {
	host1 := mockDedicatedServer()
	host2 := host1.Copy()
	assert.True(t, host1.IsEqual(host2))

	host1.Inventory = append(host1.Inventory, model.InventoryItem{
		Type:         model.InventoryItemTypeMainboard,
		Manufacturer: "Supermicro",
		Model:        "X9SCL-F",
	})
	assert.False(t, host1.IsEqual(host2))
}

//TODO: move to inventory_item.go
func TestDedicatedServer_IsEqual(t *testing.T) {
	cases := []struct {
		host1   model.DedicatedServer
		host2   model.DedicatedServer
		isEqual bool
	}{
		{
			host1: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{
						Manufacturer: "unknown",
						Model:        "unknown",
						AssetTag:     "asset tag",
						PartNumber:   "part number",
						SerialNumber: "serial number",
					},
				},
			},
			host2: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{
						Manufacturer: "unknown",
						Model:        "unknown",
						AssetTag:     "asset tag",
						PartNumber:   "part number",
						SerialNumber: "serial number",
					},
				},
			},
			isEqual: true,
		},
		{
			host1: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{Manufacturer: "unknown"},
				},
			},
			host2: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{Manufacturer: "u. n. owen"},
				},
			},
			isEqual: false,
		},
		{
			host1: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{Model: "unknown"},
				},
			},
			host2: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{Model: "u. n. owen"},
				},
			},
			isEqual: false,
		},
		{
			host1: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{PartNumber: "unknown"},
				},
			},
			host2: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{PartNumber: "u. n. owen"},
				},
			},
			isEqual: false,
		},
		{
			host1: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{AssetTag: "unknown"},
				},
			},
			host2: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{AssetTag: "u. n. owen"},
				},
			},
			isEqual: false,
		},
		{
			host1: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{SerialNumber: "unknown"},
				},
			},
			host2: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{SerialNumber: "u. n. owen"},
				},
			},
			isEqual: false,
		},
		{
			host1: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{Manufacturer: "Intel"},
					{Manufacturer: "AMD"},
				},
			},
			host2: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{Manufacturer: "AMD"},
					{Manufacturer: "Intel"},
				},
			},
			isEqual: true,
		},
		{
			host1: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{Manufacturer: "Intel"},
					{Manufacturer: "AMD"},
				},
			},
			host2: model.DedicatedServer{
				Inventory: []model.InventoryItem{
					{Manufacturer: "AMD"},
				},
			},
			isEqual: false,
		},
		//TODO: Details
	}

	for key, testcase := range cases {
		if testcase.isEqual {
			assert.True(t, testcase.host1.IsEqual(testcase.host2), "Case ID: "+strconv.Itoa(key))
		} else {
			assert.False(t, testcase.host1.IsEqual(testcase.host2), "Case ID: "+strconv.Itoa(key))
		}
	}
}
