--- osutils/pdeathsig_freebsd.go.orig	2020-10-02 13:03:31 UTC
+++ osutils/pdeathsig_freebsd.go
@@ -0,0 +1,15 @@
+// +build freebsd
+
+package osutils
+
+import (
+	"syscall"
+)
+
+// SetPDeathSig sets the parent death signal to SIGKILL so that if the
+// shim dies the container process also dies.
+func SetPDeathSig() *syscall.SysProcAttr {
+	return &syscall.SysProcAttr{
+	//Pdeathsig: syscall.SIGKILL,
+	}
+}
