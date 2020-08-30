--- daemon/container_freebsd.go.orig	2020-08-28 11:28:40.369676000 +0200
+++ daemon/container_freebsd.go	2020-08-28 11:28:18.673781000 +0200
@@ -0,0 +1,9 @@
+package daemon // import "github.com/docker/docker/daemon"
+
+import (
+	"github.com/docker/docker/container"
+)
+
+func (daemon *Daemon) saveApparmorConfig(container *container.Container) error {
+	return nil
+}
