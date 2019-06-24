--- vendor/github.com/docker/swarmkit/node/node.go.orig	2019-06-18 21:30:11 UTC
+++ vendor/github.com/docker/swarmkit/node/node.go
@@ -20,7 +20,6 @@ import (
 
 	"github.com/docker/docker/pkg/plugingetter"
 	"github.com/docker/go-metrics"
-	"github.com/docker/libnetwork/drivers/overlay/overlayutils"
 	"github.com/docker/swarmkit/agent"
 	"github.com/docker/swarmkit/agent/exec"
 	"github.com/docker/swarmkit/api"
@@ -273,11 +272,7 @@ func (n *Node) currentRole() api.NodeRole {
 
 // configVXLANUDPPort sets vxlan port in libnetwork
 func configVXLANUDPPort(ctx context.Context, vxlanUDPPort uint32) {
-	if err := overlayutils.ConfigVXLANUDPPort(vxlanUDPPort); err != nil {
-		log.G(ctx).WithError(err).Error("failed to configure VXLAN UDP port")
-		return
-	}
-	logrus.Infof("initialized VXLAN UDP port to %d ", vxlanUDPPort)
+	logrus.Infof("VXLAN UDP not supported on FreeBSD")
 }
 
 func (n *Node) run(ctx context.Context) (err error) {
