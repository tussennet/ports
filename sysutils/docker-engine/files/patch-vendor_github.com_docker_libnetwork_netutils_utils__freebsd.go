--- vendor/github.com/docker/libnetwork/netutils/utils_freebsd.go.orig	2020-10-23 14:06:21 UTC
+++ vendor/github.com/docker/libnetwork/netutils/utils_freebsd.go
@@ -6,7 +6,10 @@ import (
 	"os/exec"
 	"strings"
 
-	"github.com/docker/libnetwork/types"
+	"github.com/docker/libnetwork/ipamutils"
+	"github.com/docker/libnetwork/ns"
+	"github.com/docker/libnetwork/osl"
+	"github.com/pkg/errors"
 )
 
 // ElectInterfaceAddresses looks for an interface on the OS with the specified name
@@ -16,7 +19,43 @@ import (
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
@@ -26,8 +65,8 @@ func FindAvailableNetwork(list []*net.IPNet) (*net.IPN
 		cidr := strings.Split(avail.String(), "/")
 		ipitems := strings.Split(cidr[0], ".")
 		ip := ipitems[0] + "." +
-		      ipitems[1] + "." +
-		      ipitems[2] + "." + "1"
+			ipitems[1] + "." +
+			ipitems[2] + "." + "1"
 
 		out, err := exec.Command("/sbin/route", "get", ip).Output()
 		if err != nil {
