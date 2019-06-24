--- libcontainerd/supervisor/utils_freebsd.go.orig	2019-06-24 18:38:41 UTC
+++ libcontainerd/supervisor/utils_freebsd.go
@@ -0,0 +1,11 @@
+package supervisor // import "github.com/docker/docker/libcontainerd/supervisor"
+
+import "syscall"
+
+// containerdSysProcAttr returns the SysProcAttr to use when exec'ing
+// containerd
+func containerdSysProcAttr() *syscall.SysProcAttr {
+	return &syscall.SysProcAttr{
+		Setsid:    true,
+	}
+}
