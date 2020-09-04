--- libcontainerd/oom_unix.go.orig	2020-09-04 14:57:27 UTC
+++ libcontainerd/oom_unix.go
@@ -0,0 +1,7 @@
+// +build solaris,freebsd +build !linux
+
+package libcontainerd
+
+func setOOMScore(pid, score int) error {
+	return nil
+}
