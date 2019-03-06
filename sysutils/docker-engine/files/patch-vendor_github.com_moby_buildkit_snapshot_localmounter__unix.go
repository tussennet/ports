--- vendor/github.com/moby/buildkit/snapshot/localmounter_unix.go.orig	2019-02-26 00:29:56 UTC
+++ vendor/github.com/moby/buildkit/snapshot/localmounter_unix.go
@@ -1,4 +1,4 @@
-// +build !windows
+// +build !windows,!freebsd
 
 package snapshot
 
