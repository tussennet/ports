--- daemon/graphdriver/driver_freebsd.go.orig	2019-10-07 21:12:15 UTC
+++ daemon/graphdriver/driver_freebsd.go
@@ -1,8 +1,7 @@
 package graphdriver // import "github.com/docker/docker/daemon/graphdriver"
 
 import (
-	"syscall"
-
+	"github.com/docker/docker/pkg/mount"
 	"golang.org/x/sys/unix"
 )
 
@@ -11,10 +10,49 @@ var (
 	priority = "zfs"
 )
 
+// GetFSMagic returns the filesystem id given the path.
+func GetFSMagic(rootpath string) (FsMagic, error) {
+	var buf unix.Statfs_t
+	if err := unix.Statfs(rootpath, &buf); err != nil {
+		return 0, err
+	}
+	return FsMagic(buf.Type), nil
+}
+
+// NewFsChecker returns a checker configured for the provided FsMagic
+func NewFsChecker(t FsMagic) Checker {
+	return &fsChecker{
+		t: t,
+	}
+}
+
+type fsChecker struct {
+	t FsMagic
+}
+
+func (c *fsChecker) IsMounted(path string) bool {
+	m, _ := Mounted(c.t, path)
+	return m
+}
+
+// NewDefaultChecker returns a check that parses /proc/mountinfo to check
+// if the specified path is mounted.
+func NewDefaultChecker() Checker {
+	return &defaultChecker{}
+}
+
+type defaultChecker struct {
+}
+
+func (c *defaultChecker) IsMounted(path string) bool {
+	m, _ := mount.Mounted(path)
+	return m
+}
+
 // Mounted checks if the given path is mounted as the fs type
 func Mounted(fsType FsMagic, mountPath string) (bool, error) {
 	var buf unix.Statfs_t
-	if err := syscall.Statfs(mountPath, &buf); err != nil {
+	if err := unix.Statfs(mountPath, &buf); err != nil {
 		return false, err
 	}
 	return FsMagic(buf.Type) == fsType, nil
