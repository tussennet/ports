--- quicktime/util.c.orig
+++ quicktime/util.c
@@ -1,5 +1,5 @@
 #include <fcntl.h>
-#include <linux/cdrom.h>
+#include <sys/cdio.h>
 #include <stdio.h>
 #include <stdlib.h>
 #include <string.h>
@@ -15,9 +15,9 @@
 
 int64_t quicktime_get_file_length(char *path)
 {
-	struct stat64 status;
-	if(stat64(path, &status))
-		perror("quicktime_get_file_length stat64:");
+	struct stat status;
+	if(stat(path, &status))
+		perror("quicktime_get_file_length stat:");
 	return status.st_size;
 }
 
