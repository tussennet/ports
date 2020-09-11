--- vendor/github.com/vishvananda/netlink/filter.go.orig	2020-09-04 14:54:59 UTC
+++ vendor/github.com/vishvananda/netlink/filter.go
@@ -1,3 +1,5 @@
+// +build linux
+
 package netlink
 
 import (
