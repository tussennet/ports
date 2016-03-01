--- libmpeg3/ifo.h.orig
+++ libmpeg3/ifo.h
@@ -1,6 +1,8 @@
 #ifndef __IFO_H__
 #define __IFO_H__
 
+typedef off_t __off64_t;
+
 #ifndef DVD_VIDEO_LB_LEN
 #define DVD_VIDEO_LB_LEN 2048
 #endif
