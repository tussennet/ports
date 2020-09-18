--- hack/make.sh.orig	2020-09-18 09:00:53 UTC
+++ hack/make.sh
@@ -120,7 +120,6 @@ BUILDFLAGS=( ${BUILDFLAGS} "${ORIG_BUILDFLAGS[@]}" )
 
 LDFLAGS_STATIC_DOCKER="
 	$LDFLAGS_STATIC
-	-extldflags \"$EXTLDFLAGS_STATIC\"
 "
 
 if [ "$(uname -s)" = 'FreeBSD' ]; then
@@ -130,7 +129,7 @@ if [ "$(uname -s)" = 'FreeBSD' ]; then
 
 	# "-extld clang" is a workaround for
 	# https://code.google.com/p/go/issues/detail?id=6845
-	LDFLAGS="$LDFLAGS -extld clang"
+	LDFLAGS="$LDFLAGS -extld clang -extldflags -Wl,-z,notext"
 fi
 
 bundle() {
