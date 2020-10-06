--- libcontainer/specconv/spec_freebsd.go.orig	2020-10-02 12:15:25 UTC
+++ libcontainer/specconv/spec_freebsd.go
@@ -0,0 +1,53 @@
+// Package specconv implements conversion of specifications to libcontainer
+// configurations
+package specconv
+
+import (
+	"fmt"
+	"os"
+	"path/filepath"
+
+	"github.com/opencontainers/runc/libcontainer/configs"
+	"github.com/opencontainers/runtime-spec/specs-go"
+)
+
+type CreateOpts struct {
+	NoPivotRoot  bool
+	NoNewKeyring bool
+	Spec         *specs.Spec
+	Rootless     bool
+}
+
+// given specification and a cgroup name
+func CreateLibcontainerConfig(opts *CreateOpts) (*configs.Config, error) {
+	// runc's cwd will always be the bundle path
+	rcwd, err := os.Getwd()
+	if err != nil {
+		return nil, err
+	}
+	cwd, err := filepath.Abs(rcwd)
+	if err != nil {
+		return nil, err
+	}
+	spec := opts.Spec
+	rootfsPath := spec.Root.Path
+	if !filepath.IsAbs(rootfsPath) {
+		rootfsPath = filepath.Join(cwd, rootfsPath)
+	}
+	labels := []string{}
+	for k, v := range spec.Annotations {
+		labels = append(labels, fmt.Sprintf("%s=%s", k, v))
+	}
+	config := &configs.Config{
+		Rootfs:       rootfsPath,
+		NoPivotRoot:  opts.NoPivotRoot,
+		Readonlyfs:   spec.Root.Readonly,
+		Hostname:     spec.Hostname,
+		Labels:       append(labels, fmt.Sprintf("bundle=%s", cwd)),
+		NoNewKeyring: opts.NoNewKeyring,
+		RootlessEUID:     opts.Rootless,
+		RootlessCgroups:  opts.Rootless,
+		Version:      specs.Version,
+	}
+	return config, err
+}
