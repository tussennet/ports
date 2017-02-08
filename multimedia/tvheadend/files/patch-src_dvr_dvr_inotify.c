--- src/dvr/dvr_inotify.c.orig	2017-01-20 19:41:51.000000000 +0100
+++ src/dvr/dvr_inotify.c	2017-02-08 11:21:07.222439000 +0100
@@ -119,7 +119,7 @@
   if (filename == NULL || fd < 0)
     return;
 
-  path = strdupa(filename);
+  path = tvh_strdupa(filename);
 
   SKEL_ALLOC(dvr_inotify_entry_skel);
   dvr_inotify_entry_skel->path = dirname(path);
