--- cinelerra/renderfarmclient.C.orig	2015-08-13 14:04:04 UTC
+++ cinelerra/renderfarmclient.C
@@ -130,7 +130,7 @@ void RenderFarmClient::main_loop()
 	else
 	{
 		struct sockaddr_un addr;
-		addr.sun_family = AF_FILE;
+		addr.sun_family = AF_LOCAL;
 		strcpy(addr.sun_path, deamon_path);
 		int size = (offsetof(struct sockaddr_un, sun_path) + 
 			strlen(deamon_path) + 1);
