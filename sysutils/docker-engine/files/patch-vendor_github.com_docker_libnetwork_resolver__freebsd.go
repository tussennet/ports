--- vendor/github.com/docker/libnetwork/resolver_freebsd.go.orig	2020-09-04 09:13:43 UTC
+++ vendor/github.com/docker/libnetwork/resolver_freebsd.go
@@ -0,0 +1,12 @@
+package libnetwork
+
+import (
+	"fmt"
+)
+
+func init() {
+}
+
+func (r *resolver) setupIPTable() error {
+	return fmt.Errorf("IPTables not supported on FreeBSD")
+}
