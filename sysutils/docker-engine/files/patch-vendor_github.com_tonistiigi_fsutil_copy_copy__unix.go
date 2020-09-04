--- vendor/github.com/tonistiigi/fsutil/copy/copy_unix.go.orig	2019-10-07 21:12:15 UTC
+++ vendor/github.com/tonistiigi/fsutil/copy/copy_unix.go
@@ -3,6 +3,7 @@
 package fs
 
 import (
+	"io"
 	"os"
 	"syscall"
 
@@ -50,10 +51,33 @@ func (c *copier) copyFileInfo(fi os.FileInfo, name str
 	return nil
 }
 
+func copyFile(source, target string) error {
+	src, err := os.Open(source)
+	if err != nil {
+		return errors.Wrapf(err, "failed to open source %s", source)
+	}
+	defer src.Close()
+	tgt, err := os.Create(target)
+	if err != nil {
+		return errors.Wrapf(err, "failed to open target %s", target)
+	}
+	defer tgt.Close()
+
+	return copyFileContent(tgt, src)
+}
+
+func copyFileContent(dst, src *os.File) error {
+	_, err := io.Copy(dst, src)
+	if(err != nil) {
+		return err
+	}
+	return nil
+}
+
 func copyDevice(dst string, fi os.FileInfo) error {
 	st, ok := fi.Sys().(*syscall.Stat_t)
 	if !ok {
 		return errors.New("unsupported stat type")
 	}
-	return unix.Mknod(dst, uint32(fi.Mode()), int(st.Rdev))
+	return unix.Mknod(dst, uint32(fi.Mode()), st.Rdev)
 }
