--- vendor/github.com/docker/libnetwork/drivers_freebsd.go.orig	2020-09-04 14:54:57 UTC
+++ vendor/github.com/docker/libnetwork/drivers_freebsd.go
@@ -1,12 +1,14 @@
 package libnetwork
 
 import (
+	"github.com/docker/libnetwork/drivers/freebsd/bridge"
 	"github.com/docker/libnetwork/drivers/null"
 	"github.com/docker/libnetwork/drivers/remote"
 )
 
 func getInitializers(experimental bool) []initializer {
 	return []initializer{
+		{bridge.Init, "bridge"},
 		{null.Init, "null"},
 		{remote.Init, "remote"},
 	}
