CC              src/imagecache.o
src/imagecache.c:558:10: error: implicit declaration of function 'strdupa' is invalid in C99
      [-Werror,-Wimplicit-function-declaration]
    fn = strdupa(i->url + 7);
         ^
src/imagecache.c:558:10: note: did you mean 'strdup'?
/usr/include/string.h:85:7: note: 'strdup' declared here
char    *strdup(const char *) __malloc_like;
         ^
1 error generated.

--- src/imagecache.c.orig	2016-03-14 10:10:57.000000000 +0100
+++ src/imagecache.c	2017-02-08 09:52:44.457019000 +0100
@@ -555,7 +555,7 @@
 
   /* Local file */
   if (!strncasecmp(i->url, "file://", 7)) {
-    fn = strdupa(i->url + 7);
+    fn = tvh_strdupa(i->url + 7);
     http_deescape(fn);
     fd = open(fn, O_RDONLY);
   }
