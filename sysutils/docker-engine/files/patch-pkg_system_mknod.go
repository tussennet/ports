--- pkg/system/mknod.go.orig	2019-06-18 21:30:11 UTC
+++ pkg/system/mknod.go
@@ -8,7 +8,7 @@ import (
 
 // Mknod creates a filesystem node (file, device special file or named pipe) named path
 // with attributes specified by mode and dev.
-func Mknod(path string, mode uint32, dev int) error {
+func Mknod(path string, mode uint32, dev uint64) error {
 	return unix.Mknod(path, mode, dev)
 }
 
