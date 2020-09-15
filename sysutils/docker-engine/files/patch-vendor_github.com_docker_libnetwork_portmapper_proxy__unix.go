--- vendor/github.com/docker/libnetwork/portmapper/proxy_unix.go.orig	2020-09-04 14:57:27 UTC
+++ vendor/github.com/docker/libnetwork/portmapper/proxy_unix.go
@@ -0,0 +1,36 @@
+// +build solaris,freebsd +build !linux
+
+package portmapper
+
+import (
+	"net"
+	"os/exec"
+	"strconv"
+)
+
+func newProxyCommand(proto string, hostIP net.IP, hostPort int, containerIP net.IP, containerPort int, proxyPath string) (userlandProxy, error) {
+	path := proxyPath
+	if proxyPath == "" {
+		cmd, err := exec.LookPath(userlandProxyCommandName)
+		if err != nil {
+			return nil, err
+		}
+		path = cmd
+	}
+
+	args := []string{
+		path,
+		"-proto", proto,
+		"-host-ip", hostIP.String(),
+		"-host-port", strconv.Itoa(hostPort),
+		"-container-ip", containerIP.String(),
+		"-container-port", strconv.Itoa(containerPort),
+	}
+
+	return &proxyCommand{
+		cmd: &exec.Cmd{
+			Path: path,
+			Args: args,
+		},
+	}, nil
+}
