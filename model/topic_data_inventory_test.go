package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopicDataInventorySerializeDeserialize(t *testing.T) {
	{
		td := mockInventoryTopicData(false, false)

		data, err := td.Serialize()

		assert.Nil(t, err)
		assert.NotEmpty(t, data)

		tdi := &TopicDataInventory{}
		err = tdi.Deserialize(data)
		assert.Nil(t, err)
	}
}

func TestTopicDataInventoryValid(t *testing.T) {
	td1 := mockInventoryTopicData(false, false)
	assert.True(t, td1.Valid())

	td2 := mockInventoryTopicData(true, false)
	assert.True(t, td2.Valid())

	td3 := mockInventoryTopicData(false, true)
	assert.True(t, td3.Valid())

	td4 := mockInventoryTopicData(true, true)
	assert.True(t, td4.Valid())

	td5 := &TopicDataInventory{}
	assert.False(t, td5.Valid())
}

func mockInventoryTopicData(withDetails bool, withTags bool) TopicData {
	itm := NewInventoryItem()
	itm.Type = InventoryItemTypeMainboard
	itm.Manufacturer = "Supermicro"
	itm.Model = "X11DPT-L"
	itm.PartNumber = "1234567890"
	itm.SerialNumber = "1234567890"
	itm.AssetTag = "1234567890"

	// TODO: implement withDetails

	// TODO: implement withTags

	return &TopicDataInventory{
		[]InventoryItem{*itm},
	}
}
