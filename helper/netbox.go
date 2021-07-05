package helper

import (
	"fmt"
	"net"
	"reflect"
	"strings"

	"github.com/hosting-de-labs/go-crisp/model"
	"github.com/hosting-de-labs/go-netbox/netbox/models"
)

type Netbox struct{}

func DeviceConvert(device interface{}) (out *model.DedicatedServer, err error) {
	out = model.NewDedicatedServer()

	switch device.(type) {
	case models.Device:
		d := device.(models.Device)

		if d.Name != nil {
			out.Hostname = *d.Name
		}

		if d.AssetTag != nil {
			out.AssetTag = *d.AssetTag
		}

		out.SerialNumber = d.Serial
		out.Comments = strings.Split(d.Comments, "\n") //TODO: use utils.ParseVMComment
	case models.DeviceWithConfigContext:
		d := device.(models.DeviceWithConfigContext)

		if d.Name != nil {
			out.Hostname = *d.Name
		}

		if d.AssetTag != nil {
			out.AssetTag = *d.AssetTag
		}

		out.SerialNumber = d.Serial
		out.Comments = strings.Split(d.Comments, "\n") //TODO: use utils.ParseVMComment
	default:
		return nil, fmt.Errorf("unsupported type for device: %s", reflect.TypeOf(device))
	}

	return out, nil
}

func (n Netbox) DeviceInterfaceConvert(nbIf models.Interface) (*model.NetworkInterface, error) {
	netIf := model.NewNetworkInterface()

	if nbIf.Type != nil {
		netIf.Type = model.InterfaceType(*nbIf.Type.Value)
	}

	netIf.Enabled = nbIf.Enabled

	if nbIf.Name != nil {
		netIf.Name = *nbIf.Name
	}

	if nbIf.MacAddress != nil {
		mac, err := net.ParseMAC(*nbIf.MacAddress)
		if err != nil {
			return nil, err
		}

		netIf.MACAddress = mac
	}

	if nbIf.UntaggedVlan != nil {
		vlan, err := VlanConvert(*nbIf.UntaggedVlan)
		if err != nil {
			return nil, err
		}

		netIf.UntaggedVlan = vlan
	}

	if len(nbIf.TaggedVlans) > 0 {
		for _, taggedVlan := range nbIf.TaggedVlans {
			vlan, err := VlanConvert(*taggedVlan)
			if err != nil {
				return nil, err
			}

			netIf.TaggedVlans = append(netIf.TaggedVlans, *vlan)
		}
	}

	return netIf, nil
}

func (n Netbox) InventoryItemConvert(i models.InventoryItem) (out model.InventoryItem) {
	out = *model.NewInventoryItem()

	out.Manufacturer = *i.Manufacturer.Name
	out.PartNumber = i.PartID
	out.SerialNumber = i.Serial
	out.AssetTag = *i.AssetTag

	result := strings.SplitN(*i.Name, ":", 1)
	switch len(result) {
	case 0, 1:
		out.Type = model.InventoryItemTypeOther
		out.Model = *i.Name
	default:
		out.Type, _ = model.InventoryItemTypeParse(result[0])
		out.Model = result[1]
	}

	return out
}

func VlanConvert(netboxVlan interface{}) (*model.VLAN, error) {
	vlan := model.VLAN{}
	var vlanStatus string

	switch v := netboxVlan.(type) {
	case models.VLAN:
		vlan.ID = uint16(*v.Vid)
		vlan.Name = *v.Name
		vlan.Description = v.Description

		for _, tag := range v.Tags {
			vlan.Tags = append(vlan.Tags, *tag.Name)
		}

		vlanStatus = *v.Status.Value
	case models.NestedVLAN:
		vlan.ID = uint16(*v.Vid)
		vlan.Name = *v.Name
	default:
		return nil, fmt.Errorf("vlan has to be of type VLAN oder NestedVLAN, type is %T", netboxVlan)
	}

	switch vlanStatus {
	case "":
		fallthrough
	case "unknown":
		vlan.Status = model.VLANStatusUnknown
	case "active":
		vlan.Status = model.VLANStatusActive
	case "reserved":
		vlan.Status = model.VLANStatusReserved
	case "deprecated":
		vlan.Status = model.VLANStatusDeprecated

	default:
		return nil, fmt.Errorf("unknown vlan status %s", vlanStatus)
	}

	return &vlan, nil
}

func VMInterfaceConvert(netboxInterface models.VMInterface) (*model.NetworkInterface, error) {
	netIf := model.NewNetworkInterface()

	if netboxInterface.Name != nil {
		netIf.Name = *netboxInterface.Name
	}

	if netboxInterface.MacAddress != nil {
		mac, err := net.ParseMAC(*netboxInterface.MacAddress)
		if err != nil {
			return nil, err
		}

		netIf.MACAddress = mac
	}

	if netboxInterface.UntaggedVlan != nil {
		vlan, err := VlanConvert(*netboxInterface.UntaggedVlan)
		if err != nil {
			return nil, err
		}

		netIf.UntaggedVlan = vlan
	}

	if len(netboxInterface.TaggedVlans) > 0 {
		for _, taggedVlan := range netboxInterface.TaggedVlans {
			vlan, err := VlanConvert(*taggedVlan)
			if err != nil {
				return nil, err
			}

			netIf.TaggedVlans = append(netIf.TaggedVlans, *vlan)
		}
	}

	return netIf, nil
}

func VMConvert(netboxVM interface{}) (out *model.VirtualServer, err error) {
	out = model.NewVirtualServer()

	switch netboxVM.(type) {
	case models.VirtualMachineWithConfigContext:
		vm := netboxVM.(models.VirtualMachineWithConfigContext)

		out.Hostname = *vm.Name

		if vm.Vcpus != nil {
			out.Resources.Cores = int(*vm.Vcpus)
		}

		if vm.Memory != nil {
			out.Resources.Memory = *vm.Memory
		}

		if vm.Disk != nil {
			out.Resources.Disks = append(out.Resources.Disks, model.VirtualServerDisk{
				Size: *vm.Disk * 1024,
			})
		}
	default:
		return nil, fmt.Errorf("unsupported type for device: %s", reflect.TypeOf(netboxVM))
	}

	for _, tag := range out.Tags {
		if tag == "managed" {
			out.IsManaged = true
			break
		}
	}

	return out, nil
}
