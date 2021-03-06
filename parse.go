package go_dhcp_management

import (
	"bufio"
	leases "github.com/npotts/go-dhcpd-leases"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type FileConfig struct {
	DomainName string `json:"domain-name"`
	DomainNameServers []string `json:"domain-name-servers"`
	DefaultLeaseTime int64 `json:"default-lease-time"`
	MaxLeaseTime int64 `json:"max-lease-time"`
	SubnetConfig SubnetConfig `json:"subnet-config"`
}

type SubnetConfig struct {
	Subnet net.IP `json:"subnet"`
	Netmask string `json:"netmask"`
	RangeStart net.IP `json:"range-start"`
	RangeEnd net.IP `json:"range-end"`
	OptionRouter string `json:"option-router"`
}

func Parse(fileLocation string) FileConfig {
	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatalf("Error opening file: %s - with error: %s", fileLocation, err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	subnetConfig := SubnetConfig{}
	fileConfig := FileConfig{SubnetConfig: subnetConfig}
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "#") {
			continue
		}

		if strings.Contains(scanner.Text(), "domain-name ") {
			s := strings.Trim(scanner.Text(), "option domain-name ")
			s = strings.ReplaceAll(s, "\"", "")
			fileConfig.DomainName = strings.ReplaceAll(s, ";", "")
		} else if strings.Contains(scanner.Text(), "domain-name-servers ") {
			s := strings.ReplaceAll(scanner.Text(), "option domain-name-servers", "")
			s = strings.ReplaceAll(s, " ", "")
			s = strings.ReplaceAll(s, ";", "")
			fileConfig.DomainNameServers = strings.Split(s, ",")
		} else if strings.Contains(scanner.Text(), "default-lease-time ") {
			s := strings.Trim(scanner.Text(), "default-lease-time ")
			s = strings.Trim(s, ";")
			finalS, _ := strconv.ParseInt(s, 0, 64)
			fileConfig.DefaultLeaseTime = finalS
		} else if strings.Contains(scanner.Text(), "max-lease-time ") {
			s := strings.Trim(scanner.Text(), "max-lease-time ")
			s = strings.Trim(s, ";")
			finalS, _ := strconv.ParseInt(s, 0, 64)
			fileConfig.MaxLeaseTime = finalS
		} else if strings.Contains(scanner.Text(), "subnet") && strings.Contains(scanner.Text(), "netmask") {
			subnetIP := strings.Trim(scanner.Text(), "subnet")
			subnetIP = strings.TrimSpace(subnetIP)
			subnetIP = strings.ReplaceAll(subnetIP, " {", "")
			subnetSplit := strings.Split(subnetIP, "netmask")
			fileConfig.SubnetConfig.Subnet = net.ParseIP(strings.TrimSpace(subnetSplit[0]))
			fileConfig.SubnetConfig.Netmask = strings.TrimRight(subnetSplit[1], "{")
			fileConfig.SubnetConfig.Netmask = strings.TrimSpace(subnetSplit[1])
		} else if strings.Contains(scanner.Text(), "range ") {
			rangeBlock := strings.Trim(scanner.Text(), "range ")
			rangeBlock = strings.Trim(rangeBlock, ";")
			rangeBlock = strings.ReplaceAll(rangeBlock, "   ", " ")
			rangeBlock = strings.TrimSpace(rangeBlock)
			rangeSplit := strings.Split(rangeBlock, " ")
			fileConfig.SubnetConfig.RangeStart = net.ParseIP(strings.TrimSpace(rangeSplit[0]))
			fileConfig.SubnetConfig.RangeEnd = net.ParseIP(strings.TrimSpace(rangeSplit[1]))
		} else if strings.Contains(scanner.Text(), "option routers ") {
			s := strings.TrimSpace(scanner.Text())
			s = strings.ReplaceAll(s, "option routers ", "")
			s = strings.Trim(s, ";")
			fileConfig.SubnetConfig.OptionRouter = s
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	return fileConfig

}

func ParseLeases(fileLocation string) []leases.Lease {
	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatalf("Error opening file: %s - with error: %s", fileLocation, err)
	}

	defer file.Close()

	leaseData := leases.Parse(file)

	return leaseData
}