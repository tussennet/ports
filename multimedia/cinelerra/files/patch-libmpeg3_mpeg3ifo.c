--- libmpeg3/mpeg3ifo.c.orig	2015-08-13 14:04:04 UTC
+++ libmpeg3/mpeg3ifo.c
@@ -1,4 +1,4 @@
-#include <byteswap.h>
+//#include <byteswap.h>
 #include <dirent.h>
 #include <fcntl.h>
 #include <stdlib.h>
@@ -10,6 +10,10 @@
 #include "mpeg3private.h"
 #include "mpeg3protos.h"
 
+#include <sys/endian.h>
+#define bswap_16(x) bswap16(x)
+#define bswap_32(x) bswap32(x)
+
 typedef struct
 {
 // Bytes relative to start of stream.
