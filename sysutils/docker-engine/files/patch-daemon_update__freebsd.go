--- daemon/update_freebsd.go.orig	2020-09-18 09:01:00 UTC
+++ daemon/update_freebsd.go
@@ -0,0 +1,11 @@
+package daemon
+
+import (
+	"github.com/docker/docker/api/types/container"
+	libcontainerdtypes "github.com/docker/docker/libcontainerd/types"
+)
+
+func toContainerdResources(resources container.Resources) *libcontainerdtypes.Resources {
+	var r *libcontainerdtypes.Resources
+	return r
+}
