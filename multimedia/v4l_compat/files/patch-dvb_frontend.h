--- dvb/frontend.h.orig	2018-07-12 16:04:02 UTC
+++ dvb/frontend.h
@@ -908,7 +908,13 @@ struct dtv_properties {
 #define FE_DISHNETWORK_SEND_LEGACY_CMD _IO('o', 80) /* unsigned int */
 
 #define FE_SET_PROPERTY		   _IOW('o', 82, struct dtv_properties)
-#define FE_GET_PROPERTY		   _IOR('o', 83, struct dtv_properties)
+/* 
+ * This is broken on linux as well but they workaround it in the driver.
+ * Since this is impossible to do on FreeBSD fix the header instead.
+ * Detailed and discussion :
+ * http://lists.freebsd.org/pipermail/freebsd-multimedia/2010-April/010958.html
+ */
+#define FE_GET_PROPERTY		   _IOW('o', 83, struct dtv_properties)
 
 #if defined(__DVB_CORE__) || !defined(__KERNEL__)
 
