- Fix CVE-2015-3395

Obtained from:	https://git.libav.org/?p=libav.git;a=commit;h=5ecabd3c54b7c802522dc338838c9a4c2dc42948
--- mythtv/external/FFmpeg/libavcodec/msrledec.c.orig	2016-02-01 22:41:41.000000000 +0100
+++ mythtv/external/FFmpeg/libavcodec/msrledec.c	2016-04-13 13:09:45.999643000 +0200
@@ -39,7 +39,7 @@
     unsigned int pixel_ptr = 0;
     int row_dec = pic->linesize[0];
     int row_ptr = (avctx->height - 1) * row_dec;
-    int frame_size = row_dec * avctx->height;
+    int frame_size = FFABS(row_dec) * avctx->height;
     int i;
 
     while (row_ptr >= 0) {
