--- daemon/volumes_freebsd.go.orig	2020-09-04 09:13:42 UTC
+++ daemon/volumes_freebsd.go
@@ -0,0 +1,9 @@
+package daemon // import "github.com/docker/docker/daemon"
+
+import (
+	"github.com/docker/docker/api/types/mount"
+)
+
+func (daemon *Daemon) validateBindDaemonRoot(m mount.Mount) (bool, error) {
+	return false, nil
+}
