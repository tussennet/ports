--- src/httpc.c.orig	2017-01-20 19:41:51.000000000 +0100
+++ src/httpc.c	2017-02-08 11:22:13.686848000 +0100
@@ -1253,7 +1253,7 @@
 
   if (args == NULL)
     return;
-  p = strdupa(args);
+  p = tvh_strdupa(args);
   while (*p) {
     while (*p && *p <= ' ') p++;
     if (*p == '\0') break;
