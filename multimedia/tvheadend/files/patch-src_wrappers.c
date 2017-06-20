--- src/wrappers.c.orig	2017-04-21 08:32:22.000000000 +0000
+++ src/wrappers.c	2017-06-20 14:13:41.940132000 +0000
@@ -191,6 +191,8 @@
   pid_t tid;
   tid = gettid();
   ret = setpriority(PRIO_PROCESS, tid, value);
+#elif PLATFORM_FREEBSD
+  ret = setpriority(PRIO_PROCESS, 0, value);
 #else
 #warning "Implement renice for your platform!"
 #endif
