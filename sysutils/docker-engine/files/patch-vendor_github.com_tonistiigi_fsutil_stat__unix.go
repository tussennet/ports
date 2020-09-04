--- vendor/github.com/tonistiigi/fsutil/stat_unix.go.orig	2019-10-07 21:12:15 UTC
+++ vendor/github.com/tonistiigi/fsutil/stat_unix.go
@@ -45,7 +45,7 @@ func setUnixOpt(fi os.FileInfo, stat *types.Stat, path
 			stat.Devminor = int64(minor(uint64(s.Rdev)))
 		}
 
-		ino := s.Ino
+		ino := uint64(s.Ino)
 		linked := false
 		if seenFiles != nil {
 			if s.Nlink > 1 {
