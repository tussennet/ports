--- libcontainerd/supervisor/remote_daemon_options_freebsd.go.orig	2020-09-04 09:13:43 UTC
+++ libcontainerd/supervisor/remote_daemon_options_freebsd.go
@@ -0,0 +1,9 @@
+package supervisor // import "github.com/docker/docker/libcontainerd/supervisor"
+
+// WithOOMScore defines the oom_score_adj to set for the containerd process.
+func WithOOMScore(score int) DaemonOpt {
+	return func(r *remote) error {
+		r.OOMScore = score
+		return nil
+	}
+}
