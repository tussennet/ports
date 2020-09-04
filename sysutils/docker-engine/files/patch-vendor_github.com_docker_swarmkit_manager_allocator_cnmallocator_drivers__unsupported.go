--- vendor/github.com/docker/swarmkit/manager/allocator/cnmallocator/drivers_unsupported.go.orig	2019-10-07 21:12:15 UTC
+++ vendor/github.com/docker/swarmkit/manager/allocator/cnmallocator/drivers_unsupported.go
@@ -6,7 +6,7 @@ import (
 	"github.com/docker/swarmkit/manager/allocator/networkallocator"
 )
 
-const initializers = nil
+var initializers = []initializer{}
 
 // PredefinedNetworks returns the list of predefined network structures
 func PredefinedNetworks() []networkallocator.PredefinedNetworkData {
