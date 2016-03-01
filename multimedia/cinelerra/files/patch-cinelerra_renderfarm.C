--- cinelerra/renderfarm.C.orig
+++ cinelerra/renderfarm.C
@@ -173,7 +173,7 @@ int RenderFarmServerThread::open_client(char *hostname, int port)
 		else
 		{
 			struct sockaddr_un addr;
-			addr.sun_family = AF_FILE;
+			addr.sun_family = AF_LOCAL;
 			strcpy(addr.sun_path, hostname);
 			int size = (offsetof(struct sockaddr_un, sun_path) + 
 				strlen(hostname) + 1);
