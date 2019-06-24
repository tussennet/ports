--- pkg/chrootarchive/chroot_unix.go.orig	2019-06-24 11:59:08 UTC
+++ pkg/chrootarchive/chroot_unix.go
@@ -10,3 +10,8 @@ func chroot(path string) error {
 	}
 	return unix.Chdir("/")
 }
+
+func realChroot(path string) error {
+	return chroot(path)
+}
+
