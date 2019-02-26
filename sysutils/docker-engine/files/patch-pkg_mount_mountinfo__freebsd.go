--- pkg/mount/mountinfo_freebsd.go.orig	2019-02-06 23:39:49 UTC
+++ pkg/mount/mountinfo_freebsd.go
@@ -37,7 +37,7 @@ func parseMountTable(filter FilterFunc) ([]*Info, erro
 
 		if filter != nil {
 			// filter out entries we're not interested in
-			skip, stop = filter(p)
+			skip, stop = filter(&mountinfo)
 			if skip {
 				continue
 			}
