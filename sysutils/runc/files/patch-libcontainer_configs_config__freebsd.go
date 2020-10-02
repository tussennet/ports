--- libcontainer/configs/config_freebsd.go.orig	2020-10-02 12:15:25 UTC
+++ libcontainer/configs/config_freebsd.go
@@ -0,0 +1,27 @@
+package configs
+
+// HostUID gets the translated uid for the process on host which could be
+// different when user namespaces are enabled.
+func (c Config) HostUID(containerId int) (int, error) {
+	// Return unchanged id.
+	return containerId, nil
+}
+
+// HostRootUID gets the root uid for the process on host which could be non-zero
+// when user namespaces are enabled.
+func (c Config) HostRootUID() (int, error) {
+	return c.HostUID(0)
+}
+
+// HostGID gets the translated gid for the process on host which could be
+// different when user namespaces are enabled.
+func (c Config) HostGID(containerId int) (int, error) {
+	// Return unchanged id.
+	return containerId, nil
+}
+
+// HostRootGID gets the root gid for the process on host which could be non-zero
+// when user namespaces are enabled.
+func (c Config) HostRootGID() (int, error) {
+	return c.HostGID(0)
+}
