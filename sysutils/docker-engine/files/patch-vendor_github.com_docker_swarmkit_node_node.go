--- vendor/github.com/docker/swarmkit/node/node.go.orig	2019-10-07 21:12:15 UTC
+++ vendor/github.com/docker/swarmkit/node/node.go
@@ -21,7 +21,6 @@ import (
 
 	"github.com/docker/docker/pkg/plugingetter"
 	"github.com/docker/go-metrics"
-	"github.com/docker/libnetwork/drivers/overlay/overlayutils"
 	"github.com/docker/swarmkit/agent"
 	"github.com/docker/swarmkit/agent/exec"
 	"github.com/docker/swarmkit/api"
@@ -274,11 +273,7 @@ func (n *Node) currentRole() api.NodeRole {
 
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
