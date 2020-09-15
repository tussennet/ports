--- daemon/exec_freebsd.go.orig	2020-09-04 09:13:42 UTC
+++ daemon/exec_freebsd.go
@@ -0,0 +1,11 @@
+package daemon // import "github.com/docker/docker/daemon"
+
+import (
+	"github.com/docker/docker/container"
+	"github.com/docker/docker/daemon/exec"
+	specs "github.com/opencontainers/runtime-spec/specs-go"
+)
+
+func (daemon *Daemon) execSetPlatformOpt(c *container.Container, ec *exec.Config, p *specs.Process) error {
+	return nil
+}
