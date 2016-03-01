--- cinelerra/renderfarm.C.orig	2015-08-13 14:04:04 UTC
+++ cinelerra/renderfarm.C
@@ -163,7 +163,7 @@ int RenderFarmServerThread::open_client(
 		else
 		{
 			struct sockaddr_un addr;
-			addr.sun_family = AF_FILE;
+			addr.sun_family = AF_LOCAL;
 			strcpy(addr.sun_path, hostname);
 			int size = (offsetof(struct sockaddr_un, sun_path) + 
 				strlen(hostname) + 1);
