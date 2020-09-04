--- daemon/container_operations_unix.go.orig	2019-10-07 21:12:15 UTC
+++ daemon/container_operations_unix.go
@@ -143,10 +143,10 @@ func (daemon *Daemon) setupIpcDirs(c *container.Contai
 				return err
 			}
 
-			shmproperty := "mode=1777,size=" + strconv.FormatInt(c.HostConfig.ShmSize, 10)
-			if err := unix.Mount("shm", shmPath, "tmpfs", uintptr(unix.MS_NOEXEC|unix.MS_NOSUID|unix.MS_NODEV), label.FormatMountLabel(shmproperty, c.GetMountLabel())); err != nil {
-				return fmt.Errorf("mounting shm tmpfs: %s", err)
-			}
+			// shmproperty := "mode=1777,size=" + strconv.FormatInt(c.HostConfig.ShmSize, 10)
+			// if err := unix.Mount("shm", shmPath, "tmpfs", uintptr(unix.MS_NOEXEC|unix.MS_NOSUID|unix.MS_NODEV), label.FormatMountLabel(shmproperty, c.GetMountLabel())); err != nil {
+			// 	return fmt.Errorf("mounting shm tmpfs: %s", err)
+			// }
 			if err := os.Chown(shmPath, rootIDs.UID, rootIDs.GID); err != nil {
 				return err
 			}
