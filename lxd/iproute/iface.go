package iproute

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/grant-he/lxd/shared"
)

//// Network interface operations ////
// TODO: convert these to use netlink instead of ip, per #7871

// InterfaceAddTAP creates TAP interface
func InterfaceAddTAP(name string) error {
	_, err := shared.RunCommand("ip", "tuntap", "add", "name", name, "mode", "tap")
	return err
}

// InterfaceAddRoute adds a route to device
func InterfaceAddRoute(route string, dev string, src string) error {
	cmd := []string{"ip", "route", "add", route, "dev", dev}
	if src != "" {
		cmd = append(cmd, []string{"src", src}...)
	}
	_, err := shared.RunCommand(cmd[0], cmd[1:]...)
	return err
}

// InterfaceRemove removes a network interface by name.
func InterfaceRemove(nic string) error {
	_, err := shared.RunCommand("ip", "link", "del", "dev", nic)
	return err
}

// InterfaceExists returns true if network interface exists.
func InterfaceExists(nic string) bool {
	return shared.PathExists(fmt.Sprintf("/sys/class/net/%s", nic))
}

// InterfaceRename changes the name of a network interface.
func InterfaceRename(nic string, name string) error {
	_, err := shared.RunCommand("ip", "link", "set", "dev", nic, "name", name)
	return err
}

// InterfaceBringUp enables a network interface by name (as by `ip link set <name> up`).
func InterfaceBringUp(nic string) error {
	_, err := shared.RunCommand("ip", "link", "set", "dev", nic, "up")
	return err
}

// InterfaceBringDown disables a network interface by name (as by `ip link set <name> down`).
func InterfaceBringDown(nic string) error {
	_, err := shared.RunCommand("ip", "link", "set", "dev", nic, "down")
	return err
}

// InterfaceSetNamespace sets the network namespace of a network interface.
func InterfaceSetNamespace(nic string, netNS string) error {
	_, err := shared.RunCommand("ip", "link", "set", "dev", nic, "netns", netNS)
	return err
}

// InterfaceSetMAC sets the hardware address of a network interface.
func InterfaceSetMAC(nic string, mac string) error {
	if mac != "" {
		_, err := shared.RunCommand("ip", "link", "set", "dev", nic, "address", mac)
		return err
	}

	return nil
}

// InterfaceSetMTU sets the MTU of a network interface.
func InterfaceSetMTU(nic string, mtu string) error {
	if mtu != "" {
		_, err := shared.RunCommand("ip", "link", "set", "dev", nic, "mtu", mtu)
		if err != nil {
			return errors.Wrapf(err, "Failed setting MTU %q on %q", mtu, nic)
		}
	}

	return nil
}

// InterfaceSetMaster sets the master interface of a network interface.
func InterfaceSetMaster(nic string, master string) error {
	_, err := shared.RunCommand("ip", "link", "set", "dev", nic, "master", master)
	return err
}

// InterfaceSetNoMaster detaches a network interface from its master.
func InterfaceSetNoMaster(nic string) error {
	_, err := shared.RunCommand("ip", "link", "set", "dev", nic, "nomaster")
	return err
}

// InterfaceFlushAddresses removes all addresses from a network interface.
func InterfaceFlushAddresses(nic string) error {
	_, err := shared.RunCommand("ip", "address", "flush", "dev", nic)
	return err
}

// IPv4AddAddress adds address to device
func IPv4AddAddress(name string, address string) error {
	_, err := shared.RunCommand("ip", "-4", "addr", "add", "dev", name, address)
	return err
}

// IPv6AddAddress adds address to device
func IPv6AddAddress(name string, address string) error {
	_, err := shared.RunCommand("ip", "-6", "addr", "add", "dev", name, address)
	return err
}

// IPv4AddRoute adds an IPv4 route to device
func IPv4AddRoute(route string, dev string, tableID string, rtProto string) error {
	cmd := []string{"ip", "-4", "route", "add", route, "dev", dev}
	if tableID != "" {
		cmd = append(cmd, []string{"table", tableID}...)
	}
	if rtProto != "" {
		cmd = append(cmd, []string{"proto", rtProto}...)
	}
	_, err := shared.RunCommand(cmd[0], cmd[1:]...)
	return err
}

// IPv6AddRoute adds an IPv6 route to device
func IPv6AddRoute(route string, dev string, tableID string, rtProto string) error {
	cmd := []string{"ip", "-6", "route", "add", route, "dev", dev}
	if tableID != "" {
		cmd = append(cmd, []string{"table", tableID}...)
	}
	if rtProto != "" {
		cmd = append(cmd, []string{"proto", rtProto}...)
	}
	_, err := shared.RunCommand(cmd[0], cmd[1:]...)
	return err
}

