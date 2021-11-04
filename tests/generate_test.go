package tests

import (
	"github.com/stretchr/testify/assert"
	go_dhcp_management "github.com/zulandar/go-dhcp-management"
	"net"
	"testing"
)

func TestGenerator(t *testing.T) {
	domainNameServers := []string{"ns5.foobar.local", "ns6.foobar.local"}
	fileConfig := go_dhcp_management.FileConfig{
		DomainName:        "foobar.local",
		DomainNameServers: domainNameServers,
		DefaultLeaseTime:  100,
		MaxLeaseTime:      3600,
		SubnetConfig:      go_dhcp_management.SubnetConfig{
			Subnet: net.ParseIP("10.0.0.0"),
			Netmask: "255.255.255.0",
			RangeStart: net.ParseIP("10.0.0.2"),
			RangeEnd: net.ParseIP("10.0.0.254"),
			OptionRouter: "softwaredev2",
		},
	}

	generator := go_dhcp_management.Generate(fileConfig, "./dhcpd.conf.example")

	assert.True(t, generator)

	// Let's parse to check the config file write correctly

	parse := go_dhcp_management.Parse("./dhcpd.conf.example")

	assert.Equal(t, fileConfig.DomainName, parse.DomainName)
	assert.Equal(t, fileConfig.DomainNameServers[0], parse.DomainNameServers[0])
	assert.Equal(t, fileConfig.DomainNameServers[1], parse.DomainNameServers[1])
	assert.Equal(t, fileConfig.DefaultLeaseTime, parse.DefaultLeaseTime)
	assert.Equal(t, fileConfig.MaxLeaseTime, parse.MaxLeaseTime)
	assert.Equal(t, fileConfig.SubnetConfig.Subnet, parse.SubnetConfig.Subnet)
	assert.Equal(t, fileConfig.SubnetConfig.Netmask, parse.SubnetConfig.Netmask)
	assert.Equal(t, fileConfig.SubnetConfig.RangeStart, parse.SubnetConfig.RangeStart)
	assert.Equal(t, fileConfig.SubnetConfig.RangeEnd, parse.SubnetConfig.RangeEnd)
	assert.Equal(t, fileConfig.SubnetConfig.OptionRouter, parse.SubnetConfig.OptionRouter)


}