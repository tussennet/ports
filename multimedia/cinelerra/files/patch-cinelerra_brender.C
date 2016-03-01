--- cinelerra/brender.C.orig	2015-08-13 14:04:04 UTC
+++ cinelerra/brender.C
@@ -144,14 +144,22 @@ void BRender::run()
 
 
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
