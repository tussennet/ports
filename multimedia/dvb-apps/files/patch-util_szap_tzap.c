--- util/szap/tzap.c.orig	2014-03-21 19:26:36 UTC
+++ util/szap/tzap.c
@@ -650,7 +650,7 @@ int main(int argc, char **argv)
 	if (record) {
 		if (filename!=NULL) {
 			if (strcmp(filename,"-")!=0) {
-				file_fd = open(filename,O_WRONLY|O_LARGEFILE|O_CREAT,0644);
+				file_fd = open(filename,O_WRONLY|O_CREAT,0644);
 				if (file_fd<0) {
 					PERROR("open of '%s' failed",filename);
 					return -1;
