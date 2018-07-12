--- ../linux/dvb/ca.h.orig	2018-06-20 14:17:16.975230000 +0200
+++ ../linux/dvb/ca.h	2018-06-20 14:22:58.803607000 +0200
@@ -134,9 +134,13 @@
 
 #define CA_RESET          _IO('o', 128)
 #define CA_GET_CAP        _IOR('o', 129, struct ca_caps)
-#define CA_GET_SLOT_INFO  _IOR('o', 130, struct ca_slot_info)
+
+/* At least CA_GET_SLOT_INFO and CA_GET_MSG need to be _IOWR not _IOR.
+ * This is wrong on Linux too but there the driver doesn't care.
+ */
+#define CA_GET_SLOT_INFO  _IOWR('o', 130, struct ca_slot_info)
 #define CA_GET_DESCR_INFO _IOR('o', 131, struct ca_descr_info)
-#define CA_GET_MSG        _IOR('o', 132, struct ca_msg)
+#define CA_GET_MSG        _IOWR('o', 132, struct ca_msg)
 #define CA_SEND_MSG       _IOW('o', 133, struct ca_msg)
 #define CA_SET_DESCR      _IOW('o', 134, struct ca_descr)
 
