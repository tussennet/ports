--- src/download.c.orig	2017-01-20 19:41:51.000000000 +0100
+++ src/download.c	2017-02-08 11:18:55.155793000 +0100
@@ -172,7 +172,7 @@
     sbuf_alloc(&dn->pipe_sbuf, 2048);
     len = sbuf_read(&dn->pipe_sbuf, dn->pipe_fd);
     if (len == 0) {
-      s = dn->url ? strdupa(dn->url) : strdupa("");
+      s = dn->url ? tvh_strdupa(dn->url) : tvh_strdupa("");
       p = strchr(s, ' ');
       if (p)
         *p = '\0';
@@ -248,7 +248,7 @@
     goto done;
 
   if (strncmp(dn->url, "file://", 7) == 0) {
-    char *f = strdupa(dn->url + 7);
+    char *f = tvh_strdupa(dn->url + 7);
     http_deescape(f);
     download_file(dn, f);
     goto done;
