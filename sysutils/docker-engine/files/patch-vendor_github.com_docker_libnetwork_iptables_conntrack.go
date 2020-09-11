--- vendor/github.com/docker/libnetwork/iptables/conntrack.go.orig	2020-09-04 14:54:57 UTC
+++ vendor/github.com/docker/libnetwork/iptables/conntrack.go
@@ -1,3 +1,5 @@
+// +build !freebsd
+
 package iptables
 
 import (
