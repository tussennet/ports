--- quicktime/qtprivate.h.orig
+++ quicktime/qtprivate.h
@@ -30,12 +30,13 @@
 #include <stdio.h>
 #include <stdint.h>
 #include <stdlib.h>
+#include <sys/types.h>
 
 
 
 
-#define FTELL ftello64
-#define FSEEK fseeko64
+#define FTELL ftello
+#define FSEEK fseeko
 
 
 // ffmpeg requires global variable initialization
