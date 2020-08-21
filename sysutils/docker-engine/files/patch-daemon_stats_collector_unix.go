--- daemon/stats/collector_unix.go.orig	2019-10-07 23:12:15.000000000 +0200
+++ daemon/stats/collector_unix.go	2020-08-21 18:25:40.785012000 +0200
@@ -7,15 +7,11 @@
 	"os"
 	"strconv"
 	"strings"
-
-	"github.com/opencontainers/runc/libcontainer/system"
-	"golang.org/x/sys/unix"
 )
 
 // platformNewStatsCollector performs platform specific initialisation of the
 // Collector structure.
 func platformNewStatsCollector(s *Collector) {
-	s.clockTicksPerSecond = uint64(system.GetClockTicks())
 }
 
 const nanoSecondsPerSecond = 1e9
@@ -66,10 +62,5 @@
 }
 
 func (s *Collector) getNumberOnlineCPUs() (uint32, error) {
-	var cpuset unix.CPUSet
-	err := unix.SchedGetaffinity(0, &cpuset)
-	if err != nil {
-		return 0, err
-	}
-	return uint32(cpuset.Count()), nil
+	return 0, nil
 }
