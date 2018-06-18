--- test/test_dvr.c.orig	2014-03-21 19:26:36 UTC
+++ test/test_dvr.c
@@ -129,7 +129,7 @@ int main(int argc, char *argv[])
 
 	fprintf(stderr, "using '%s' and '%s'\n"
 		"writing to '%s'\n", dmxdev, dvrdev, argv[1]);
-	tsfd = open(argv[1], O_WRONLY | O_CREAT | O_TRUNC | O_LARGEFILE, 0664);
+	tsfd = open(argv[1], O_WRONLY | O_CREAT | O_TRUNC, 0664);
 	if (tsfd == -1) {
 		perror("cannot write output file");
 		return 1;
