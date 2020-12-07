package iproute

import "github.com/grant-he/lxd/shared"

// VFSetMAC sets MAC on VF
func VFSetMAC(parent string, vfID string, mac string) error {
	_, err := shared.TryRunCommand("ip", "link", "set", "dev", parent, "vf", vfID, "mac", mac)
	return err
}

// VFSetSpoofchk sets spoof checking to mode
func VFSetSpoofchk(parent string, vfID string, mode string) error {
	_, err := shared.TryRunCommand("ip", "link", "set", "dev", parent, "vf", vfID, "spoofchk", mode)
	return err
}

// VFSetVLAN sets up VF VLAN
func VFSetVLAN(parent string, vfID string, vlan string) error {
	_, err := shared.TryRunCommand("ip", "link", "set", "dev", parent, "vf", vfID, "vlan", vlan)
	return err
}
