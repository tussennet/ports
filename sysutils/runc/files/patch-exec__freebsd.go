--- exec_freebsd.go.orig	2020-10-02 12:15:25 UTC
+++ exec_freebsd.go
@@ -0,0 +1,125 @@
+package main
+
+import (
+	"encoding/json"
+	"fmt"
+	"os"
+
+	"github.com/opencontainers/runc/libcontainer"
+	"github.com/opencontainers/runc/libcontainer/utils"
+	"github.com/opencontainers/runtime-spec/specs-go"
+	"github.com/urfave/cli"
+)
+
+var execCommand = cli.Command{
+	Name:  "exec",
+	Usage: "execute new process inside the container",
+	ArgsUsage: `<container-id> <command> [command options]  || -p process.json <container-id>
+
+Where "<container-id>" is the name for the instance of the container and
+"<command>" is the command to be executed in the container.
+"<command>" can't be empty unless a "-p" flag provided.
+
+EXAMPLE:
+For example, if the container is configured to run the linux ps command the
+following will output a list of processes running in the container:
+
+       # runc exec <container-id> ps`,
+	Flags: []cli.Flag{
+		cli.StringFlag{
+			Name:  "cwd",
+			Usage: "current working directory in the container",
+		},
+		cli.StringFlag{
+			Name:  "pid-file",
+			Value: "",
+			Usage: "specify the file to write the process id to",
+		},
+		cli.StringFlag{
+			Name:  "process, p",
+			Usage: "path to the process.json",
+		},
+	},
+	Action: func(context *cli.Context) error {
+		if err := checkArgs(context, 1, minArgs); err != nil {
+			return err
+		}
+		if err := revisePidFile(context); err != nil {
+			return err
+		}
+		status, err := execProcess(context)
+		if err == nil {
+			os.Exit(status)
+		}
+		return fmt.Errorf("exec failed: %v", err)
+	},
+	SkipArgReorder: true,
+}
+
+func execProcess(context *cli.Context) (int, error) {
+	container, err := getContainer(context)
+	if err != nil {
+		return -1, err
+	}
+	status, err := container.Status()
+	if err != nil {
+		return -1, err
+	}
+	if status == libcontainer.Stopped {
+		return -1, fmt.Errorf("cannot exec a container that has stopped")
+	}
+	path := context.String("process")
+	if path == "" && len(context.Args()) == 1 {
+		return -1, fmt.Errorf("process args cannot be empty")
+	}
+	state, err := container.State()
+	if err != nil {
+		return -1, err
+	}
+
+	bundle := utils.SearchLabels(state.Config.Labels, "bundle")
+	p, err := getProcess(context, bundle)
+	if err != nil {
+		return -1, err
+	}
+	r := &runner{
+		enableSubreaper: false,
+		shouldDestroy:   false,
+		container:       container,
+		consoleSocket:   context.String("console-socket"),
+		pidFile:         context.String("pid-file"),
+		action:          CT_ACT_RUN,
+	}
+	return r.run(p)
+}
+
+func getProcess(context *cli.Context, bundle string) (*specs.Process, error) {
+	if path := context.String("process"); path != "" {
+		f, err := os.Open(path)
+		if err != nil {
+			return nil, err
+		}
+		defer f.Close()
+		var p specs.Process
+		if err := json.NewDecoder(f).Decode(&p); err != nil {
+			return nil, err
+		}
+		return &p, validateProcessSpec(&p)
+	}
+	// process via cli flags
+	if err := os.Chdir(bundle); err != nil {
+		return nil, err
+	}
+	spec, err := loadSpec(specConfig)
+	if err != nil {
+		return nil, err
+	}
+	p := spec.Process
+	p.Args = context.Args()[1:]
+	// override the cwd, if passed
+	if context.String("cwd") != "" {
+		p.Cwd = context.String("cwd")
+	}
+
+	return p, nil
+}
