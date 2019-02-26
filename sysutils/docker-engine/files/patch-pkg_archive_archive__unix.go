--- pkg/archive/archive_unix.go.orig	2019-02-06 23:39:49 UTC
+++ pkg/archive/archive_unix.go
@@ -62,7 +62,7 @@ func getInodeFromStat(stat interface{}) (inode uint64,
 	s, ok := stat.(*syscall.Stat_t)
 
 	if ok {
-		inode = s.Ino
+		inode = uint64(s.Ino)
 	}
 
 	return
