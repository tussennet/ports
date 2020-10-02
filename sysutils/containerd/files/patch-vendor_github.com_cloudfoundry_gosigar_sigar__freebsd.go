--- vendor/github.com/cloudfoundry/gosigar/sigar_freebsd.go.orig	2020-10-02 13:03:31 UTC
+++ vendor/github.com/cloudfoundry/gosigar/sigar_freebsd.go
@@ -0,0 +1,52 @@
+// -------------------------------------------------------
+// FreeBSD stub for compiling sigar
+// @(kris-nova)
+//
+// Adding out stubs here to get sigar compiling without an
+// implementation, this patch should handle getting docker
+// working, and the Go compiler happy and NOTHING more.
+
+// +build freebsd
+
+package sigar
+
+import "syscall"
+
+func (self *FileSystemUsage) Get(path string) error {
+	stat := syscall.Statfs_t{}
+	err := syscall.Statfs(path, &stat)
+	if err != nil {
+		return err
+	}
+
+	bsize := stat.Bsize / 512
+
+	self.Total = (uint64(stat.Blocks) * uint64(bsize)) >> 1
+	self.Free = (uint64(stat.Bfree) * uint64(bsize)) >> 1
+	self.Avail = (uint64(stat.Bavail) * uint64(bsize)) >> 1
+	self.Used = self.Total - self.Free
+	self.Files = stat.Files
+	self.FreeFiles = uint64(stat.Ffree)
+
+	return nil
+}
+
+func (self *Cpu) Get() error {
+	return nil
+}
+
+func (self *Mem) Get() error {
+	return nil
+}
+
+func (self *Swap) Get() error {
+	return nil
+}
+
+func (self *LoadAverage) Get() error {
+	return nil
+}
+
+func (self *CpuList) Get() error {
+	return nil
+}
