--- libcontainer/criu_opts_freebsd.go.orig	2020-10-02 12:15:25 UTC
+++ libcontainer/criu_opts_freebsd.go
@@ -0,0 +1,6 @@
+package libcontainer
+
+// TODO Windows: This can ultimately be entirely factored out as criu is
+// a Unix concept not relevant on Windows.
+type CriuOpts struct {
+}
