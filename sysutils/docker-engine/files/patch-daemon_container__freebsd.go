--- daemon/container_freebsd.go.orig	2020-09-04 09:13:42 UTC
+++ daemon/container_freebsd.go
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
