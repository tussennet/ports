--- daemon/volumes_freebsd.go.orig	2020-08-28 11:41:27.586718000 +0200
+++ daemon/volumes_freebsd.go	2020-08-28 11:41:15.346599000 +0200
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
