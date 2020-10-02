--- main_freebsd.go.orig	2020-10-02 12:15:25 UTC
+++ main_freebsd.go
@@ -0,0 +1,13 @@
+package main
+
+import "github.com/urfave/cli"
+
+var (
+	checkpointCommand cli.Command
+	eventsCommand     cli.Command
+	restoreCommand    cli.Command
+	//initCommand       cli.Command
+	pauseCommand      cli.Command
+	resumeCommand     cli.Command
+	updateCommand     cli.Command
+)
