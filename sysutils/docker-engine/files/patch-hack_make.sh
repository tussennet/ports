--- hack/make.sh.orig	2019-10-07 21:12:15 UTC
+++ hack/make.sh
@@ -128,7 +128,6 @@ BUILDFLAGS=( ${BUILDFLAGS} "${ORIG_BUILDFLAGS[@]}" )
 
 LDFLAGS_STATIC_DOCKER="
 	$LDFLAGS_STATIC
-	-extldflags \"$EXTLDFLAGS_STATIC\"
 "
 
 if [ "$(uname -s)" = 'FreeBSD' ]; then
@@ -138,7 +137,7 @@ if [ "$(uname -s)" = 'FreeBSD' ]; then
 
 	# "-extld clang" is a workaround for
 	# https://code.google.com/p/go/issues/detail?id=6845
-	LDFLAGS="$LDFLAGS -extld clang"
+	LDFLAGS="$LDFLAGS -extld clang -extldflags -Wl,-z,notext"
 fi
 
 bundle() {
