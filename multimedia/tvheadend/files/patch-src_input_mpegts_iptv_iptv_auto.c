--- src/input/mpegts/iptv/iptv_auto.c.orig	2017-01-20 19:41:51.000000000 +0100
+++ src/input/mpegts/iptv/iptv_auto.c	2017-02-08 11:23:44.790547000 +0100
@@ -123,7 +123,7 @@
   tags = htsmsg_get_str(item, "tvh-tags");
   if (!tags) tags = htsmsg_get_str(item, "group-title");
   if (tags) {
-    tags = n = strdupa(tags);
+    tags = n = tvh_strdupa(tags);
     while (*n) {
       if (*n == '|')
         *n = '\n';