// IPv4DelRoute deletes an IPv4 route from device
func IPv4DelRoute(route string, dev string, tableID string) error {
	cmd := []string{"ip", "-6", "route", "delete", route, "dev", dev}
	if tableID != "" {
		cmd = append(cmd, []string{"table", tableID}...)
	}
	_, err := shared.RunCommand(cmd[0], cmd[1:]...)
	return err
}

// IPv6DelRoute deletes an IPv6 route from device
func IPv6DelRoute(route string, dev string, tableID string) error {
	cmd := []string{"ip", "-6", "route", "delete", route, "dev", dev}
	if tableID != "" {
		cmd = append(cmd, []string{"table", tableID}...)
	}
	_, err := shared.RunCommand(cmd[0], cmd[1:]...)
	return err
}

// IPv4FlushAddresses flushes all IPv4 address from device
func IPv4FlushAddresses(dev string, scope string) error {
	cmd := []string{"ip", "-4", "addr", "flush", "dev", dev}
	if scope != "" {
		cmd = append(cmd, []string{"scope", scope}...)
	}
	_, err := shared.RunCommand(cmd[0], cmd[1:]...)
	return err
}

// IPv6FlushAddresses flushes all IPv6 address from device
func IPv6FlushAddresses(dev string, scope string) error {
	cmd := []string{"ip", "-6", "addr", "flush", "dev", dev}
	if scope != "" {
		cmd = append(cmd, []string{"scope", scope}...)
	}
	_, err := shared.RunCommand(cmd[0], cmd[1:]...)
	return err
}

// IPv4FlushRoute flushes an IPv4 route from device
func IPv4FlushRoute(route string, dev string, rtProto string) error {
	cmd := []string{"ip", "-4", "route", "flush"}
	if route != "" {
		cmd = append(cmd, route)
	}

	cmd = append(cmd, []string{"dev", dev}...)

	if rtProto != "" {
		cmd = append(cmd, []string{"proto", rtProto}...)
	}
	_, err := shared.RunCommand(cmd[0], cmd[1:]...)
	return err
}

// IPv6FlushRoute flushes an IPv6 route from device
func IPv6FlushRoute(route string, dev string, rtProto string) error {
	cmd := []string{"ip", "-6", "route", "flush"}
	if route != "" {
		cmd = append(cmd, route)
	}

	cmd = append(cmd, []string{"dev", dev}...)

	if rtProto != "" {
		cmd = append(cmd, []string{"proto", rtProto}...)
	}
	_, err := shared.RunCommand(cmd[0], cmd[1:]...)
	return err
}

func IPv4ReplaceRoute(name string, routeFields []string) error {
	ipArgs := append([]string{"-4", "route", "replace", "dev", name, "proto", "boot"}, routeFields...)
	_, err := shared.RunCommand("ip", ipArgs...)
	return err
}

func IPv6ReplaceRoute(name string, routeFields []string) error {
	ipArgs := append([]string{"-6", "route", "replace", "dev", name, "proto", "boot"}, routeFields...)
	_, err := shared.RunCommand("ip", ipArgs...)
	return err
}

func IPLinkAddDummy(devName string, mtu string) error {
	_, err := shared.RunCommand("ip", "link", "add", "dev", devName, "mtu", mtu, "type", "dummy")
	return err
}

func IPLinkAddVeth(devName string, peerName string) error {
	_, err := shared.RunCommand("ip", "link", "add", "dev", devName, "type", "veth", "peer", "name", peerName)
	return err
}

func IPLinkAddBridge(devName string) error {
	_, err := shared.RunCommand("ip", "link", "add", "dev", devName, "type", "bridge")
	return err
}

func IPLinkAddMacvlan(devName string, parentName string, mode string) error {
	_, err := shared.RunCommand("ip", "link", "add", "dev", devName, "link", parentName, "type", "macvlan", "mode", mode)
	return err
}

func IPLinkAddMacvtap(devName string, parentName string, mode string) error {
	_, err := shared.RunCommand("ip", "link", "add", "dev", devName, "link", parentName, "type", "macvtap", "mode", mode)
	return err
}

func IPLinkAddVxlan(tunName string, vxlanID string, devName string, dstport string, devAddr string, fanMap string) error {
	_, err := shared.RunCommand("ip", "link", "add", tunName, "type", "vxlan", "id", vxlanID, "dev", devName, "dstport", dstport, "local", devAddr, "fan-map", fanMap)
	return err
}

// IPLinkAddVlan creates a VLAN interface named `vlanDevice` in VLAN ID `vlanId` and attaches it to `parent`.
func IPLinkAddVlan(parent string, vlanDevice string, vlanID string) error {
	_, err := shared.RunCommand("ip", "link", "add", "link", parent, "name", vlanDevice, "up", "type", "vlan", "id", vlanID)
	return err
}

func IPLinkChangeIpip(devName string, fanmap string) error {
	_, err := shared.RunCommand("ip", "link", "change", "dev", devName, "type", "ipip", "fan-map", fanmap)
	return err
}
