--- vendor/github.com/moby/buildkit/snapshot/localmounter_unix.go.orig	2019-10-07 21:12:15 UTC
+++ vendor/github.com/moby/buildkit/snapshot/localmounter_unix.go
@@ -1,4 +1,4 @@
-// +build !windows
+// +build !windows,!freebsd
 
 package snapshot
 
