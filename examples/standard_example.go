package main

import (
	"fmt"
	"github.com/zulandar/go-dhcp-management"
)


func main() {
	response := go_dhcp_management.Parse("./tests/dhcpd.conf")

	fmt.Printf("Results %v", response)
}
