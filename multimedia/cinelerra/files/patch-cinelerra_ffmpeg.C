--- cinelerra/ffmpeg.C.orig	2015-08-13 14:04:04 UTC
+++ cinelerra/ffmpeg.C
@@ -2,7 +2,7 @@
 
 #ifdef HAVE_SWSCALER
 extern "C" {
-#include <swscale.h>
+#include <ffmpeg/swscale.h>
 }
 #endif
 
