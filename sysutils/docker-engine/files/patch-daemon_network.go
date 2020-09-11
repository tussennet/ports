--- daemon/network.go.orig	2020-09-04 14:54:50 UTC
+++ daemon/network.go
@@ -796,10 +796,10 @@ func buildCreateEndpointOptions(c *container.Container
 
 	defaultNetName := runconfig.DefaultDaemonNetworkMode().NetworkName()
 
-	if (!serviceDiscoveryOnDefaultNetwork() && n.Name() == defaultNetName) ||
-		c.NetworkSettings.IsAnonymousEndpoint {
-		createOptions = append(createOptions, libnetwork.CreateOptionAnonymous())
-	}
+	// if (!serviceDiscoveryOnDefaultNetwork() && n.Name() == defaultNetName) ||
+	// 	c.NetworkSettings.IsAnonymousEndpoint {
+	// 	createOptions = append(createOptions, libnetwork.CreateOptionAnonymous())
+	// }
 
 	if epConfig != nil {
 		ipam := epConfig.IPAMConfig
