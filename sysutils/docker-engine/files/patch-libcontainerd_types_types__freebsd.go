--- libcontainerd/types/types_freebsd.go.orig	2019-06-24 11:36:48 UTC
+++ libcontainerd/types/types_freebsd.go
@@ -0,0 +1,24 @@
+package types // import "github.com/docker/docker/libcontainerd/types"
+
+import (
+	"time"
+
+	"github.com/opencontainers/runtime-spec/specs-go"
+)
+
+// Summary is not used on FreeBSD
+type Summary struct{}
+
+// Stats holds metrics properties as returned by containerd
+type Stats struct {}
+
+// InterfaceToStats returns a stats object from the platform-specific interface.
+func InterfaceToStats(read time.Time, v interface{}) *Stats {
+	return &Stats{}
+}
+
+// Resources defines updatable container resource values. TODO: it must match containerd upcoming API
+type Resources specs.LinuxResources
+
+// Checkpoints contains the details of a checkpoint
+type Checkpoints struct{}
