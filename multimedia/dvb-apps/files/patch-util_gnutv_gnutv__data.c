--- util/gnutv/gnutv_data.c.orig	2014-03-21 19:26:36 UTC
+++ util/gnutv/gnutv_data.c
@@ -33,8 +33,10 @@
 #include <pthread.h>
 #include <errno.h>
 #include <sys/poll.h>
+#include <sys/types.h>
 #include <sys/socket.h>
 #include <netinet/in.h>
+#include <netdb.h>
 #include <arpa/inet.h>
 #include <libdvbapi/dvbdemux.h>
 #include <libdvbapi/dvbaudio.h>
@@ -98,7 +100,7 @@ void gnutv_data_start(int _output_type,
 	case OUTPUT_TYPE_FILE:
 		if (output_type == OUTPUT_TYPE_FILE) {
 			// open output file
-			outfd = open(outfile, O_WRONLY|O_CREAT|O_LARGEFILE|O_TRUNC, 0644);
+			outfd = open(outfile, O_WRONLY|O_CREAT|O_TRUNC, 0644);
 			if (outfd < 0) {
 				fprintf(stderr, "Failed to open output file\n");
 				exit(1);
@@ -137,7 +139,8 @@ void gnutv_data_start(int _output_type,
 
 		// bind to local interface if requested
 		if (outif != NULL) {
-			if (setsockopt(outfd, SOL_SOCKET, SO_BINDTODEVICE, outif, strlen(outif)) < 0) {
+			if (1 /*setsockopt(outfd, SOL_SOCKET, SO_BINDTODEVICE, outif, strlen(outif)) < 0*/) {
+				fprintf(stderr, "SO_BINDTODEVICE not supported on FreeBSD\n");
 				fprintf(stderr, "Failed to bind to interface %s\n", outif);
 				exit(1);
 			}
