--- src/wrappers.c.orig	2017-01-20 19:41:51.000000000 +0100
+++ src/wrappers.c	2017-02-08 12:18:01.540169000 +0100
@@ -191,6 +191,8 @@
   pid_t tid;
   tid = gettid();
   ret = setpriority(PRIO_PROCESS, tid, value);
+#elif PLATFORM_FREEBSD
+  ret = setpriority(PRIO_PROCESS, 0, value);
 #else
 #warning "Implement renice for your platform!"
 #endif
@@ -298,9 +300,16 @@
   int r;
   if (us <= 0)
     return 0;
+
   ts.tv_sec = us / 1000000LL;
   ts.tv_nsec = (us % 1000000LL) * 1000LL;
+
+#ifdef clock_nanosleep
   r = clock_nanosleep(CLOCK_MONOTONIC, 0, &ts, &ts);
+#else
+  r = nanosleep(&ts, NULL);
+#endif
+
   val = (ts.tv_sec * 1000000LL) + ((ts.tv_nsec + 500LL) / 1000LL);
   if (ERRNO_AGAIN(r))
     return val;
@@ -315,9 +324,18 @@
   int r;
   if (us <= 0)
     return 0;
+
+#ifdef clock_nanosleep
   ts.tv_sec = us / 1000000LL;
   ts.tv_nsec = (us % 1000000LL) * 1000LL;
   r = clock_nanosleep(CLOCK_MONOTONIC, TIMER_ABSTIME, &ts, &ts);
+#else
+  us -= getmonoclock();
+
+  ts.tv_sec = us / 1000000LL;
+  ts.tv_nsec = (us % 1000000LL) * 1000LL;
+  r = nanosleep(&ts, NULL);
+#endif
   val = (ts.tv_sec * 1000000LL) + ((ts.tv_nsec + 500LL) / 1000LL);
   if (ERRNO_AGAIN(r))
     return val;
