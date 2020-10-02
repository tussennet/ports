--- specs/spec_freebsd.go.orig	2020-10-02 13:03:31 UTC
+++ specs/spec_freebsd.go
@@ -0,0 +1,10 @@
+package specs
+
+import ocs "github.com/opencontainers/runtime-spec/specs-go"
+
+type (
+	// ProcessSpec aliases the platform process specs
+	ProcessSpec ocs.Process
+	// Spec aliases the platform oci spec
+	Spec ocs.Spec
+)
