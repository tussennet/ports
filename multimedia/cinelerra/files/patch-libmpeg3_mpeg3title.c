--- libmpeg3/mpeg3title.c.orig	2015-08-13 14:04:04 UTC
+++ libmpeg3/mpeg3title.c
@@ -5,6 +5,7 @@
 
 #include <stdlib.h>
 #include <string.h>
+#include <unistd.h>
 
 mpeg3_title_t* mpeg3_new_title(mpeg3_t *file, char *path)
 {
