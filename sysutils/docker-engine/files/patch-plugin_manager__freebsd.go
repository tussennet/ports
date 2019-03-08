Fix build on FreeBSD by copying Windows stub

--- plugin/manager_freebsd.go.orig	2019-03-08 09:00:07 UTC
+++ plugin/manager_freebsd.go
@@ -0,0 +1,28 @@
+package plugin // import "github.com/docker/docker/plugin"
+
+import (
+	"fmt"
+
+	"github.com/docker/docker/plugin/v2"
+	specs "github.com/opencontainers/runtime-spec/specs-go"
+)
+
+func (pm *Manager) enable(p *v2.Plugin, c *controller, force bool) error {
+	return fmt.Errorf("Not implemented")
+}
+
+func (pm *Manager) initSpec(p *v2.Plugin) (*specs.Spec, error) {
+	return nil, fmt.Errorf("Not implemented")
+}
+
+func (pm *Manager) disable(p *v2.Plugin, c *controller) error {
+	return fmt.Errorf("Not implemented")
+}
+
+func (pm *Manager) restore(p *v2.Plugin, c *controller) error {
+	return fmt.Errorf("Not implemented")
+}
+
+// Shutdown plugins
+func (pm *Manager) Shutdown() {
+}
