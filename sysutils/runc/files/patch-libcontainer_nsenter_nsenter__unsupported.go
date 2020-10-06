--- libcontainer/nsenter/nsenter_unsupported.go.orig	2020-10-02 12:15:24 UTC
+++ libcontainer/nsenter/nsenter_unsupported.go
@@ -1,3 +1,3 @@
-// +build !linux !cgo
+// +build !linux !cgo !freebsd
 
 package nsenter
