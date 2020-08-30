--- daemon/update_freebsd.go.orig	2020-08-28 11:34:43.179485000 +0200
+++ daemon/update_freebsd.go	2020-08-28 11:34:07.820028000 +0200
@@ -0,0 +1,10 @@
+package daemon // import "github.com/docker/docker/daemon"
+
+import (
+	"github.com/docker/docker/api/types/container"
+	libcontainerdtypes "github.com/docker/docker/libcontainerd/types"
+)
+
+func toContainerdResources(resources container.Resources) *libcontainerdtypes.Resources {
+	return nil
+}
