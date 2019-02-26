--- pkg/archive/changes_unix.go.orig	2019-02-06 23:39:49 UTC
+++ pkg/archive/changes_unix.go
@@ -29,7 +29,7 @@ func (info *FileInfo) isDir() bool {
 }
 
 func getIno(fi os.FileInfo) uint64 {
-	return fi.Sys().(*syscall.Stat_t).Ino
+	return uint64(fi.Sys().(*syscall.Stat_t).Ino)
 }
 
 func hasHardlinks(fi os.FileInfo) bool {
