--- daemon/exec_freebsd.go.orig	2020-08-28 10:21:49.828328000 +0200
+++ daemon/exec_freebsd.go	2020-08-28 10:21:42.012410000 +0200
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
