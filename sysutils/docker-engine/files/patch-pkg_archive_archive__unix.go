--- pkg/archive/archive_unix.go.orig	2019-06-24 10:21:29 UTC
+++ pkg/archive/archive_unix.go
@@ -96,7 +96,7 @@ func handleTarTypeBlockCharFifo(hdr *tar.Header, path 
 		mode |= unix.S_IFIFO
 	}
 
-	return system.Mknod(path, mode, int(system.Mkdev(hdr.Devmajor, hdr.Devminor)))
+	return system.Mknod(path, mode, uint64(system.Mkdev(hdr.Devmajor, hdr.Devminor)))
 }
 
 func handleLChmod(hdr *tar.Header, path string, hdrInfo os.FileInfo) error {
