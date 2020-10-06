--- libcontainer/nsenter/nsenter.go.orig	2020-10-02 12:15:24 UTC
+++ libcontainer/nsenter/nsenter.go
@@ -1,4 +1,4 @@
-// +build linux,!gccgo
+// +build linux,!gccgo,!freebsd
 
 package nsenter
 
