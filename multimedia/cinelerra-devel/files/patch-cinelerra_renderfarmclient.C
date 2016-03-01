--- cinelerra/renderfarmclient.C.orig
+++ cinelerra/renderfarmclient.C
@@ -124,7 +124,7 @@ void RenderFarmClient::main_loop()
 	else
 	{
 		struct sockaddr_un addr;
-		addr.sun_family = AF_FILE;
+		addr.sun_family = AF_LOCAL;
 		strcpy(addr.sun_path, deamon_path);
 		int size = (offsetof(struct sockaddr_un, sun_path) + 
 			strlen(deamon_path) + 1);
