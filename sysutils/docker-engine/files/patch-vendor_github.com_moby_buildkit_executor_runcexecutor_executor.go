--- vendor/github.com/moby/buildkit/executor/runcexecutor/executor.go.orig	2019-02-06 23:39:49 UTC
+++ vendor/github.com/moby/buildkit/executor/runcexecutor/executor.go
@@ -13,7 +13,6 @@ import (
 	"syscall"
 	"time"
 
-	"github.com/containerd/containerd/contrib/seccomp"
 	"github.com/containerd/containerd/mount"
 	containerdoci "github.com/containerd/containerd/oci"
 	"github.com/containerd/continuity/fs"
@@ -24,7 +23,6 @@ import (
 	"github.com/moby/buildkit/identity"
 	"github.com/moby/buildkit/solver/pb"
 	"github.com/moby/buildkit/util/network"
-	rootlessspecconv "github.com/moby/buildkit/util/rootless/specconv"
 	"github.com/moby/buildkit/util/system"
 	specs "github.com/opencontainers/runtime-spec/specs-go"
 	"github.com/pkg/errors"
@@ -177,7 +175,7 @@ func (w *runcExecutor) Exec(ctx context.Context, meta 
 
 	opts := []containerdoci.SpecOpts{oci.WithUIDGID(uid, gid, sgids)}
 	if system.SeccompSupported() {
-		opts = append(opts, seccomp.WithDefaultProfile())
+		// TODO
 	}
 	if meta.ReadonlyRootFS {
 		opts = append(opts, containerdoci.WithRootFSReadonly())
@@ -216,9 +214,7 @@ func (w *runcExecutor) Exec(ctx context.Context, meta 
 		return err
 	}
 	if w.rootless {
-		if err := rootlessspecconv.ToRootless(spec); err != nil {
-			return err
-		}
+		// TODO
 	}
 
 	if err := json.NewEncoder(f).Encode(spec); err != nil {
