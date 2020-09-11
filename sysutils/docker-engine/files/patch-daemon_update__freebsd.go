--- daemon/update_freebsd.go.orig	2020-09-04 14:57:27 UTC
+++ daemon/update_freebsd.go
@@ -0,0 +1,12 @@
+package daemon // import "github.com/docker/docker/daemon"
+
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
