--- daemon/oci_freebsd.go.orig	2020-09-18 09:01:00 UTC
+++ daemon/oci_freebsd.go
@@ -0,0 +1,74 @@
+package daemon
+
+import (
+	"fmt"
+	"sort"
+
+	containertypes "github.com/docker/docker/api/types/container"
+	"github.com/docker/docker/container"
+	"github.com/docker/docker/oci"
+	"github.com/opencontainers/runtime-spec/specs-go"
+)
+
+func setResources(s *specs.Spec, r containertypes.Resources) error {
+	return nil
+}
+
+func setUser(s *specs.Spec, c *container.Container) error {
+	return nil
+}
+
+func getUser(c *container.Container, username string) (uint32, uint32, []uint32, error) {
+	return 0, 0, nil, nil
+}
+
+// mergeUlimits merge the Ulimits from HostConfig with daemon defaults, and update HostConfig
+// It will do nothing on non-Linux platform
+func (daemon *Daemon) mergeUlimits(c *containertypes.HostConfig) {
+	return
+}
+
+func (daemon *Daemon) createSpec(c *container.Container) (*specs.Spec, error) {
+	s := oci.DefaultSpec()
+	if err := daemon.populateCommonSpec(&s, c); err != nil {
+		return nil, err
+	}
+
+	if err := setResources(&s, c.HostConfig.Resources); err != nil {
+		return nil, fmt.Errorf("runtime spec resources: %v", err)
+	}
+
+	if err := setUser(&s, c); err != nil {
+		return nil, fmt.Errorf("spec user: %v", err)
+	}
+
+	if err := daemon.setNetworkInterface(&s, c); err != nil {
+		return nil, err
+	}
+
+	if err := daemon.setupIpcDirs(c); err != nil {
+		return nil, err
+	}
+
+	ms, err := daemon.setupMounts(c)
+	if err != nil {
+		return nil, err
+	}
+	ms = append(ms, c.IpcMounts()...)
+	tmpfsMounts, err := c.TmpfsMounts()
+	if err != nil {
+		return nil, err
+	}
+	ms = append(ms, tmpfsMounts...)
+	sort.Sort(mounts(ms))
+
+	return (*specs.Spec)(&s), nil
+}
+
+func (daemon *Daemon) setNetworkInterface(s *specs.Spec, c *container.Container) error {
+	return nil
+}
+
+func (daemon *Daemon) populateCommonSpec(s *specs.Spec, c *container.Container) error {
+	return nil
+}
