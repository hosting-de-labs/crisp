package model_test

import (
	"strconv"
	"testing"

	"github.com/hosting-de-labs/go-crisp/model"
	"github.com/stretchr/testify/assert"
)

func MockHost() model.Host {
	h := model.NewHost()
	h.Hostname = "host1"
	h.Platform = "Debian "

	return *h
}

func TestHost_HasTag(t *testing.T) {
	host := MockHost()
	assert.False(t, host.HasTag("tag1"))

	host.AddTag("tag2")
	assert.True(t, host.HasTag("tag2"))
}

func TestHost_AddTag(t *testing.T) {
	host := MockHost()

	host.AddTag("tag1")
	assert.True(t, host.HasTag("tag1"))
	assert.False(t, host.HasTag("tag2"))

	host.AddTag("tag2")
	assert.True(t, host.HasTag("tag1"))
	assert.True(t, host.HasTag("tag2"))

	host.AddTag("tag1")
	assert.True(t, host.HasTag("tag1"))
	assert.True(t, host.HasTag("tag2"))
}

func TestHost_Copy(t *testing.T) {
	host1 := MockHost()
	host2 := host1.Copy()
	assert.Equal(t, host1, host2)

	host2.Hostname = "host2"
	assert.NotEqual(t, host1, host2)

	host3 := MockHost()
	host3.AddTag("tag1")
	host4 := host3.Copy()
	assert.Equal(t, host3, host4)

	host3.AddTag("tag2")
	assert.NotEqual(t, host3, host4)
}

func TestHost_IsEqual(t *testing.T) {
	cases := []struct {
		host1   model.Host
		host2   model.Host
		isEqual bool
	}{
		{
			host1:   model.Host{},
			host2:   model.Host{},
			isEqual: true,
		},
		{
			host1: model.Host{
				Hostname:  "Server",
				IsManaged: true,
				Comments: []string{
					"Comment1",
					"Comment2",
				},
				Tags: []string{
					"Tag1",
					"Tag2",
				},
			},
			host2: model.Host{
				Hostname:  "Server",
				IsManaged: true,
				Comments: []string{
					"Comment1",
					"Comment2",
				},
				Tags: []string{
					"Tag1",
					"Tag2",
				},
			},
			isEqual: true,
		},
		{
			host1: model.Host{
				Hostname: "Server1",
			},
			host2: model.Host{
				Hostname: "Server2",
			},
			isEqual: false,
		},
		{
			host1: model.Host{
				IsManaged: true,
			},
			host2: model.Host{
				IsManaged: false,
			},
			isEqual: false,
		},
		{
			host1: model.Host{
				Tags: []string{"Tag1"},
			},
			host2: model.Host{
				Tags: []string{"Tag2"},
			},
			isEqual: false,
		},
		{
			host1: model.Host{
				Tags: []string{"Tag1", "Tag2"},
			},
			host2: model.Host{
				Tags: []string{"Tag1"},
			},
			isEqual: false,
		},
		{
			host1: model.Host{
				Comments: []string{"Comment1"},
			},
			host2: model.Host{
				Comments: []string{"Comment2"},
			},
			isEqual: false,
		},
	}

	for key, testcase := range cases {
		if testcase.isEqual {
			assert.True(t, testcase.host1.IsEqual(testcase.host2), "Case ID: "+strconv.Itoa(key))
		} else {
			assert.False(t, testcase.host1.IsEqual(testcase.host2), "Case ID: "+strconv.Itoa(key))
		}
	}
}
