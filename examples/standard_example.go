package main

import (
	"fmt"
	"github.com/zulandar/go-dhcp-management"
	"net"
)


func main() {
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


	parseResponse := go_dhcp_management.Parse("./tests/dhcpd.conf")
	generatedResponse := go_dhcp_management.Generate(fileConfig, "./tests/dhcpd.conf.example")
	parseLeases := go_dhcp_management.ParseLeases("./tests/dhcpd.leases")

	fmt.Printf("Results %v\n", parseResponse)

	fmt.Printf("Generated Results: %v", generatedResponse)

	fmt.Printf("Generated Results: %v", parseLeases)
}
