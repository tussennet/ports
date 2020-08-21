--- daemon/graphdriver/zfs/zfs.go.orig	2019-10-07 23:12:15.000000000 +0200
+++ daemon/graphdriver/zfs/zfs.go	2020-08-21 17:52:03.350611000 +0200
@@ -414,7 +414,7 @@
 
 	logger.Debugf(`unmount("%s")`, mountpoint)
 
-	if err := unix.Unmount(mountpoint, unix.MNT_DETACH); err != nil {
+	if err := unix.Unmount(mountpoint, 0); err != nil {
 		logger.Warnf("Failed to unmount %s mount %s: %v", id, mountpoint, err)
 	}
 	if err := unix.Rmdir(mountpoint); err != nil && !os.IsNotExist(err) {
