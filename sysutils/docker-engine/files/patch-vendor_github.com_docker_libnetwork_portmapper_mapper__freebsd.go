--- vendor/github.com/docker/libnetwork/portmapper/mapper_freebsd.go.orig	2020-09-04 09:13:43 UTC
+++ vendor/github.com/docker/libnetwork/portmapper/mapper_freebsd.go
@@ -0,0 +1,31 @@
+package portmapper
+
+import (
+	"net"
+	"sync"
+
+	"github.com/docker/libnetwork/portallocator"
+)
+
+// PortMapper manages the network address translation
+type PortMapper struct {
+	bridgeName string
+
+	// udp:ip:port
+	currentMappings map[string]*mapping
+	lock            sync.Mutex
+
+	proxyPath string
+
+	Allocator *portallocator.PortAllocator
+}
+
+// AppendForwardingTableEntry adds a port mapping to the forwarding table
+func (pm *PortMapper) AppendForwardingTableEntry(proto string, sourceIP net.IP, sourcePort int, containerIP string, containerPort int) error {
+	return nil
+}
+
+// DeleteForwardingTableEntry removes a port mapping from the forwarding table
+func (pm *PortMapper) DeleteForwardingTableEntry(proto string, sourceIP net.IP, sourcePort int, containerIP string, containerPort int) error {
+	return nil
+}
