--- daemon/container_freebsd.go.orig	2020-09-18 09:01:00 UTC
+++ daemon/container_freebsd.go
@@ -0,0 +1,9 @@
+package daemon
+
+import (
+	"github.com/docker/docker/container"
+)
+
+func (daemon *Daemon) saveApparmorConfig(container *container.Container) error {
+	return nil
+}
