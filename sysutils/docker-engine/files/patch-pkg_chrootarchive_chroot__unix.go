--- pkg/chrootarchive/chroot_unix.go.orig	2019-08-22 20:57:25 UTC
+++ pkg/chrootarchive/chroot_unix.go
@@ -14,3 +14,8 @@ func chroot(path string) error {
 func realChroot(path string) error {
 	return chroot(path)
 }
+
+
+func realChroot(path string) error {
+	return chroot(path)
+}
