--- guicast/filesystem.C.orig
+++ guicast/filesystem.C
@@ -393,7 +393,7 @@ int FileSystem::test_filter(FileItem *file)
 int FileSystem::update(char *new_dir)
 {
 	DIR *dirstream;
-	struct dirent64 *new_filename;
+	struct dirent *new_filename;
 	struct stat ostat;
 	struct tm *mod_time;
 	int i, j, k, include_this;
@@ -408,7 +408,7 @@ int FileSystem::update(char *new_dir)
 	dirstream = opendir(current_dir);
 	if(!dirstream) return 1;          // failed to open directory
 
-	while(new_filename = readdir64(dirstream))
+	while(new_filename = readdir(dirstream))
 	{
 		include_this = 1;
 
