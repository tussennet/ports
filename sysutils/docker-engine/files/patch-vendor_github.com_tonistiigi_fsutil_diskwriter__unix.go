--- vendor/github.com/tonistiigi/fsutil/diskwriter_unix.go.orig	2019-10-07 21:12:15 UTC
+++ vendor/github.com/tonistiigi/fsutil/diskwriter_unix.go
@@ -45,7 +45,7 @@ func handleTarTypeBlockCharFifo(path string, stat *typ
 		mode |= syscall.S_IFBLK
 	}
 
-	if err := syscall.Mknod(path, mode, int(mkdev(stat.Devmajor, stat.Devminor))); err != nil {
+	if err := syscall.Mknod(path, mode, uint64(mkdev(stat.Devmajor, stat.Devminor))); err != nil {
 		return err
 	}
 	return nil
