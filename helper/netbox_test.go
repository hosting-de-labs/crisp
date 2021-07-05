package helper_test

import (
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/hosting-de-labs/go-crisp/helper"
	"github.com/hosting-de-labs/go-crisp/model"
	"github.com/hosting-de-labs/go-netbox/netbox/models"
	"github.com/stretchr/testify/assert"
)

func TestDeviceConvertFromNetbox(t *testing.T) {
	device := mockNetboxDevice()
	assert.NotNil(t, device)

	assert.Equal(t, "host1", *device.Name)
	assert.Equal(t, "123-456", *device.AssetTag)
	assert.Equal(t, "1234567890", device.Serial)
}

func TestDeviceConvertFromNetbox_WithWrongType(t *testing.T) {
	var device interface{}
	res, err := helper.DeviceConvert(device)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestDeviceConvertFromNetbox_WithIPAddresses(t *testing.T) {
	res, err := helper.DeviceConvert(mockNetboxDevice())
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestVlanConvertFromNetbox(t *testing.T) {
	vlan400, err := helper.VlanConvert(mockNetboxVlan())
	assert.Nil(t, err)

	assert.Equal(t, uint16(400), vlan400.ID)
	assert.Equal(t, "Public VLAN", vlan400.Name)

	assert.Equal(t, model.VLANStatus(model.VLANStatusActive), vlan400.Status)

	assert.Equal(t, "This is Public VLAN description", vlan400.Description)
	assert.Equal(t, []string{"Green Tag"}, vlan400.Tags)
}

func TestVlanConvertFromNetboxWithNestedVlan(t *testing.T) {
	vlan600, err := helper.VlanConvert(mockNetboxNestedVlan())
	assert.Nil(t, err)

	assert.Equal(t, uint16(600), vlan600.ID)
	assert.Equal(t, "Private VLAN", vlan600.Name)
}

func TestVirtualMachineConvertFromNetbox_WithUnknownType(t *testing.T) {

	type UnknownVMType interface{}
	var vm UnknownVMType

	_, err := helper.VMConvert(vm)
	assert.Error(t, err)
}

func TestVirtualMachineConvertFromNetbox_WithResources(t *testing.T) {
	vm := mockNetboxVirtualMachine(true, false, false)

	res, err := helper.VMConvert(vm)
	assert.Equal(t, err, nil)

	assert.Equal(t, res.Resources.Cores, 1)
	assert.Equal(t, res.Resources.Memory, int64(4096))
	assert.Equal(t, len(res.Resources.Disks), 1)
	assert.Equal(t, res.Resources.Disks[0].Size, int64(10240*1024))
}

func TestVirtualMachineConvertFromNetbox_WithIPAddresses(t *testing.T) {
	vm := mockNetboxVirtualMachine(false, false, false)

	_, err := helper.VMConvert(vm)
	assert.Equal(t, err, nil)
}

func TestVirtualMachineConvertFromNetbox_WithCustomFields(t *testing.T) {
	vm := mockNetboxVirtualMachine(false, false, true)
	res, err := helper.VMConvert(vm)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func mockNetboxDevice() models.DeviceWithConfigContext {
	return models.DeviceWithConfigContext{
		AssetTag:    swag.String("123-456"),
		Created:     strfmt.Date(time.Now()),
		DisplayName: "",
		ID:          10,
		LastUpdated: strfmt.DateTime(time.Now()),
		Name:        swag.String("host1"),
		Serial:      "1234567890",
		Status: &models.DeviceWithConfigContextStatus{
			Label: swag.String("Active"),
			Value: swag.String("active"),
		},
		PrimaryIp4: &models.NestedIPAddress{Address: swag.String("123.456.789.101112")},
		PrimaryIp6: &models.NestedIPAddress{Address: swag.String("::827")},
	}
}

func mockNetboxDeviceInterface() models.Interface {
	return models.Interface{
		Enabled:  true,
		ID:       1,
		Label:    "eth0",
		MgmtOnly: false,
		Name:     swag.String("eth0"),
	}
}

func mockNetboxVlan() models.VLAN {
	return models.VLAN{
		ID:   10,
		Vid:  swag.Int64(400),
		Name: swag.String("Public VLAN"),
		Status: &models.VLANStatus{
			Value: swag.String("active"),
			Label: swag.String("Active"),
		},
		Description: "This is Public VLAN description",
		Tags: []*models.NestedTag{
			{
				Color: "green",
				ID:    1,
				Name:  swag.String("Green Tag"),
				Slug:  "green-tag",
				URL:   "https://netbox/api",
			},
		},
	}
}

func mockNetboxNestedVlan() models.NestedVLAN {
	return models.NestedVLAN{
		ID:   20,
		Vid:  swag.Int64(600),
		Name: swag.String("Private VLAN"),
	}
}

func mockNetboxVirtualMachine(addResources bool, addIPAddresses bool, addCustomFields bool) (out models.VirtualMachineWithConfigContext) {
	out.ID = 10
	out.Name = swag.String("VM1")

	if addResources {
		out.Vcpus = swag.Int64(1)
		out.Memory = swag.Int64(4096)
		out.Disk = swag.Int64(10240)
	}

	if addIPAddresses {
		//TODO: add interfaces when adding ip addresses
		out.PrimaryIp4 = &models.NestedIPAddress{Address: swag.String("127.0.0.1/32")}
		out.PrimaryIp6 = &models.NestedIPAddress{Address: swag.String("::1/128")}
	}

	if addCustomFields {
		customFields := make(map[string]interface{})
		customFields["hypervisor_label"] = "Hypervisor1"

		out.CustomFields = customFields
	}

	return out
}
