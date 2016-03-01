--- cinelerra/mwindow.C.orig	2010-12-19 11:09:09.000000000 +0200
+++ cinelerra/mwindow.C	2010-12-27 12:23:59.000000000 +0200
@@ -105,7 +105,8 @@
 #include "exportedl.h"
 
 #include <string.h>
-
+#include <sys/types.h>
+#include <sys/sysctl.h>
 
 
 extern "C"
@@ -1251,6 +1252,23 @@
 
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
@@ -1278,6 +1296,7 @@
 			"echo \"0x7fffffff\" > /proc/sys/kernel/shmmax\n",
 			result);
 	}
+#endif
 }
 
 
