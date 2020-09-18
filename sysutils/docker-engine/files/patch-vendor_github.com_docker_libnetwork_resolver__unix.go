--- vendor/github.com/docker/libnetwork/resolver_unix.go.orig	2020-09-18 09:00:58 UTC
+++ vendor/github.com/docker/libnetwork/resolver_unix.go
@@ -1,4 +1,4 @@
-// +build !windows
+// +build !windows,!freebsd
 
 package libnetwork
 
