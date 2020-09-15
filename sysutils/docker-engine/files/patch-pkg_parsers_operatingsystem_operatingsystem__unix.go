--- pkg/parsers/operatingsystem/operatingsystem_unix.go.orig	2020-09-04 14:54:52 UTC
+++ pkg/parsers/operatingsystem/operatingsystem_unix.go
@@ -1,4 +1,4 @@
-// +build freebsd darwin
+// +build darwin
 
 package operatingsystem // import "github.com/docker/docker/pkg/parsers/operatingsystem"
 
@@ -20,6 +20,5 @@ func GetOperatingSystem() (string, error) {
 // IsContainerized returns true if we are running inside a container.
 // No-op on FreeBSD and Darwin, always returns false.
 func IsContainerized() (bool, error) {
-	// TODO: Implement jail detection for freeBSD
 	return false, errors.New("Cannot detect if we are in container")
 }
