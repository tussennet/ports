--- dvb/osd.h.orig	2018-07-12 16:04:03 UTC
+++ dvb/osd.h
@@ -25,7 +25,7 @@
 #ifndef _DVBOSD_H_
 #define _DVBOSD_H_
 
-#include <linux/compiler.h>
+#include <sys/types.h>
 
 typedef enum {
   // All functions return -2 on "not open"
