--- vendor/github.com/docker/libnetwork/netutils/utils_freebsd.go.orig	2020-09-04 14:54:57 UTC
+++ vendor/github.com/docker/libnetwork/netutils/utils_freebsd.go
@@ -1,7 +1,10 @@
 package netutils
 
 import (
+	"fmt"
 	"net"
+	"os/exec"
+	"strings"
 
 	"github.com/docker/libnetwork/types"
 )
@@ -19,5 +22,32 @@ func ElectInterfaceAddresses(name string) ([]*net.IPNe
 // FindAvailableNetwork returns a network from the passed list which does not
 // overlap with existing interfaces in the system
 func FindAvailableNetwork(list []*net.IPNet) (*net.IPNet, error) {
-	return nil, types.NotImplementedErrorf("not supported on freebsd")
+	for _, avail := range list {
+		cidr := strings.Split(avail.String(), "/")
+		ipitems := strings.Split(cidr[0], ".")
+		ip := ipitems[0] + "." +
+		      ipitems[1] + "." +
+		      ipitems[2] + "." + "1"
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
