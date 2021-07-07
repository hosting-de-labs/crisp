package model_test

import (
	"testing"

	"github.com/hosting-de-labs/go-crisp/model"
	"github.com/stretchr/testify/assert"
)

func MockIpv4Address() (out model.IPAddress) {
	return model.IPAddress{
		Family:      model.IPAddressFamilyIPv4,
		Address:     "192.168.10.1",
		CIDR:        24,
		Status:      model.IPAddressStatusActive,
		Description: "An internal ip address",
		Tags:        []string{"internal", "netbox-sync"},
	}
}

func MockIpv6Address() (out model.IPAddress) {
	return model.IPAddress{
		Family:      model.IPAddressFamilyIPv6,
		Address:     "fc00::1",
		CIDR:        128,
		Status:      model.IPAddressStatusActive,
		Description: "An internal ip address",
		Tags:        []string{"internal", "netbox-sync"},
	}
}

func TestIPAddress_String(t *testing.T) {
	ip1 := MockIpv4Address()
	assert.Equal(t, ip1.String(), "192.168.10.1/24")

	ip2 := MockIpv6Address()
	assert.Equal(t, ip2.String(), "fc00::1/128")
}

func TestIPAddress_IsEqual(t *testing.T) {
	ip1 := MockIpv4Address()

	ip2 := ip1.Clone()
	assert.Equal(t, ip1, ip2)
	assert.True(t, ip1.IsEqual(ip2))

	ip2.Address = "192.168.10.2"
	assert.NotEqual(t, ip1, ip2)
}
