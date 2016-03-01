--- libmpeg3/mpeg3tocutil.c.orig	2010-12-26 16:14:27.556227108 +0300
+++ libmpeg3/mpeg3tocutil.c	2010-12-26 16:14:34.918227865 +0300
@@ -1415,8 +1415,8 @@
 
 int64_t mpeg3_calculate_source_date(char *path)
 {
-	struct stat64 ostat;
-	bzero(&ostat, sizeof(struct stat64));
-	stat64(path, &ostat);
+	struct stat ostat;
+	bzero(&ostat, sizeof(struct stat));
+	stat(path, &ostat);
 	return ostat.st_mtime;
 }
