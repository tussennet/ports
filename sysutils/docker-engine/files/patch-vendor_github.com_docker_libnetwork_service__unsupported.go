--- vendor/github.com/docker/libnetwork/service_unsupported.go.orig	2019-10-07 21:12:15 UTC
+++ vendor/github.com/docker/libnetwork/service_unsupported.go
@@ -1,4 +1,4 @@
-// +build !linux,!windows
+// +build !linux,!windows,!freebsd
 
 package libnetwork
 
