--- daemon/container_operations_unix.go.orig	2020-08-28 10:10:01.856418000 +0200
+++ daemon/container_operations_unix.go	2020-08-28 10:10:16.078675000 +0200
@@ -144,9 +144,9 @@
 			}
 
 			shmproperty := "mode=1777,size=" + strconv.FormatInt(c.HostConfig.ShmSize, 10)
-			if err := unix.Mount("shm", shmPath, "tmpfs", uintptr(unix.MS_NOEXEC|unix.MS_NOSUID|unix.MS_NODEV), label.FormatMountLabel(shmproperty, c.GetMountLabel())); err != nil {
-				return fmt.Errorf("mounting shm tmpfs: %s", err)
-			}
+			// if err := unix.Mount("shm", shmPath, "tmpfs", uintptr(unix.MS_NOEXEC|unix.MS_NOSUID|unix.MS_NODEV), label.FormatMountLabel(shmproperty, c.GetMountLabel())); err != nil {
+			// 	return fmt.Errorf("mounting shm tmpfs: %s", err)
+			// }
 			if err := os.Chown(shmPath, rootIDs.UID, rootIDs.GID); err != nil {
 				return err
 			}
