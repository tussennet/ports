--- daemon/update_freebsd.go.orig	2020-09-04 09:13:42 UTC
+++ daemon/update_freebsd.go
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
