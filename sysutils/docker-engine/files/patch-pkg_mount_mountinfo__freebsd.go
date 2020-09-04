--- pkg/mount/mountinfo_freebsd.go.orig	2019-10-07 21:12:15 UTC
+++ pkg/mount/mountinfo_freebsd.go
@@ -37,7 +37,7 @@ func parseMountTable(filter FilterFunc) ([]*Info, erro
 
 		if filter != nil {
 			// filter out entries we're not interested in
-			skip, stop = filter(p)
+			skip, stop = filter(&mountinfo)
 			if skip {
 				continue
 			}
