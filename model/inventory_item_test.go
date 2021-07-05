package model_test

import (
	"testing"

	"github.com/hosting-de-labs/go-crisp/model"
	"github.com/stretchr/testify/assert"
)

func MockInventoryItem() model.InventoryItem {
	itm := model.NewInventoryItem()
	itm.Type = model.InventoryItemTypeProcessor
	itm.Manufacturer = "Intel"
	itm.Model = "Xeon X5670"
	itm.AssetTag = "Asset Tag"
	itm.PartNumber = "Part Number"
	itm.SerialNumber = "Serial Number"

	return *itm
}

func TestInventoryItem_Copy(t *testing.T) {
	item := MockInventoryItem()
	item2 := item.Copy()

	assert.Equal(t, item, item2)
	assert.True(t, item.IsEqual(item2))
}

func TestInventoryItem_IsEqual(t *testing.T) {
	item := MockInventoryItem()
	item.Details["Cores"] = "2"
	item.Details["Threads"] = "4"

	item2 := item.Copy()
	assert.Equal(t, item, item2)

	item2.Details["L3 Cache"] = "12MB"
	assert.NotEqual(t, item, item2)
}
