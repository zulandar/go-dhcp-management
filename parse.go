package go_dhcp_management

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type FileConfig struct {
	DomainName string `json:"domain-name"`
	DomainNameServers []string `json:"domain-name-servers"`
	DefaultLeaseTime int `json:"default-lease-time"`
	MaxLeaseTime int `json:"max-lease-time"`
	SubnetConfig SubnetConfig `json:"subnet-config"`
}

type SubnetConfig struct {
	Subnet net.IP `json:"subnet"`
	Netmask net.IPMask `json:"netmask"`
	RangeStart net.IP `json:"range-start"`
	RangeEnd net.IP `json:"range-end"`
	OptionRouter string `json:"option-router"`
}

func Parse(fileLocation string) {
	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatalf("Error opening file: %s - with error: %s", fileLocation, err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "#") {
			continue
		}

		if strings.Contains(scanner.Text(), "domain-name ") {
			fmt.Printf("Domain name found")
		} else if strings.Contains(scanner.Text(), "domain-name-servers ") {
			fmt.Printf("Domain servers")
		} else if strings.Contains(scanner.Text(), "default-lease-time ") {
			fmt.Printf("Default lease time")
		} else if strings.Contains(scanner.Text(), "max-lease-time ") {
			fmt.Printf("Max Lease Time")
		} else if strings.Contains(scanner.Text(), "subnet ") && strings.Contains(scanner.Text(), "netmask ") {
			fmt.Printf("Subnet and Netmask")
		} else if strings.Contains(scanner.Text(), "range ") {
			fmt.Printf("range")
		} else if strings.Contains(scanner.Text(), "option routers ") {
			fmt.Printf("Option Routers")
		}
		//fmt.Printf("%s\n",scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}


}