--- pkg/archive/archive_unix.go.orig	2019-06-18 21:30:11 UTC
+++ pkg/archive/archive_unix.go
@@ -63,7 +63,7 @@ func getInodeFromStat(stat interface{}) (inode uint64,
 	s, ok := stat.(*syscall.Stat_t)
 
 	if ok {
-		inode = s.Ino
+		inode = uint64(s.Ino)
 	}
 
 	return
