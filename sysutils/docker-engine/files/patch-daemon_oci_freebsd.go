--- daemon/oci_freebsd.go.orig	2020-08-28 10:57:58.901293000 +0200
+++ daemon/oci_freebsd.go	2020-08-28 10:57:33.391424000 +0200
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
