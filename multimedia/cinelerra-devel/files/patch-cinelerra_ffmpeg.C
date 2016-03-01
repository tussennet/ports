--- cinelerra/ffmpeg.C.orig
+++ cinelerra/ffmpeg.C
@@ -2,7 +2,7 @@
 
 #ifdef HAVE_SWSCALER
 extern "C" {
-#include <swscale.h>
+#include <ffmpeg/swscale.h>
 }
 #endif
 
