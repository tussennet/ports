--- ctr/checkpoint_unix.go.orig	2020-10-02 13:03:31 UTC
+++ ctr/checkpoint_unix.go
@@ -0,0 +1,38 @@
+// +build !linux
+
+package main
+
+import (
+	"fmt"
+
+	"github.com/urfave/cli"
+)
+
+var checkpointSubCmds = []cli.Command{
+	listCheckpointCommand,
+}
+
+var checkpointCommand = cli.Command{
+	Name:        "checkpoints",
+	Usage:       "list all checkpoints",
+	ArgsUsage:   "COMMAND [arguments...]",
+	Subcommands: checkpointSubCmds,
+	Description: func() string {
+		desc := "\n    COMMAND:\n"
+		for _, command := range checkpointSubCmds {
+			desc += fmt.Sprintf("    %-10.10s%s\n", command.Name, command.Usage)
+		}
+		return desc
+	}(),
+	Action: listCheckpoints,
+}
+
+var listCheckpointCommand = cli.Command{
+	Name:   "list",
+	Usage:  "list all checkpoints for a container",
+	Action: listCheckpoints,
+}
+
+func listCheckpoints(context *cli.Context) {
+	fatal("checkpoint command is not supported on Solaris", ExitStatusUnsupported)
+}
