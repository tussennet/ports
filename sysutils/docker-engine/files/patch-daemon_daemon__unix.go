--- daemon/daemon_unix.go.orig	2020-10-23 14:06:17 UTC
+++ daemon/daemon_unix.go
@@ -7,6 +7,7 @@ import (
 	"context"
 	"fmt"
 	"io/ioutil"
+	"net"
 	"os"
 	"path/filepath"
 	"runtime"
@@ -37,8 +38,11 @@ import (
 	volumemounts "github.com/docker/docker/volume/mounts"
 	"github.com/docker/libnetwork"
 	nwconfig "github.com/docker/libnetwork/config"
+	"github.com/docker/libnetwork/drivers/freebsd/bridge"
 	"github.com/docker/libnetwork/netlabel"
+	"github.com/docker/libnetwork/netutils"
 	"github.com/docker/libnetwork/options"
+	lntypes "github.com/docker/libnetwork/types"
 
 	// "github.com/opencontainers/runc/libcontainer/cgroups"
 	rsystem "github.com/opencontainers/runc/libcontainer/system"
@@ -913,6 +917,134 @@ func driverOptions(config *config.Config) []nwconfig.O
 }
 
 func initBridgeDriver(controller libnetwork.NetworkController, config *config.Config) error {
+	bridgeName := bridge.DefaultBridgeName
+	if config.BridgeConfig.Iface != "" {
+		bridgeName = config.BridgeConfig.Iface
+	}
+	netOption := map[string]string{
+		bridge.BridgeName:         bridgeName,
+		bridge.DefaultBridge:      strconv.FormatBool(true),
+		netlabel.DriverMTU:        strconv.Itoa(config.Mtu),
+		bridge.EnableIPMasquerade: strconv.FormatBool(config.BridgeConfig.EnableIPMasq),
+		bridge.EnableICC:          strconv.FormatBool(config.BridgeConfig.InterContainerCommunication),
+	}
+
+	// --ip processing
+	if config.BridgeConfig.DefaultIP != nil {
+		netOption[bridge.DefaultBindingIP] = config.BridgeConfig.DefaultIP.String()
+	}
+
+	var (
+		ipamV4Conf *libnetwork.IpamConf
+		ipamV6Conf *libnetwork.IpamConf
+	)
+
+	ipamV4Conf = &libnetwork.IpamConf{AuxAddresses: make(map[string]string)}
+
+	nwList, nw6List, err := netutils.ElectInterfaceAddresses(bridgeName)
+	if err != nil {
+		return errors.Wrap(err, "list bridge addresses failed")
+	}
+
+	nw := nwList[0]
+	if len(nwList) > 1 && config.BridgeConfig.FixedCIDR != "" {
+		_, fCIDR, err := net.ParseCIDR(config.BridgeConfig.FixedCIDR)
+		if err != nil {
+			return errors.Wrap(err, "parse CIDR failed")
+		}
+		// Iterate through in case there are multiple addresses for the bridge
+		for _, entry := range nwList {
+			if fCIDR.Contains(entry.IP) {
+				nw = entry
+				break
+			}
+		}
+	}
+
+	ipamV4Conf.PreferredPool = lntypes.GetIPNetCanonical(nw).String()
+	hip, _ := lntypes.GetHostPartIP(nw.IP, nw.Mask)
+	if hip.IsGlobalUnicast() {
+		ipamV4Conf.Gateway = nw.IP.String()
+	}
+
+	if config.BridgeConfig.IP != "" {
+		ip, ipNet, err := net.ParseCIDR(config.BridgeConfig.IP)
+		if err != nil {
+			return err
+		}
+		ipamV4Conf.PreferredPool = ipNet.String()
+		ipamV4Conf.Gateway = ip.String()
+	} else if bridgeName == bridge.DefaultBridgeName && ipamV4Conf.PreferredPool != "" {
+		logrus.Infof("Default bridge (%s) is assigned with an IP address %s. Daemon option --bip can be used to set a preferred IP address", bridgeName, ipamV4Conf.PreferredPool)
+	}
+
+	if config.BridgeConfig.FixedCIDR != "" {
+		_, fCIDR, err := net.ParseCIDR(config.BridgeConfig.FixedCIDR)
+		if err != nil {
+			return err
+		}
+
+		ipamV4Conf.SubPool = fCIDR.String()
+	}
+
+	if config.BridgeConfig.DefaultGatewayIPv4 != nil {
+		ipamV4Conf.AuxAddresses["DefaultGatewayIPv4"] = config.BridgeConfig.DefaultGatewayIPv4.String()
+	}
+
+	var deferIPv6Alloc bool
+	if config.BridgeConfig.FixedCIDRv6 != "" {
+		_, fCIDRv6, err := net.ParseCIDR(config.BridgeConfig.FixedCIDRv6)
+		if err != nil {
+			return err
+		}
+
+		// In case user has specified the daemon flag --fixed-cidr-v6 and the passed network has
+		// at least 48 host bits, we need to guarantee the current behavior where the containers'
+		// IPv6 addresses will be constructed based on the containers' interface MAC address.
+		// We do so by telling libnetwork to defer the IPv6 address allocation for the endpoints
+		// on this network until after the driver has created the endpoint and returned the
+		// constructed address. Libnetwork will then reserve this address with the ipam driver.
+		ones, _ := fCIDRv6.Mask.Size()
+		deferIPv6Alloc = ones <= 80
+
+		if ipamV6Conf == nil {
+			ipamV6Conf = &libnetwork.IpamConf{AuxAddresses: make(map[string]string)}
+		}
+		ipamV6Conf.PreferredPool = fCIDRv6.String()
+
+		// In case the --fixed-cidr-v6 is specified and the current docker0 bridge IPv6
+		// address belongs to the same network, we need to inform libnetwork about it, so
+		// that it can be reserved with IPAM and it will not be given away to somebody else
+		for _, nw6 := range nw6List {
+			if fCIDRv6.Contains(nw6.IP) {
+				ipamV6Conf.Gateway = nw6.IP.String()
+				break
+			}
+		}
+	}
+
+	if config.BridgeConfig.DefaultGatewayIPv6 != nil {
+		if ipamV6Conf == nil {
+			ipamV6Conf = &libnetwork.IpamConf{AuxAddresses: make(map[string]string)}
+		}
+		ipamV6Conf.AuxAddresses["DefaultGatewayIPv6"] = config.BridgeConfig.DefaultGatewayIPv6.String()
+	}
+
+	v4Conf := []*libnetwork.IpamConf{ipamV4Conf}
+	v6Conf := []*libnetwork.IpamConf{}
+	if ipamV6Conf != nil {
+		v6Conf = append(v6Conf, ipamV6Conf)
+	}
+	// Initialize default network on "bridge" with the same name
+	_, err = controller.NewNetwork("bridge", "bridge", "",
+		libnetwork.NetworkOptionEnableIPv6(config.BridgeConfig.EnableIPv6),
+		libnetwork.NetworkOptionDriverOpts(netOption),
+		libnetwork.NetworkOptionIpam("default", "", v4Conf, v6Conf, nil),
+		libnetwork.NetworkOptionDeferIPv6Alloc(deferIPv6Alloc))
+	if err != nil {
+		return fmt.Errorf("Error creating default \"bridge\" network: %v", err)
+	}
+
 	return nil
 }
 
