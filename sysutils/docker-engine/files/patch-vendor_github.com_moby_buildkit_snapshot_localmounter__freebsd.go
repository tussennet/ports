--- vendor/github.com/moby/buildkit/snapshot/localmounter_freebsd.go.orig	2019-03-06 19:37:34 UTC
+++ vendor/github.com/moby/buildkit/snapshot/localmounter_freebsd.go
@@ -0,0 +1,26 @@
+package snapshot
+
+import (
+	"os"
+
+	"github.com/containerd/containerd/mount"
+)
+
+func (lm *localMounter) Unmount() error {
+	lm.mu.Lock()
+	defer lm.mu.Unlock()
+
+	if lm.target != "" {
+		if err := mount.Unmount(lm.target, 0); err != nil {
+			return err
+		}
+		os.RemoveAll(lm.target)
+		lm.target = ""
+	}
+
+	if lm.mountable != nil {
+		return lm.mountable.Release()
+	}
+
+	return nil
+}
