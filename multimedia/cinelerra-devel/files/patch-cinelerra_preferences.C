--- cinelerra/preferences.C.orig	2010-12-20 11:54:26.000000000 +0200
+++ cinelerra/preferences.C	2010-12-20 12:13:02.000000000 +0200
@@ -35,10 +35,30 @@
 #include "videoconfig.h"
 #include "videodevice.inc"
 #include <string.h>
+#if defined(__FreeBSD__) || defined(__FreeBSD_kernel__)
+#include <sys/types.h>
+#include <sys/sysctl.h>
+#include <sys/errno.h>
+#endif
 
 //#define CLAMP(x, y, z) (x) = ((x) < (y) ? (y) : ((x) > (z) ? (z) : (x)))
 
-
+#if defined(__FreeBSD__) || defined(__FreeBSD_kernel__)
+#define GETSYSCTL(name, var)    getsysctl(name, &(var), sizeof(var))
+static int getsysctl(const char *name, void *ptr, size_t len)
+{
+ size_t nlen = len;
+ if (sysctlbyname(name, ptr, &nlen, NULL, 0) == -1)
+ {
+  return -1;
+ }
+ if (nlen != len && errno == ENOMEM)
+ {
+  return -1;
+ }
+ return 0;
+}
+#endif
 
 
 
@@ -610,6 +630,18 @@
 {
 /* Get processor count */
 	int result = 1;
+#if defined(__FreeBSD__) || defined(__FreeBSD_kernel__)
+        size_t cpu_count_len = sizeof(result);
+        
+        if (GETSYSCTL("hw.ncpu", result) == 0)
+        {
+        }
+        else
+        {
+         fprintf(stderr, "Cannot get hw.ncpu\n");
+         result = 1;
+        }
+#else
 	FILE *proc;
 
 	if(force_uniprocessor && !interactive) return 1;
@@ -642,7 +674,7 @@
 		}
 		fclose(proc);
 	}
-
+#endif
 	return result;
 }
 
