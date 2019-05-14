--- src/init/init.sh.orig	2019-05-14 21:25:40 UTC
+++ src/init/init.sh
@@ -164,7 +164,7 @@ runInit()
         return 0;
     fi
 
-    if [ "X${UN}" = "XOpenBSD" -o "X${UN}" = "XNetBSD" -o "X${UN}" = "XFreeBSD" -o "X${UN}" = "XDragonFly" ]; then
+    if [ "X${UN}" = "XOpenBSD" -o "X${UN}" = "XNetBSD" -o "X${UN}" = "XDragonFly" ]; then
         # Checking for the presence of ossec-control on rc.local
         grep ossec-control /etc/rc.local > /dev/null 2>&1
         if [ $? != 0 ]; then
