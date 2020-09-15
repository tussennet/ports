--- pkg/parsers/operatingsystem/operatingsystem_freebsd.go.orig	2020-09-04 14:57:27 UTC
+++ pkg/parsers/operatingsystem/operatingsystem_freebsd.go
@@ -0,0 +1,31 @@
+// +build freebsd
+
+package operatingsystem
+
+import (
+	"errors"
+	"os/exec"
+	"syscall"
+)
+
+// GetOperatingSystem gets the name of the current operating system.
+func GetOperatingSystem() (string, error) {
+	cmd := exec.Command("uname", "-s")
+	osName, err := cmd.Output()
+	if err != nil {
+		return "", err
+	}
+	return string(osName), nil
+}
+
+// IsContainerized returns true if we are running inside a container.
+func IsContainerized() (bool, error) {
+	jailed, err := syscall.Sysctl("security.jail.jailed")
+	if err != nil {
+		return false, errors.New("Cannot detect if we are in a jail")
+	}
+	if jailed[0] == 1 {
+		return true, nil
+	}
+	return false, nil
+}
