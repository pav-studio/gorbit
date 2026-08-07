package utils

import (
	"net"
	"strings"
)

type NetworkAddress struct {
	Name  string
	Value string
}

func GetNetworkAddresses() ([]NetworkAddress, error) {
	addresses := make([]NetworkAddress, 0)

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	// Ignore common virtual interfaces
	skip := []string{
		"docker",
		"wsl",
		"hyper-v",
		"virtual",
		"vmware",
		"vbox",
		"loopback",
		"teredo",
		"isatap",
		"bluetooth",
	}

	for _, iface := range interfaces {

		// Interface must be up
		if iface.Flags&net.FlagUp == 0 {
			continue
		}

		// Skip loopback interface
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		name := strings.ToLower(iface.Name)

		ignore := false
		for _, s := range skip {
			if strings.Contains(name, s) {
				ignore = true
				break
			}
		}

		if ignore {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {

			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			ip := ipNet.IP.To4()
			if ip == nil {
				continue
			}

			addresses = append(addresses, NetworkAddress{
				Name:  iface.Name,
				Value: ip.String(),
			})
		}
	}

	return addresses, nil
}