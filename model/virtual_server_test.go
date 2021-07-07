package model_test

import (
	"strconv"
	"testing"

	"github.com/hosting-de-labs/go-crisp/model"
	"github.com/stretchr/testify/assert"
)

func MockVirtualServer() model.VirtualServer {
	vm := model.NewVirtualServer()
	vm.Resources.Cores = 4
	vm.Resources.Disks = []model.VirtualServerDisk{{Size: 10}}
	return *vm
}

func TestVirtualServer_Copy(t *testing.T) {
	vm1 := MockVirtualServer()
	vm1.Resources.Cores = 4
	vm1.Resources.Disks = []model.VirtualServerDisk{{Size: 10}}

	vm2 := vm1.Copy()

	assert.Equal(t, vm1, vm2)
}

func TestVirtualServer_IsEqual(t *testing.T) {
	cases := []struct {
		vm1     model.VirtualServer
		vm2     model.VirtualServer
		isEqual bool
	}{
		{
			vm1: model.VirtualServer{
				Resources: model.VirtualServerResources{
					Cores: 4,
					Disks: []model.VirtualServerDisk{
						{Size: 10},
					},
				},
			},
			vm2: model.VirtualServer{
				Resources: model.VirtualServerResources{
					Cores: 4,
					Disks: []model.VirtualServerDisk{
						{Size: 10},
					},
				},
			},
			isEqual: true,
		},
		{
			vm1: model.VirtualServer{
				Resources: model.VirtualServerResources{
					Cores: 4,
					Disks: []model.VirtualServerDisk{
						{Size: 10},
					},
				},
			},
			vm2: model.VirtualServer{
				Resources: model.VirtualServerResources{
					Cores: 4,
					Disks: []model.VirtualServerDisk{
						{Size: 20},
					},
				},
			},
			isEqual: false,
		},
		{
			vm1: model.VirtualServer{
				Resources: model.VirtualServerResources{
					Cores: 4,
					Disks: []model.VirtualServerDisk{
						{Size: 10},
					},
				},
			},
			vm2: model.VirtualServer{
				Resources: model.VirtualServerResources{
					Cores: 4,
					Disks: []model.VirtualServerDisk{
						{Size: 10},
						{Size: 20},
					},
				},
			},
			isEqual: false,
		},
		{
			vm1: model.VirtualServer{
				Resources: model.VirtualServerResources{
					Cores: 4,
				},
			},
			vm2: model.VirtualServer{
				Resources: model.VirtualServerResources{
					Cores: 2,
				},
			},
			isEqual: false,
		},
		{
			vm1: model.VirtualServer{
				Resources: model.VirtualServerResources{
					Memory: 1024,
				},
			},
			vm2: model.VirtualServer{
				Resources: model.VirtualServerResources{
					Memory: 2048,
				},
			},
			isEqual: false,
		},
	}

	for key, testcase := range cases {
		if testcase.isEqual {
			assert.Equal(t, testcase.vm1, testcase.vm2, "Case ID: "+strconv.Itoa(key))
			assert.True(t, testcase.vm1.IsEqual(testcase.vm2), "Case ID: "+strconv.Itoa(key))
		} else {
			assert.NotEqual(t, testcase.vm1, testcase.vm2, "Case ID: "+strconv.Itoa(key))
			assert.False(t, testcase.vm1.IsEqual(testcase.vm2), "Case ID: "+strconv.Itoa(key))
		}
	}
}
