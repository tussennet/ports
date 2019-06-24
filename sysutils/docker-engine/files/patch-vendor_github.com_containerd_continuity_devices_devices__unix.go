--- vendor/github.com/containerd/continuity/devices/devices_unix.go.orig	2019-06-18 21:30:11 UTC
+++ vendor/github.com/containerd/continuity/devices/devices_unix.go
@@ -55,7 +55,7 @@ func Mknod(p string, mode os.FileMode, maj, min int) e
 		m |= unix.S_IFIFO
 	}
 
-	return unix.Mknod(p, m, int(dev))
+	return unix.Mknod(p, m, dev)
 }
 
 // syscallMode returns the syscall-specific mode bits from Go's portable mode bits.
