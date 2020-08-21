--- vendor/github.com/containerd/cgroups/memory.go.orig	2019-10-07 23:12:15.000000000 +0200
+++ vendor/github.com/containerd/cgroups/memory.go	2020-08-21 18:41:07.332697000 +0200
@@ -176,7 +176,7 @@
 		return 0, err
 	}
 	defer f.Close()
-	fd, _, serr := unix.RawSyscall(unix.SYS_EVENTFD2, 0, unix.EFD_CLOEXEC, 0)
+	fd, _, serr := unix.RawSyscall(0, 0, 0, 0)//unix.RawSyscall(unix.SYS_EVENTFD2, 0, unix.EFD_CLOEXEC, 0)
 	if serr != 0 {
 		return 0, serr
 	}
