--- vendor/github.com/containerd/cgroups/memory.go.orig	2020-09-18 09:00:55 UTC
+++ vendor/github.com/containerd/cgroups/memory.go
@@ -208,7 +208,7 @@ func (m *memoryController) OOMEventFD(path string) (ui
 		return 0, err
 	}
 	defer f.Close()
-	fd, _, serr := unix.RawSyscall(unix.SYS_EVENTFD2, 0, unix.EFD_CLOEXEC, 0)
+	fd, _, serr := unix.RawSyscall(0, 0, 0, 0)//unix.RawSyscall(unix.SYS_EVENTFD2, 0, unix.EFD_CLOEXEC, 0)
 	if serr != 0 {
 		return 0, serr
 	}
