- Add missing include noticed on FreeBSD 11.

src/spawn.c:306:7: warning: implicit declaration of function 'kill' is invalid in C99 [-Wimplicit-function-declaration]
      kill(-(s->pid), SIGKILL);
      ^
src/spawn.c:306:23: error: use of undeclared identifier 'SIGKILL'
      kill(-(s->pid), SIGKILL);
                      ^
src/spawn.c:655:3: warning: implicit declaration of function 'pthread_kill' is invalid in C99
      [-Wimplicit-function-declaration]
  pthread_kill(spawn_pipe_tid, SIGTERM);
  ^
src/spawn.c:655:32: error: use of undeclared identifier 'SIGTERM'
  pthread_kill(spawn_pipe_tid, SIGTERM);
--- src/spawn.c.orig	2015-09-25 13:57:59 UTC
+++ src/spawn.c
@@ -28,6 +28,7 @@
 #include <syslog.h>
 #include <fcntl.h>
 #include <dirent.h>
+#include <signal.h>
 
 #include "tvheadend.h"
 #include "tvhpoll.h"
