--- plugins/titler/title.C.orig
+++ plugins/titler/title.C
@@ -48,8 +48,7 @@
 #include <stdint.h>
 #include <stdio.h>
 #include <string.h>
-#include <endian.h>
-#include <byteswap.h>
+#include <sys/endian.h>
 #include <iconv.h>
 #include <sys/stat.h>
 
@@ -1517,7 +1516,8 @@ void TitleMain::draw_glyphs()
 
 			size_t inbytes,outbytes;
 			char inbuf;
-			char *inp = (char*)&inbuf, *outp = (char *)&char_code;
+			const char *inp = (const char*)&inbuf;
+			char *outp = (char *)&char_code;
 			
 			inbuf = (char)c;
 			inbytes = 1;
@@ -1525,7 +1525,7 @@ void TitleMain::draw_glyphs()
 	
 			iconv (cd, &inp, &inbytes, &outp, &outbytes);
 #if     __BYTE_ORDER == __LITTLE_ENDIAN
-				char_code = bswap_32(char_code);
+				char_code = bswap32(char_code);
 #endif                          /* Big endian.  */
 
 		} else {
