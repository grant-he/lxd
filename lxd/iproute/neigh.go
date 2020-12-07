package iproute

import (
	"net"
	"strings"

	"github.com/grant-he/lxd/shared"
)

// GetNeighbourIPs returns the IP addresses in the neighbour cache for a particular interface and MAC.
func GetNeighbourIPs(interfaceName string, hwaddr string) ([]net.IP, error) {
	addresses := []net.IP{}

	// Look for neighbour entries for IPv.
	out, err := shared.RunCommand("ip", "neigh", "show", "dev", interfaceName)
	if err == nil {
		for _, line := range strings.Split(out, "\n") {
			// Split fields and early validation.
			fields := strings.Fields(line)
			if len(fields) != 4 {
				continue
			}

			if fields[2] != hwaddr {
				continue
			}

			ip := net.ParseIP(fields[0])
			if ip != nil {
				addresses = append(addresses, ip)
			}
		}
	}

	return addresses, nil
}
