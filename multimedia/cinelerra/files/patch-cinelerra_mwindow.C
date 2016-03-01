--- cinelerra/mwindow.C.orig	2015-08-13 14:04:04 UTC
+++ cinelerra/mwindow.C
@@ -111,7 +111,8 @@
 #include <sys/stat.h>
 #include <fcntl.h>
 #include <string.h>
-
+#include <sys/types.h>
+#include <sys/sysctl.h>
 
 
 extern "C"
@@ -1338,6 +1339,23 @@ void MWindow::test_plugins(EDL *new_edl,
 
 void MWindow::init_shm()
 {
+#if defined(__FreeBSD__) || defined(__FreeBSD_kernel__)
+ uint64_t result=0;
+ size_t len = sizeof(result);
+ 
+ if (sysctlbyname("kern.ipc.shmmax", &result, &len, NULL, 0) == -1) {
+                MainError::show_error("MWindow::init_shm: couldn't get kern.ipc.shmmax\n");
+                return;
+ }
+ 
+ if(result < 0x7fffffff)
+ {       
+         eprintf("WARNING: kern.ipc.shmmax is 0x%llx, which is too low.\n"
+                 "Before running Cinelerra do the following as root:\n"
+                 "sysctl -w kern.ipc.shmmax=\"2147483647\"\n",
+                 result);
+ }
+#else
 // Fix shared memory
 	FILE *fd = fopen("/proc/sys/kernel/shmmax", "w");
 	if(fd)
@@ -1365,6 +1383,7 @@ void MWindow::init_shm()
 			"echo \"0x7fffffff\" > /proc/sys/kernel/shmmax\n",
 			result);
 	}
+#endif
 }
 
 void MWindow::create_objects(int want_gui, 
