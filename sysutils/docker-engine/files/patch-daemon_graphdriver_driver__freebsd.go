Fix build error on FreeBSD:

daemon/graphdriver/driver_freebsd.go:17:38: cannot use &buf (type *unix.Statfs_t) as type *syscall.Statfs_t in argument to syscall.Statfs

--- daemon/graphdriver/driver_freebsd.go.orig	2019-02-26 00:29:56 UTC
+++ daemon/graphdriver/driver_freebsd.go
@@ -1,8 +1,6 @@
 package graphdriver // import "github.com/docker/docker/daemon/graphdriver"
 
 import (
-	"syscall"
-
 	"golang.org/x/sys/unix"
 )
 
@@ -14,7 +12,7 @@ var (
 // Mounted checks if the given path is mounted as the fs type
 func Mounted(fsType FsMagic, mountPath string) (bool, error) {
 	var buf unix.Statfs_t
-	if err := syscall.Statfs(mountPath, &buf); err != nil {
+	if err := unix.Statfs(mountPath, &buf); err != nil {
 		return false, err
 	}
 	return FsMagic(buf.Type) == fsType, nil
