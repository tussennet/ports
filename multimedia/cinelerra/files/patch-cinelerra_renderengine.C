--- cinelerra/renderengine.C.orig	2015-08-13 14:04:04 UTC
+++ cinelerra/renderengine.C
@@ -419,6 +419,7 @@ int RenderEngine::start_command()
 		Thread::start();
 		start_lock->lock("RenderEngine::start_command 2");
 		start_lock->unlock();
+		interrupt_lock->unlock();
 	}
 	return 0;
 }
