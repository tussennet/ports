--- src/http.c.orig	2017-01-20 19:41:51.000000000 +0100
+++ src/http.c	2017-02-08 10:23:07.867231000 +0100
@@ -31,6 +31,7 @@
 #include <netinet/tcp.h>
 #include <arpa/inet.h>
 #include <openssl/md5.h>
+#include <sys/socket.h>
 
 #include "tvheadend.h"
 #include "tcp.h"
