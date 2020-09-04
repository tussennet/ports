--- daemon/inspect.go.orig	2019-10-07 21:12:15 UTC
+++ daemon/inspect.go
@@ -22,7 +22,7 @@ import (
 func (daemon *Daemon) ContainerInspect(name string, size bool, version string) (interface{}, error) {
 	switch {
 	case versions.LessThan(version, "1.20"):
-		return daemon.containerInspectPre120(name)
+		return nil, errors.New("Port pre-1.20 not supported on freeBSD")
 	case versions.Equal(version, "1.20"):
 		return daemon.containerInspect120(name)
 	}
@@ -135,7 +135,7 @@ func (daemon *Daemon) getInspectData(container *contai
 	}
 
 	// We merge the Ulimits from hostConfig with daemon default
-	daemon.mergeUlimits(&hostConfig)
+	// daemon.mergeUlimits(&hostConfig)
 
 	var containerHealth *types.Health
 	if container.State.Health != nil {
