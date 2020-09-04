--- vendor/github.com/containerd/continuity/devices/devices_unix.go.orig	2019-10-07 21:12:15 UTC
+++ vendor/github.com/containerd/continuity/devices/devices_unix.go
@@ -55,7 +55,7 @@ func Mknod(p string, mode os.FileMode, maj, min int) e
 		m |= unix.S_IFIFO
 	}
 
-	return unix.Mknod(p, m, int(dev))
+	return unix.Mknod(p, m, dev)
 }
 
 // syscallMode returns the syscall-specific mode bits from Go's portable mode bits.
