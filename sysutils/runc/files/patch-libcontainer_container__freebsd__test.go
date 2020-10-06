--- libcontainer/container_freebsd_test.go.orig	2020-10-02 12:15:25 UTC
+++ libcontainer/container_freebsd_test.go
@@ -0,0 +1,36 @@
+// +build freebsd
+
+package libcontainer
+
+import (
+	"testing"
+
+	"github.com/opencontainers/runc/libcontainer/configs"
+)
+
+func TestGetContainerState(t *testing.T) {
+	var (
+		pid       int
+		startTime uint64
+	)
+	pid = 123
+	startTime = 456789
+	container := &freebsdContainer{
+		jailId:               "mockJailId",
+		initProcessPid:       pid,
+		initProcessStartTime: startTime,
+		id:                   "myid",
+		config:               &configs.Config{},
+	}
+	container.state = &createdState{c: container}
+	state, err := container.State()
+	if err != nil {
+		t.Fatal(err)
+	}
+	if state.InitProcessPid != pid {
+		t.Fatalf("expected pid %d but received %d", pid, state.InitProcessPid)
+	}
+	if state.InitProcessStartTime != startTime {
+		t.Fatalf("expected process start time 10 but received %d", state.InitProcessStartTime)
+	}
+}
