--- daemon/graphdriver/zfs/zfs.go.orig	2019-10-07 21:12:15 UTC
+++ daemon/graphdriver/zfs/zfs.go
@@ -414,7 +414,7 @@ func (d *Driver) Put(id string) error {
 
 	logger.Debugf(`unmount("%s")`, mountpoint)
 
-	if err := unix.Unmount(mountpoint, unix.MNT_DETACH); err != nil {
+	if err := unix.Unmount(mountpoint, 0); err != nil {
 		logger.Warnf("Failed to unmount %s mount %s: %v", id, mountpoint, err)
 	}
 	if err := unix.Rmdir(mountpoint); err != nil && !os.IsNotExist(err) {
