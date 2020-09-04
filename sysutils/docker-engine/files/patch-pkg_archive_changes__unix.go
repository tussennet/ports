--- pkg/archive/changes_unix.go.orig	2019-10-07 21:12:15 UTC
+++ pkg/archive/changes_unix.go
@@ -35,7 +35,7 @@ func (info *FileInfo) isDir() bool {
 }
 
 func getIno(fi os.FileInfo) uint64 {
-	return fi.Sys().(*syscall.Stat_t).Ino
+	return uint64(fi.Sys().(*syscall.Stat_t).Ino)
 }
 
 func hasHardlinks(fi os.FileInfo) bool {
