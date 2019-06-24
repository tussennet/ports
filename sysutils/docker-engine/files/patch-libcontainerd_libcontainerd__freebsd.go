This is a copy from libcontainerd_linux.go

--- libcontainerd/libcontainerd_freebsd.go.orig	2019-06-24 09:47:19 UTC
+++ libcontainerd/libcontainerd_freebsd.go
@@ -0,0 +1,14 @@
+package libcontainerd // import "github.com/docker/docker/libcontainerd"
+
+import (
+	"context"
+
+	"github.com/containerd/containerd"
+	"github.com/docker/docker/libcontainerd/remote"
+	libcontainerdtypes "github.com/docker/docker/libcontainerd/types"
+)
+
+// NewClient creates a new libcontainerd client from a containerd client
+func NewClient(ctx context.Context, cli *containerd.Client, stateDir, ns string, b libcontainerdtypes.Backend) (libcontainerdtypes.Client, error) {
+	return remote.NewClient(ctx, cli, stateDir, ns, b)
+}
