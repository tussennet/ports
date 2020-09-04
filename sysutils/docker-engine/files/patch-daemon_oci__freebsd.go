--- daemon/oci_freebsd.go.orig	2020-09-04 09:13:42 UTC
+++ daemon/oci_freebsd.go
@@ -0,0 +1,12 @@
+package daemon // import "github.com/docker/docker/daemon"
+
+import (
+	"errors"
+
+	"github.com/docker/docker/container"
+	"github.com/opencontainers/runtime-spec/specs-go"
+)
+
+func (daemon *Daemon) createSpec(c *container.Container) (retSpec *specs.Spec, err error) {
+	return nil, errors.New("not implemented")
+}
