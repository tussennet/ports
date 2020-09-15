--- daemon/container_operations.go.orig	2020-09-04 14:54:50 UTC
+++ daemon/container_operations.go
@@ -68,7 +68,7 @@ func (daemon *Daemon) buildSandboxOptions(container *c
 		sboxOptions = append(sboxOptions, libnetwork.OptionUseExternalKey())
 	}
 
-	if err = daemon.setupPathsAndSandboxOptions(container, &sboxOptions); err != nil {
+	if err = setupPathsAndSandboxOptions(container, &sboxOptions); err != nil {
 		return nil, err
 	}
 
@@ -618,9 +618,9 @@ func validateNetworkingConfig(n libnetwork.Network, ep
 		if hasUserDefinedIPAddress(epConfig.IPAMConfig) && !enableIPOnPredefinedNetwork() {
 			return runconfig.ErrUnsupportedNetworkAndIP
 		}
-		if len(epConfig.Aliases) > 0 && !serviceDiscoveryOnDefaultNetwork() {
-			return runconfig.ErrUnsupportedNetworkAndAlias
-		}
+		// if len(epConfig.Aliases) > 0 && !serviceDiscoveryOnDefaultNetwork() {
+		// 	return runconfig.ErrUnsupportedNetworkAndAlias
+		// }
 	}
 	if !hasUserDefinedIPAddress(epConfig.IPAMConfig) {
 		return nil
@@ -935,7 +935,7 @@ func (daemon *Daemon) initializeNetworking(container *
 			return err
 		}
 
-		err = daemon.initializeNetworkingPaths(container, nc)
+		err = initializeNetworkingPaths(container, nc)
 		if err != nil {
 			return err
 		}
