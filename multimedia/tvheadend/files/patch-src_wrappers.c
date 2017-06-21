--- src/wrappers.c.orig	2017-05-16 11:15:24.000000000 +0000
+++ src/wrappers.c	2017-06-21 11:23:54.840229000 +0000
@@ -191,6 +191,8 @@
   pid_t tid;
   tid = gettid();
   ret = setpriority(PRIO_PROCESS, tid, value);
+#elif PLATFORM_FREEBSD
+  ret = setpriority(PRIO_PROCESS, 0, value);
 #else
 #warning "Implement renice for your platform!"
 #endif
@@ -290,6 +292,19 @@
   } while (r > 0);
 }
 
+#ifdef PLATFORM_FREEBSD
+int64_t
+tvh_usleep(int64_t us)
+{
+  return usleep(us);
+}
+
+int64_t
+tvh_usleep_abs(int64_t us)
+{
+  return usleep(us - getfastmonoclock());
+}
+#else
 int64_t
 tvh_usleep(int64_t us)
 {
@@ -323,6 +338,7 @@
     return val;
   return r ? -r : 0;
 }
+#endif
 
 /*
  * qsort
