--- libcontainer/process_freebsd.go.orig	2020-10-02 12:15:25 UTC
+++ libcontainer/process_freebsd.go
@@ -0,0 +1,72 @@
+package libcontainer
+
+import (
+	"errors"
+	"os"
+	"os/exec"
+	"syscall"
+
+	"github.com/opencontainers/runc/libcontainer/system"
+)
+
+type initProcess struct {
+	cmd       *exec.Cmd
+	container *freebsdContainer
+	fds       []string
+	process   *Process
+}
+
+func (p *initProcess) start() {
+	p.process.ops = p
+}
+
+func (p *initProcess) pid() int {
+	return p.cmd.Process.Pid
+}
+
+func (p *initProcess) externalDescriptors() []string {
+	return p.fds
+}
+
+func (p *initProcess) wait() (*os.ProcessState, error) {
+	err := p.cmd.Wait()
+	if err != nil {
+		return p.cmd.ProcessState, err
+	}
+	return p.cmd.ProcessState, nil
+}
+
+func (p *initProcess) terminate() error {
+	if p.cmd.Process == nil {
+		return nil
+	}
+	err := p.cmd.Process.Kill()
+	if _, werr := p.wait(); err == nil {
+		err = werr
+	}
+	return err
+}
+
+func (p *initProcess) startTime() (string, error) {
+	return system.GetProcessStartTime(p.pid())
+}
+
+/*
+func (p *initProcess) sendConfig() error {
+	// send the config to the container's init process, we don't use JSON Encode
+	// here because there might be a problem in JSON decoder in some cases, see:
+	// https://github.com/docker/docker/issues/14203#issuecomment-174177790
+	return utils.WriteJSON(p.parentPipe, p.config)
+}
+*/
+func (p *initProcess) signal(sig os.Signal) error {
+	s, ok := sig.(syscall.Signal)
+	if !ok {
+		return errors.New("os: unsupported signal type")
+	}
+	return syscall.Kill(p.pid(), s)
+}
+
+func (p *initProcess) setExternalDescriptors(newFds []string) {
+	p.fds = newFds
+}
