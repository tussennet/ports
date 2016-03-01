--- cinelerra/brender.C.orig	2010-12-20 12:41:00.000000000 +0200
+++ cinelerra/brender.C	2010-12-20 13:07:00.000000000 +0200
@@ -144,14 +144,22 @@
 
 
 // Construct executable command with the designated filesystem port
+#if defined(__FreeBSD__) || defined(__FreeBSD_kernel__)
+	fd = fopen("/compat/linux/proc/self/cmdline", "r");
+#else
 	fd = fopen("/proc/self/cmdline", "r");
+#endif
 	if(fd)
 	{
 		fread(string, 1, BCTEXTLEN, fd);
 		fclose(fd);
 	}
 	else
+#if defined(__FreeBSD__) || defined(__FreeBSD_kernel__)
+		perror(_("BRender::fork_background: can't open /compat/linux/proc/self/cmdline.\n"));
+#else
 		perror(_("BRender::fork_background: can't open /proc/self/cmdline.\n"));
+#endif
 
 	arguments[0] = new char[strlen(string) + 1];
 	strcpy(arguments[0], string);
