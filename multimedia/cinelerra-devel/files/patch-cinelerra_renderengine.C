--- cinelerra/renderengine.C.orig	2011-11-13 22:03:22.000000000 +0200
+++ cinelerra/renderengine.C	2011-11-13 22:04:15.000000000 +0200
@@ -419,6 +419,7 @@
 		Thread::start();
 		start_lock->lock("RenderEngine::start_command 2");
 		start_lock->unlock();
+		interrupt_lock->unlock();
 	}
 	return 0;
 }
