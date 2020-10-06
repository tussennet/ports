--- libcontainer/utils/utils_unix.go.orig	2020-10-02 12:15:24 UTC
+++ libcontainer/utils/utils_unix.go
@@ -16,9 +16,6 @@ func EnsureProcHandle(fh *os.File) error {
 	if err := unix.Fstatfs(int(fh.Fd()), &buf); err != nil {
 		return fmt.Errorf("ensure %s is on procfs: %v", fh.Name(), err)
 	}
-	if buf.Type != unix.PROC_SUPER_MAGIC {
-		return fmt.Errorf("%s is not on procfs", fh.Name())
-	}
 	return nil
 }
 
