--- vendor/github.com/docker/libnetwork/netutils/utils_freebsd.go.orig	2020-10-23 18:37:21 UTC
+++ vendor/github.com/docker/libnetwork/netutils/utils_freebsd.go
@@ -1,9 +1,15 @@
 package netutils
 
 import (
+	"fmt"
 	"net"
+	"os/exec"
+	"strings"
 
-	"github.com/docker/libnetwork/types"
+	"github.com/docker/libnetwork/ipamutils"
+	"github.com/docker/libnetwork/ns"
+	"github.com/docker/libnetwork/osl"
+	"github.com/pkg/errors"
 )
 
 // ElectInterfaceAddresses looks for an interface on the OS with the specified name
@@ -13,11 +19,74 @@ import (
 // list the first IPv4 address which does not conflict with other
 // interfaces on the system.
 func ElectInterfaceAddresses(name string) ([]*net.IPNet, []*net.IPNet, error) {
-	return nil, nil, types.NotImplementedErrorf("not supported on freebsd")
+	var (
+		v4Nets []*net.IPNet
+		v6Nets []*net.IPNet
+	)
+
+	defer osl.InitOSContext()()
+
+	link, _ := ns.NlHandle().LinkByName(name)
+	// disabled on freebsd for now
+	// if link != nil {
+	// 	v4addr, err := ns.NlHandle().AddrList(link, netlink.FAMILY_V4)
+	// 	if err != nil {
+	// 		return nil, nil, err
+	// 	}
+	// 	v6addr, err := ns.NlHandle().AddrList(link, netlink.FAMILY_V6)
+	// 	if err != nil {
+	// 		return nil, nil, err
+	// 	}
+	// 	for _, nlAddr := range v4addr {
+	// 		v4Nets = append(v4Nets, nlAddr.IPNet)
+	// 	}
+	// 	for _, nlAddr := range v6addr {
+	// 		v6Nets = append(v6Nets, nlAddr.IPNet)
+	// 	}
+	// }
+
+	if link == nil || len(v4Nets) == 0 {
+		// Choose from predefined local scope networks
+		v4Net, err := FindAvailableNetwork(ipamutils.PredefinedLocalScopeDefaultNetworks)
+		if err != nil {
+			return nil, nil, errors.Wrapf(err, "PredefinedLocalScopeDefaultNetworks List: %+v",
+				ipamutils.PredefinedLocalScopeDefaultNetworks)
+		}
+		v4Nets = append(v4Nets, v4Net)
+	}
+
+	return v4Nets, v6Nets, nil
 }
 
 // FindAvailableNetwork returns a network from the passed list which does not
 // overlap with existing interfaces in the system
 func FindAvailableNetwork(list []*net.IPNet) (*net.IPNet, error) {
-	return nil, types.NotImplementedErrorf("not supported on freebsd")
+	for _, avail := range list {
+		cidr := strings.Split(avail.String(), "/")
+		ipitems := strings.Split(cidr[0], ".")
+		ip := ipitems[0] + "." +
+			ipitems[1] + "." +
+			ipitems[2] + "." + "1"
+
+		out, err := exec.Command("/sbin/route", "get", ip).Output()
+		if err != nil {
+			fmt.Println("failed to run route get command")
+			return nil, err
+		}
+		lines := strings.Split(string(out), "\n")
+		for _, l := range lines {
+			s := strings.Split(string(l), ":")
+			if len(s) == 2 {
+				k, v := s[0], strings.TrimSpace(s[1])
+				if k == "destination" {
+					if v == "default" {
+						return avail, nil
+					}
+					break
+				}
+			}
+		}
+	}
+	return nil, fmt.Errorf("no available network")
+	//types.NotImplementedErrorf("not supported on freebsd")
 }
