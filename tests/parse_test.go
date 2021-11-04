package tests

import (
	"github.com/stretchr/testify/assert"
	go_dhcp_management "github.com/zulandar/go-dhcp-management"
	"net"
	"testing"
)

func TestParse(t *testing.T) {
	parse := go_dhcp_management.Parse("./dhcpd.conf")

	assert.Equal(t, "example.org", parse.DomainName)
	assert.Equal(t, "ns1.example.org", parse.DomainNameServers[0])
	assert.Equal(t, "ns2.example.org", parse.DomainNameServers[1])
	assert.Equal(t, int64(600), parse.DefaultLeaseTime)
	assert.Equal(t, int64(7200), parse.MaxLeaseTime)
	assert.Equal(t, net.ParseIP("192.168.1.0"), parse.SubnetConfig.Subnet)
	assert.Equal(t, "255.255.255.0", parse.SubnetConfig.Netmask)
	assert.Equal(t, net.ParseIP("192.168.1.2"), parse.SubnetConfig.RangeStart)
	assert.Equal(t, net.ParseIP("192.168.1.100"), parse.SubnetConfig.RangeEnd)
	assert.Equal(t, "softwaredev", parse.SubnetConfig.OptionRouter)
}