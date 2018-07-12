--- lib/libdvbv5/dvb-dev-remote.c.orig	2017-12-27 13:50:55 UTC
+++ lib/libdvbv5/dvb-dev-remote.c
@@ -32,9 +32,13 @@
 #include <libudev.h>
 #include <stdio.h>
 #include <stdlib.h>
+#include <stdarg.h>
 #include <locale.h>
 #include <pthread.h>
 #include <unistd.h>
+#include <sys/types.h>
+#include <netinet/in.h>
+#include <arpa/nameser.h>
 #include <resolv.h>
 #include <string.h>
 #include <sys/socket.h>
@@ -50,6 +54,10 @@
 # define _(string) string
 #endif
 
+#ifndef MSG_MORE
+#define MSG_MORE 0
+#endif
+
 
 /*
  * Expected server version
