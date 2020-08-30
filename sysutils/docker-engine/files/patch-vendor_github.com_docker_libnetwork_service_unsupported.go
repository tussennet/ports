--- vendor/github.com/docker/libnetwork/service_unsupported.go.orig	2019-10-07 23:12:15.000000000 +0200
+++ vendor/github.com/docker/libnetwork/service_unsupported.go	2020-08-21 16:24:55.067626000 +0200
@@ -1,5 +1,5 @@
-// +build !linux,!windows
-
+// +build !linux,!windows,!freebsd
+
 package libnetwork
 
 import (
