--- hack/make.sh.orig	2019-10-07 23:12:15.000000000 +0200
+++ hack/make.sh	2020-08-28 16:51:57.589492000 +0200
@@ -128,7 +128,6 @@
 
 LDFLAGS_STATIC_DOCKER="
 	$LDFLAGS_STATIC
-	-extldflags \"$EXTLDFLAGS_STATIC\"
 "
 
 if [ "$(uname -s)" = 'FreeBSD' ]; then
@@ -138,7 +137,7 @@
 
 	# "-extld clang" is a workaround for
 	# https://code.google.com/p/go/issues/detail?id=6845
-	LDFLAGS="$LDFLAGS -extld clang"
+	LDFLAGS="$LDFLAGS -extld clang -extldflags -Wl,-z,notext"
 fi
 
 bundle() {
