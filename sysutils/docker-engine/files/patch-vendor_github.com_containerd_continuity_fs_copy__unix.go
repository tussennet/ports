--- vendor/github.com/containerd/continuity/fs/copy_unix.go.orig	2019-06-18 21:30:11 UTC
+++ vendor/github.com/containerd/continuity/fs/copy_unix.go
@@ -108,5 +108,5 @@ func copyDevice(dst string, fi os.FileInfo) error {
 	if !ok {
 		return errors.New("unsupported stat type")
 	}
-	return unix.Mknod(dst, uint32(fi.Mode()), int(st.Rdev))
+	return unix.Mknod(dst, uint32(fi.Mode()), st.Rdev)
 }
