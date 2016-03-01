--- cinelerra/preferences.C.orig	2015-08-13 14:04:04 UTC
+++ cinelerra/preferences.C
@@ -35,9 +35,31 @@
 #include "videoconfig.h"
 #include "videodevice.inc"
 #include <string.h>
+#if defined(__FreeBSD__) || defined(__FreeBSD_kernel__)
+#include <sys/types.h>
+#include <sys/sysctl.h>
+#include <sys/errno.h>
+#endif
 
 //#define CLAMP(x, y, z) (x) = ((x) < (y) ? (y) : ((x) > (z) ? (z) : (x)))
 
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
+
 Preferences::Preferences()
 {
 // Set defaults
@@ -599,6 +621,18 @@ int Preferences::calculate_processors(in
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
@@ -631,7 +665,7 @@ int Preferences::calculate_processors(in
 		}
 		fclose(proc);
 	}
-
+#endif
 	return result;
 }
 
