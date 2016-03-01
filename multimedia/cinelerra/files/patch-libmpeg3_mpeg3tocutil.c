--- libmpeg3/mpeg3tocutil.c.orig	2015-08-13 14:04:04 UTC
+++ libmpeg3/mpeg3tocutil.c
@@ -1417,8 +1417,8 @@ int64_t mpeg3_get_source_date(mpeg3_t *f
 
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
