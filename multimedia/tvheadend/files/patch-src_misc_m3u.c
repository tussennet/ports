--- src/misc/m3u.c.orig	2017-01-20 19:41:51.000000000 +0100
+++ src/misc/m3u.c	2017-02-08 11:19:51.401412000 +0100
@@ -103,7 +103,7 @@
   if (rel[0] == '/') {
     snprintf(buf, buflen, "%s%s", url, rel + 1);
   } else {
-    url2 = strdupa(url);
+    url2 = tvh_strdupa(url);
     p = strrchr(url2, '/');
     if (p == NULL)
       return rel;
