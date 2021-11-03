package main

import "github.com/zulandar/go-dhcp-management"


func main() {
	go_dhcp_management.Parse("./tests/dhcpd.conf")
}
