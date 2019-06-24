--- vendor/github.com/moby/buildkit/executor/oci/spec_unix.go.orig	2019-06-18 21:30:11 UTC
+++ vendor/github.com/moby/buildkit/executor/oci/spec_unix.go
@@ -8,7 +8,6 @@ import (
 	"sync"
 
 	"github.com/containerd/containerd/containers"
-	"github.com/containerd/containerd/contrib/seccomp"
 	"github.com/containerd/containerd/mount"
 	"github.com/containerd/containerd/namespaces"
 	"github.com/containerd/containerd/oci"
@@ -52,7 +51,8 @@ func GenerateSpec(ctx context.Context, meta executor.M
 	if meta.SecurityMode == pb.SecurityMode_INSECURE {
 		opts = append(opts, entitlements.WithInsecureSpec())
 	} else if system.SeccompSupported() && meta.SecurityMode == pb.SecurityMode_SANDBOX {
-		opts = append(opts, seccomp.WithDefaultProfile())
+		// TODO
+		return nil, nil, errors.New("TODO Seccomp Sandbox not supported on FreeBSD")
 	}
 
 	switch processMode {
