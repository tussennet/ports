--- archive/tar_unix.go.orig	2020-10-02 13:03:27 UTC
+++ archive/tar_unix.go
@@ -122,7 +122,7 @@ func handleTarTypeBlockCharFifo(hdr *tar.Header, path 
 		mode |= unix.S_IFIFO
 	}
 
-	return unix.Mknod(path, mode, int(unix.Mkdev(uint32(hdr.Devmajor), uint32(hdr.Devminor))))
+	return unix.Mknod(path, mode, unix.Mkdev(uint32(hdr.Devmajor), uint32(hdr.Devminor)))
 }
 
 func handleLChmod(hdr *tar.Header, path string, hdrInfo os.FileInfo) error {
