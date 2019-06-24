--- vendor/github.com/moby/buildkit/executor/runcexecutor/executor.go.orig	2019-06-18 21:30:11 UTC
+++ vendor/github.com/moby/buildkit/executor/runcexecutor/executor.go
@@ -24,7 +24,6 @@ import (
 	"github.com/moby/buildkit/identity"
 	"github.com/moby/buildkit/solver/pb"
 	"github.com/moby/buildkit/util/network"
-	rootlessspecconv "github.com/moby/buildkit/util/rootless/specconv"
 	specs "github.com/opencontainers/runtime-spec/specs-go"
 	"github.com/pkg/errors"
 	"github.com/sirupsen/logrus"
@@ -241,9 +240,7 @@ func (w *runcExecutor) Exec(ctx context.Context, meta 
 		return err
 	}
 	if w.rootless {
-		if err := rootlessspecconv.ToRootless(spec); err != nil {
-			return err
-		}
+		return errors.New("TODO: Rootless not implemented in FreeBSD!")
 	}
 
 	if err := json.NewEncoder(f).Encode(spec); err != nil {
