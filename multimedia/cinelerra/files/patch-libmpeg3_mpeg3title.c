--- libmpeg3/mpeg3title.c.orig	2010-12-26 15:48:33.181227771 +0300
+++ libmpeg3/mpeg3title.c	2010-12-26 15:52:53.339237239 +0300
@@ -5,6 +5,7 @@
 
 #include <stdlib.h>
 #include <string.h>
+#include <unistd.h>
 
 mpeg3_title_t* mpeg3_new_title(mpeg3_t *file, char *path)
 {
