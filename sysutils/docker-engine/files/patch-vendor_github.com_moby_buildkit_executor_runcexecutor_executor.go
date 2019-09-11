--- vendor/github.com/moby/buildkit/executor/runcexecutor/executor.go.orig	2019-08-22 20:57:25 UTC
+++ vendor/github.com/moby/buildkit/executor/runcexecutor/executor.go
@@ -22,7 +22,6 @@ import (
 	"github.com/moby/buildkit/identity"
 	"github.com/moby/buildkit/solver/pb"
 	"github.com/moby/buildkit/util/network"
-	rootlessspecconv "github.com/moby/buildkit/util/rootless/specconv"
 	"github.com/pkg/errors"
 	"github.com/sirupsen/logrus"
 )
@@ -246,9 +245,7 @@ func (w *runcExecutor) Exec(ctx context.Context, meta 
 
 	spec.Process.OOMScoreAdj = w.oomScoreAdj
 	if w.rootless {
-		if err := rootlessspecconv.ToRootless(spec); err != nil {
-			return err
-		}
+		return errors.New("TODO: Rootless not implemented in FreeBSD!")
 	}
 
 	if err := json.NewEncoder(f).Encode(spec); err != nil {
