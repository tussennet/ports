--- include/ca.h
+++ include/ca.h
@@ -85,6 +85,5 @@
 #define CA_GET_MSG        _IOR('o', 132, ca_msg_t)
 #define CA_SEND_MSG       _IOW('o', 133, ca_msg_t)
 #define CA_SET_DESCR      _IOW('o', 134, ca_descr_t)
-#define CA_SET_PID        _IOW('o', 135, ca_pid_t)
 
 #endif
