--- cinelerra/ieee1394-ioctl.h.orig
+++ cinelerra/ieee1394-ioctl.h
@@ -25,8 +25,8 @@
 #ifndef __IEEE1394_IOCTL_H
 #define __IEEE1394_IOCTL_H
 
-#include <linux/ioctl.h>
-#include <linux/types.h>
+#include <sys/ioctl.h>
+#include <sys/types.h>
 
 
 /* AMDTP Gets 6 */
