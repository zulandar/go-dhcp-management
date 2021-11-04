package go_dhcp_management

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Generate(config FileConfig, fileLocation string) bool {
	if _, err := os.Stat(fileLocation); err == nil {
		err = os.Remove(fileLocation)
		if err != nil {
			log.Fatalf("Failed to remove file: %s", err)
		}
	}

	//file, err := os.Create(fileLocation)
	//
	//if err != nil {
	//	log.Fatalf("Failed to create file: %s", err)
	//}
	//
	//defer file.Close()

	if err := os.WriteFile(fileLocation, []byte("option domain-name " + config.DomainName + "\n" +
	"option domain-name-servers " + strings.Join(config.DomainNameServers[:], ", ") + "\n" +
		"default-lease-time " + strconv.FormatInt(config.DefaultLeaseTime, 10) + "\n" +
		"max-lease-time " + strconv.FormatInt(config.MaxLeaseTime, 10) + "\n" +
		"subnet " + config.SubnetConfig.Subnet.String() + " netmask " + config.SubnetConfig.Netmask + " {\n" +
		"range " + config.SubnetConfig.RangeStart.String() + " " + config.SubnetConfig.RangeEnd.String() + ";\n" +
		"option routers " + config.SubnetConfig.OptionRouter + ";\n}\n"), 0666); err != nil {
		log.Printf("Failed to write to file: %v", err)
		return false
	}

	//_, err = fmt.Fprintln(file, "option domain-name ", config.DomainName)
	//
	//if err != nil {
	//	log.Fatalf("Failed to write line to file: %s", err)
	//}
	//
	//_, err = fmt.Fprintln(file, "option domain-name-servers ", strings.Join(config.DomainNameServers[:], ", "))
	//
	//if err != nil {
	//	log.Fatalf("Failed to write line to file: %s", err)
	//}
	//
	//_, err = fmt.Fprintln(file, "default-lease-time ", config.DefaultLeaseTime)
	//
	//if err != nil {
	//	log.Fatalf("Failed to write line to file: %s", err)
	//}
	//
	//_, err = fmt.Fprintln(file, "max-lease-time ", config.MaxLeaseTime)
	//
	//if err != nil {
	//	log.Fatalf("Failed to write line to file: %s", err)
	//}
	//
	//_, err = fmt.Fprintln(file, "subnet ", config.SubnetConfig.Subnet, " netmask ", config.SubnetConfig.Netmask, " {")
	//
	//if err != nil {
	//	log.Fatalf("Failed to write line to file: %s", err)
	//}
	//
	//_, err = fmt.Fprintln(file, "range ", config.SubnetConfig.RangeStart, " ", config.SubnetConfig.RangeEnd, ";")
	//
	//if err != nil {
	//	log.Fatalf("Failed to write line to file: %s", err)
	//}
	//
	//_, err = fmt.Fprintln(file, "option routers", strings.ReplaceAll(config.SubnetConfig.OptionRouter, " ", ""), ";")
	//
	//if err != nil {
	//	log.Fatalf("Failed to write line to file: %s", err)
	//}

	return true
}