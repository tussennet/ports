--- daemon/daemon_linux.go.orig	2020-09-04 14:54:51 UTC
+++ daemon/daemon_linux.go
@@ -145,3 +145,43 @@ func setupResolvConf(config *config.Config) {
 	}
 	config.ResolvConf = resolvconf.Path()
 }
+
+// setupDaemonProcess sets various settings for the daemon's process
+func setupDaemonProcess(config *config.Config) error {
+	// setup the daemons oom_score_adj
+	return setupOOMScoreAdj(config.OOMScoreAdjust)
+}
+
+func setupOOMScoreAdj(score int) error {
+	f, err := os.OpenFile("/proc/self/oom_score_adj", os.O_WRONLY, 0)
+	if err != nil {
+		return err
+	}
+	defer f.Close()
+	stringScore := strconv.Itoa(score)
+	_, err = f.WriteString(stringScore)
+	if os.IsPermission(err) {
+		// Setting oom_score_adj does not work in an
+		// unprivileged container. Ignore the error, but log
+		// it if we appear not to be in that situation.
+		if !rsystem.RunningInUserNS() {
+			logrus.Debugf("Permission denied writing %q to /proc/self/oom_score_adj", stringScore)
+		}
+		return nil
+	}
+
+	return err
+}
+
+func (daemon *Daemon) setupSeccompProfile() error {
+	if daemon.configStore.SeccompProfile != "" {
+		daemon.seccompProfilePath = daemon.configStore.SeccompProfile
+		b, err := ioutil.ReadFile(daemon.configStore.SeccompProfile)
+		if err != nil {
+			return fmt.Errorf("opening seccomp profile (%s) failed: %v", daemon.configStore.SeccompProfile, err)
+		}
+		daemon.seccompProfile = b
+	}
+	return nil
+}
+
