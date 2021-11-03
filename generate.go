package go_dhcp_management

import (
	"fmt"
	"log"
	"os"
)

func Generate(config FileConfig, fileLocation string) bool {
	if _, err := os.Stat(fileLocation); err == nil {
		err = os.Remove(fileLocation)
		if err != nil {
			log.Fatalf("Failed to remove file: %s", err)
		}
	}

	file, err := os.Create(fileLocation)

	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}

	defer file.Close()

	_, err = fmt.Fprintln(file, "option domain-name ", config.DomainName)

	if err != nil {
		log.Fatalf("Failed to write line to file: %s", err)
	}

	_, err = fmt.Fprintln(file, "option domain-name-servers ", config.DomainNameServers)

	if err != nil {
		log.Fatalf("Failed to write line to file: %s", err)
	}

	_, err = fmt.Fprintln(file, "default-lease-time ", config.DefaultLeaseTime)

	if err != nil {
		log.Fatalf("Failed to write line to file: %s", err)
	}

	_, err = fmt.Fprintln(file, "max-lease-time ", config.MaxLeaseTime)

	if err != nil {
		log.Fatalf("Failed to write line to file: %s", err)
	}

	_, err = fmt.Fprintln(file, "subnet ", config.SubnetConfig.Subnet, " netmask ", config.SubnetConfig.Netmask, " {")

	if err != nil {
		log.Fatalf("Failed to write line to file: %s", err)
	}

	_, err = fmt.Fprintln(file, "range ", config.SubnetConfig.RangeStart, " ", config.SubnetConfig.RangeEnd, ";")

	if err != nil {
		log.Fatalf("Failed to write line to file: %s", err)
	}

	_, err = fmt.Fprintln(file, "option routers ", config.SubnetConfig.OptionRouter, ";")

	if err != nil {
		log.Fatalf("Failed to write line to file: %s", err)
	}

	return true
}