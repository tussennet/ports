--- src/wrappers.c.orig	2017-01-20 19:41:51.000000000 +0100
+++ src/wrappers.c	2017-02-08 11:35:24.814465000 +0100
@@ -191,6 +191,8 @@
   pid_t tid;
   tid = gettid();
   ret = setpriority(PRIO_PROCESS, tid, value);
+#elif PLATFORM_FREEBSD
+  ret = setpriority(PRIO_PROCESS, 0, value);
 #else
 #warning "Implement renice for your platform!"
 #endif
