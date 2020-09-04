--- daemon/container_operations_freebsd.go.orig	2020-09-04 14:57:27 UTC
+++ daemon/container_operations_freebsd.go
@@ -0,0 +1,45 @@
+package daemon // import "github.com/docker/docker/daemon"
+
+import (
+	"github.com/docker/docker/container"
+	"github.com/docker/docker/runconfig"
+	"github.com/docker/libnetwork"
+)
+
+func (daemon *Daemon) setupLinkedContainers(container *container.Container) ([]string, error) {
+	return nil, nil
+}
+
+func (daemon *Daemon) setupIpcDirs(container *container.Container) error {
+	return nil
+}
+
+func killProcessDirectly(container *container.Container) error {
+	return nil
+}
+
+func detachMounted(path string) error {
+	return nil
+}
+
+func isLinkable(child *container.Container) bool {
+	// A container is linkable only if it belongs to the default network
+	_, ok := child.NetworkSettings.Networks[runconfig.DefaultDaemonNetworkMode().NetworkName()]
+	return ok
+}
+
+func enableIPOnPredefinedNetwork() bool {
+	return false
+}
+
+func (daemon *Daemon) isNetworkHotPluggable() bool {
+	return false
+}
+
+func setupPathsAndSandboxOptions(container *container.Container, sboxOptions *[]libnetwork.SandboxOption) error {
+	return nil
+}
+
+func initializeNetworkingPaths(container *container.Container, nc *container.Container) error {
+	return nil
+}
