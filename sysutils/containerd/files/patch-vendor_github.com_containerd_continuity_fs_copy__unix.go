--- vendor/github.com/containerd/continuity/fs/copy_unix.go.orig	2020-10-02 13:03:29 UTC
+++ vendor/github.com/containerd/continuity/fs/copy_unix.go
@@ -92,5 +92,5 @@ func copyDevice(dst string, fi os.FileInfo) error {
 	if !ok {
 		return errors.New("unsupported stat type")
 	}
-	return unix.Mknod(dst, uint32(fi.Mode()), int(st.Rdev))
+	return unix.Mknod(dst, uint32(fi.Mode()), st.Rdev)
 }
