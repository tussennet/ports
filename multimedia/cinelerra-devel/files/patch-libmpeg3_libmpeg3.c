--- libmpeg3/libmpeg3.c.orig
+++ libmpeg3/libmpeg3.c
@@ -6,6 +6,7 @@
 #include <fcntl.h>
 #include <stdlib.h>
 #include <string.h>
+#include <sys/types.h>
 
 
 int mpeg3_major()
