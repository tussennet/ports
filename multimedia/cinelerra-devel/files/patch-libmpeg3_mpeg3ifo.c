--- libmpeg3/mpeg3ifo.c.orig	2010-12-19 13:56:22.000000000 +0300
+++ libmpeg3/mpeg3ifo.c	2010-12-26 16:00:24.291223315 +0300
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
