package model

//VirtualServer represents a virtual server
type VirtualServer struct {
	Host

	Resources VirtualServerResources
}

//NewVirtualServer returns a new instance of VirtualServer
func NewVirtualServer() *VirtualServer {
	return &VirtualServer{
		Resources: VirtualServerResources{},
	}
}

//Copy creates a deep copy of a VirtualServer object
func (vm VirtualServer) Copy() (out VirtualServer) {
	out.Host = vm.Host.Copy()

	out.Resources = VirtualServerResources{
		Cores:  vm.Resources.Cores,
		Memory: vm.Resources.Memory,
	}

	//copy disks
	out.Resources.Disks = make([]VirtualServerDisk, len(vm.Resources.Disks))
	copy(out.Resources.Disks, vm.Resources.Disks)

	return out
}

//IsEqual compares the current object with another VirtualServer object
func (vm VirtualServer) IsEqual(vm2 VirtualServer) bool {
	//compare Host struct
	if !vm.Host.IsEqual(vm2.Host) {
		return false
	}

	//Resources
	if vm.Resources.Cores != vm2.Resources.Cores {
		return false
	}

	if vm.Resources.Memory != vm2.Resources.Memory {
		return false
	}

	if len(vm.Resources.Disks) != len(vm2.Resources.Disks) {
		return false
	}

	for key, disk := range vm.Resources.Disks {
		if !disk.IsEqual(vm2.Resources.Disks[key]) {
			return false
		}
	}

	return true
}
