--- libcontainerd/supervisor/remote_daemon_options_freebsd.go.orig	2020-08-28 12:28:37.910366000 +0200
+++ libcontainerd/supervisor/remote_daemon_options_freebsd.go	2020-08-28 12:28:27.814302000 +0200
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
