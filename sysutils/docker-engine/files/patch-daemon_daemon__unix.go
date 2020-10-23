--- daemon/daemon_unix.go.orig	2020-10-23 18:37:16 UTC
+++ daemon/daemon_unix.go
@@ -29,7 +29,8 @@ import (
 	"github.com/docker/docker/pkg/containerfs"
 	"github.com/docker/docker/pkg/idtools"
 	"github.com/docker/docker/pkg/ioutils"
-	"github.com/docker/docker/pkg/mount"
+
+	//"github.com/docker/docker/pkg/mount"
 	"github.com/docker/docker/pkg/parsers"
 	"github.com/docker/docker/pkg/parsers/kernel"
 	"github.com/docker/docker/pkg/sysinfo"
@@ -37,18 +38,18 @@ import (
 	volumemounts "github.com/docker/docker/volume/mounts"
 	"github.com/docker/libnetwork"
 	nwconfig "github.com/docker/libnetwork/config"
-	"github.com/docker/libnetwork/drivers/bridge"
+	"github.com/docker/libnetwork/drivers/freebsd/bridge"
 	"github.com/docker/libnetwork/netlabel"
 	"github.com/docker/libnetwork/netutils"
 	"github.com/docker/libnetwork/options"
 	lntypes "github.com/docker/libnetwork/types"
-	"github.com/opencontainers/runc/libcontainer/cgroups"
+
+	// "github.com/opencontainers/runc/libcontainer/cgroups"
 	rsystem "github.com/opencontainers/runc/libcontainer/system"
 	"github.com/opencontainers/runtime-spec/specs-go"
 	"github.com/opencontainers/selinux/go-selinux/label"
 	"github.com/pkg/errors"
 	"github.com/sirupsen/logrus"
-	"github.com/vishvananda/netlink"
 	"golang.org/x/sys/unix"
 )
 
@@ -874,11 +875,11 @@ func (daemon *Daemon) initNetworkController(config *co
 	}
 
 	// Initialize default network on "host"
-	if n, _ := controller.NetworkByName("host"); n == nil {
-		if _, err := controller.NewNetwork("host", "host", "", libnetwork.NetworkOptionPersist(true)); err != nil {
-			return nil, fmt.Errorf("Error creating default \"host\" network: %v", err)
-		}
-	}
+	// if n, _ := controller.NetworkByName("host"); n == nil {
+	// 	if _, err := controller.NewNetwork("host", "host", "", libnetwork.NetworkOptionPersist(true)); err != nil {
+	// 		return nil, fmt.Errorf("Error creating default \"host\" network: %v", err)
+	// 	}
+	// }
 
 	// Clear stale bridge network
 	if n, err := controller.NetworkByName("bridge"); err == nil {
@@ -1043,16 +1044,13 @@ func initBridgeDriver(controller libnetwork.NetworkCon
 	if err != nil {
 		return fmt.Errorf("Error creating default \"bridge\" network: %v", err)
 	}
+
 	return nil
 }
 
 // Remove default bridge interface if present (--bridge=none use case)
-func removeDefaultBridgeInterface() {
-	if lnk, err := netlink.LinkByName(bridge.DefaultBridgeName); err == nil {
-		if err := netlink.LinkDel(lnk); err != nil {
-			logrus.Warnf("Failed to remove bridge interface (%s): %v", bridge.DefaultBridgeName, err)
-		}
-	}
+func removeDefaultBridgeInterface() error {
+	return fmt.Errorf("Bridge network driver not supported on FreeBSD (yet)")
 }
 
 func setupInitLayer(idMapping *idtools.IdentityMapping) func(containerfs.ContainerFS) error {
@@ -1260,45 +1258,45 @@ func setupDaemonRoot(config *config.Config, rootDir st
 }
 
 func setupDaemonRootPropagation(cfg *config.Config) error {
-	rootParentMount, options, err := getSourceMount(cfg.Root)
-	if err != nil {
-		return errors.Wrap(err, "error getting daemon root's parent mount")
-	}
+	// rootParentMount, options, err := getSourceMount(cfg.Root)
+	// if err != nil {
+	// 	return errors.Wrap(err, "error getting daemon root's parent mount")
+	// }
 
-	var cleanupOldFile bool
-	cleanupFile := getUnmountOnShutdownPath(cfg)
-	defer func() {
-		if !cleanupOldFile {
-			return
-		}
-		if err := os.Remove(cleanupFile); err != nil && !os.IsNotExist(err) {
-			logrus.WithError(err).WithField("file", cleanupFile).Warn("could not clean up old root propagation unmount file")
-		}
-	}()
+	// var cleanupOldFile bool
+	// cleanupFile := getUnmountOnShutdownPath(cfg)
+	// defer func() {
+	// 	if !cleanupOldFile {
+	// 		return
+	// 	}
+	// 	if err := os.Remove(cleanupFile); err != nil && !os.IsNotExist(err) {
+	// 		logrus.WithError(err).WithField("file", cleanupFile).Warn("could not clean up old root propagation unmount file")
+	// 	}
+	// }()
 
-	if hasMountinfoOption(options, sharedPropagationOption, slavePropagationOption) {
-		cleanupOldFile = true
-		return nil
-	}
+	// if hasMountinfoOption(options, sharedPropagationOption, slavePropagationOption) {
+	// 	cleanupOldFile = true
+	// 	return nil
+	// }
 
-	if err := mount.MakeShared(cfg.Root); err != nil {
-		return errors.Wrap(err, "could not setup daemon root propagation to shared")
-	}
+	// if err := mount.MakeShared(cfg.Root); err != nil {
+	// 	return errors.Wrap(err, "could not setup daemon root propagation to shared")
+	// }
 
-	// check the case where this may have already been a mount to itself.
-	// If so then the daemon only performed a remount and should not try to unmount this later.
-	if rootParentMount == cfg.Root {
-		cleanupOldFile = true
-		return nil
-	}
+	// // check the case where this may have already been a mount to itself.
+	// // If so then the daemon only performed a remount and should not try to unmount this later.
+	// if rootParentMount == cfg.Root {
+	// 	cleanupOldFile = true
+	// 	return nil
+	// }
 
-	if err := os.MkdirAll(filepath.Dir(cleanupFile), 0700); err != nil {
-		return errors.Wrap(err, "error creating dir to store mount cleanup file")
-	}
+	// if err := os.MkdirAll(filepath.Dir(cleanupFile), 0700); err != nil {
+	// 	return errors.Wrap(err, "error creating dir to store mount cleanup file")
+	// }
 
-	if err := ioutil.WriteFile(cleanupFile, nil, 0600); err != nil {
-		return errors.Wrap(err, "error writing file to signal mount cleanup on shutdown")
-	}
+	// if err := ioutil.WriteFile(cleanupFile, nil, 0600); err != nil {
+	// 	return errors.Wrap(err, "error writing file to signal mount cleanup on shutdown")
+	// }
 	return nil
 }
 
@@ -1387,7 +1385,7 @@ func (daemon *Daemon) stats(c *container.Container) (*
 	if !c.IsRunning() {
 		return nil, errNotRunning(c.ID)
 	}
-	cs, err := daemon.containerd.Stats(context.Background(), c.ID)
+	_, err := daemon.containerd.Stats(context.Background(), c.ID)
 	if err != nil {
 		if strings.Contains(err.Error(), "container not found") {
 			return nil, containerNotFound(c.ID)
@@ -1395,97 +1393,97 @@ func (daemon *Daemon) stats(c *container.Container) (*
 		return nil, err
 	}
 	s := &types.StatsJSON{}
-	s.Read = cs.Read
-	stats := cs.Metrics
-	if stats.Blkio != nil {
-		s.BlkioStats = types.BlkioStats{
-			IoServiceBytesRecursive: copyBlkioEntry(stats.Blkio.IoServiceBytesRecursive),
-			IoServicedRecursive:     copyBlkioEntry(stats.Blkio.IoServicedRecursive),
-			IoQueuedRecursive:       copyBlkioEntry(stats.Blkio.IoQueuedRecursive),
-			IoServiceTimeRecursive:  copyBlkioEntry(stats.Blkio.IoServiceTimeRecursive),
-			IoWaitTimeRecursive:     copyBlkioEntry(stats.Blkio.IoWaitTimeRecursive),
-			IoMergedRecursive:       copyBlkioEntry(stats.Blkio.IoMergedRecursive),
-			IoTimeRecursive:         copyBlkioEntry(stats.Blkio.IoTimeRecursive),
-			SectorsRecursive:        copyBlkioEntry(stats.Blkio.SectorsRecursive),
-		}
-	}
-	if stats.CPU != nil {
-		s.CPUStats = types.CPUStats{
-			CPUUsage: types.CPUUsage{
-				TotalUsage:        stats.CPU.Usage.Total,
-				PercpuUsage:       stats.CPU.Usage.PerCPU,
-				UsageInKernelmode: stats.CPU.Usage.Kernel,
-				UsageInUsermode:   stats.CPU.Usage.User,
-			},
-			ThrottlingData: types.ThrottlingData{
-				Periods:          stats.CPU.Throttling.Periods,
-				ThrottledPeriods: stats.CPU.Throttling.ThrottledPeriods,
-				ThrottledTime:    stats.CPU.Throttling.ThrottledTime,
-			},
-		}
-	}
+	// s.Read = cs.Read
+	// stats := cs.Metrics
+	// if stats.Blkio != nil {
+	// 	s.BlkioStats = types.BlkioStats{
+	// 		IoServiceBytesRecursive: copyBlkioEntry(stats.Blkio.IoServiceBytesRecursive),
+	// 		IoServicedRecursive:     copyBlkioEntry(stats.Blkio.IoServicedRecursive),
+	// 		IoQueuedRecursive:       copyBlkioEntry(stats.Blkio.IoQueuedRecursive),
+	// 		IoServiceTimeRecursive:  copyBlkioEntry(stats.Blkio.IoServiceTimeRecursive),
+	// 		IoWaitTimeRecursive:     copyBlkioEntry(stats.Blkio.IoWaitTimeRecursive),
+	// 		IoMergedRecursive:       copyBlkioEntry(stats.Blkio.IoMergedRecursive),
+	// 		IoTimeRecursive:         copyBlkioEntry(stats.Blkio.IoTimeRecursive),
+	// 		SectorsRecursive:        copyBlkioEntry(stats.Blkio.SectorsRecursive),
+	// 	}
+	// }
+	// if stats.CPU != nil {
+	// 	s.CPUStats = types.CPUStats{
+	// 		CPUUsage: types.CPUUsage{
+	// 			TotalUsage:        stats.CPU.Usage.Total,
+	// 			PercpuUsage:       stats.CPU.Usage.PerCPU,
+	// 			UsageInKernelmode: stats.CPU.Usage.Kernel,
+	// 			UsageInUsermode:   stats.CPU.Usage.User,
+	// 		},
+	// 		ThrottlingData: types.ThrottlingData{
+	// 			Periods:          stats.CPU.Throttling.Periods,
+	// 			ThrottledPeriods: stats.CPU.Throttling.ThrottledPeriods,
+	// 			ThrottledTime:    stats.CPU.Throttling.ThrottledTime,
+	// 		},
+	// 	}
+	// }
 
-	if stats.Memory != nil {
-		raw := make(map[string]uint64)
-		raw["cache"] = stats.Memory.Cache
-		raw["rss"] = stats.Memory.RSS
-		raw["rss_huge"] = stats.Memory.RSSHuge
-		raw["mapped_file"] = stats.Memory.MappedFile
-		raw["dirty"] = stats.Memory.Dirty
-		raw["writeback"] = stats.Memory.Writeback
-		raw["pgpgin"] = stats.Memory.PgPgIn
-		raw["pgpgout"] = stats.Memory.PgPgOut
-		raw["pgfault"] = stats.Memory.PgFault
-		raw["pgmajfault"] = stats.Memory.PgMajFault
-		raw["inactive_anon"] = stats.Memory.InactiveAnon
-		raw["active_anon"] = stats.Memory.ActiveAnon
-		raw["inactive_file"] = stats.Memory.InactiveFile
-		raw["active_file"] = stats.Memory.ActiveFile
-		raw["unevictable"] = stats.Memory.Unevictable
-		raw["hierarchical_memory_limit"] = stats.Memory.HierarchicalMemoryLimit
-		raw["hierarchical_memsw_limit"] = stats.Memory.HierarchicalSwapLimit
-		raw["total_cache"] = stats.Memory.TotalCache
-		raw["total_rss"] = stats.Memory.TotalRSS
-		raw["total_rss_huge"] = stats.Memory.TotalRSSHuge
-		raw["total_mapped_file"] = stats.Memory.TotalMappedFile
-		raw["total_dirty"] = stats.Memory.TotalDirty
-		raw["total_writeback"] = stats.Memory.TotalWriteback
-		raw["total_pgpgin"] = stats.Memory.TotalPgPgIn
-		raw["total_pgpgout"] = stats.Memory.TotalPgPgOut
-		raw["total_pgfault"] = stats.Memory.TotalPgFault
-		raw["total_pgmajfault"] = stats.Memory.TotalPgMajFault
-		raw["total_inactive_anon"] = stats.Memory.TotalInactiveAnon
-		raw["total_active_anon"] = stats.Memory.TotalActiveAnon
-		raw["total_inactive_file"] = stats.Memory.TotalInactiveFile
-		raw["total_active_file"] = stats.Memory.TotalActiveFile
-		raw["total_unevictable"] = stats.Memory.TotalUnevictable
+	// if stats.Memory != nil {
+	// 	raw := make(map[string]uint64)
+	// 	raw["cache"] = stats.Memory.Cache
+	// 	raw["rss"] = stats.Memory.RSS
+	// 	raw["rss_huge"] = stats.Memory.RSSHuge
+	// 	raw["mapped_file"] = stats.Memory.MappedFile
+	// 	raw["dirty"] = stats.Memory.Dirty
+	// 	raw["writeback"] = stats.Memory.Writeback
+	// 	raw["pgpgin"] = stats.Memory.PgPgIn
+	// 	raw["pgpgout"] = stats.Memory.PgPgOut
+	// 	raw["pgfault"] = stats.Memory.PgFault
+	// 	raw["pgmajfault"] = stats.Memory.PgMajFault
+	// 	raw["inactive_anon"] = stats.Memory.InactiveAnon
+	// 	raw["active_anon"] = stats.Memory.ActiveAnon
+	// 	raw["inactive_file"] = stats.Memory.InactiveFile
+	// 	raw["active_file"] = stats.Memory.ActiveFile
+	// 	raw["unevictable"] = stats.Memory.Unevictable
+	// 	raw["hierarchical_memory_limit"] = stats.Memory.HierarchicalMemoryLimit
+	// 	raw["hierarchical_memsw_limit"] = stats.Memory.HierarchicalSwapLimit
+	// 	raw["total_cache"] = stats.Memory.TotalCache
+	// 	raw["total_rss"] = stats.Memory.TotalRSS
+	// 	raw["total_rss_huge"] = stats.Memory.TotalRSSHuge
+	// 	raw["total_mapped_file"] = stats.Memory.TotalMappedFile
+	// 	raw["total_dirty"] = stats.Memory.TotalDirty
+	// 	raw["total_writeback"] = stats.Memory.TotalWriteback
+	// 	raw["total_pgpgin"] = stats.Memory.TotalPgPgIn
+	// 	raw["total_pgpgout"] = stats.Memory.TotalPgPgOut
+	// 	raw["total_pgfault"] = stats.Memory.TotalPgFault
+	// 	raw["total_pgmajfault"] = stats.Memory.TotalPgMajFault
+	// 	raw["total_inactive_anon"] = stats.Memory.TotalInactiveAnon
+	// 	raw["total_active_anon"] = stats.Memory.TotalActiveAnon
+	// 	raw["total_inactive_file"] = stats.Memory.TotalInactiveFile
+	// 	raw["total_active_file"] = stats.Memory.TotalActiveFile
+	// 	raw["total_unevictable"] = stats.Memory.TotalUnevictable
 
-		if stats.Memory.Usage != nil {
-			s.MemoryStats = types.MemoryStats{
-				Stats:    raw,
-				Usage:    stats.Memory.Usage.Usage,
-				MaxUsage: stats.Memory.Usage.Max,
-				Limit:    stats.Memory.Usage.Limit,
-				Failcnt:  stats.Memory.Usage.Failcnt,
-			}
-		} else {
-			s.MemoryStats = types.MemoryStats{
-				Stats: raw,
-			}
-		}
+	// 	if stats.Memory.Usage != nil {
+	// 		s.MemoryStats = types.MemoryStats{
+	// 			Stats:    raw,
+	// 			Usage:    stats.Memory.Usage.Usage,
+	// 			MaxUsage: stats.Memory.Usage.Max,
+	// 			Limit:    stats.Memory.Usage.Limit,
+	// 			Failcnt:  stats.Memory.Usage.Failcnt,
+	// 		}
+	// 	} else {
+	// 		s.MemoryStats = types.MemoryStats{
+	// 			Stats: raw,
+	// 		}
+	// 	}
 
-		// if the container does not set memory limit, use the machineMemory
-		if s.MemoryStats.Limit > daemon.machineMemory && daemon.machineMemory > 0 {
-			s.MemoryStats.Limit = daemon.machineMemory
-		}
-	}
+	// 	// if the container does not set memory limit, use the machineMemory
+	// 	if s.MemoryStats.Limit > daemon.machineMemory && daemon.machineMemory > 0 {
+	// 		s.MemoryStats.Limit = daemon.machineMemory
+	// 	}
+	// }
 
-	if stats.Pids != nil {
-		s.PidsStats = types.PidsStats{
-			Current: stats.Pids.Current,
-			Limit:   stats.Pids.Limit,
-		}
-	}
+	// if stats.Pids != nil {
+	// 	s.PidsStats = types.PidsStats{
+	// 		Current: stats.Pids.Current,
+	// 		Limit:   stats.Pids.Limit,
+	// 	}
+	// }
 
 	return s, nil
 }
@@ -1538,24 +1536,7 @@ func setMayDetachMounts() error {
 }
 
 func setupOOMScoreAdj(score int) error {
-	f, err := os.OpenFile("/proc/self/oom_score_adj", os.O_WRONLY, 0)
-	if err != nil {
-		return err
-	}
-	defer f.Close()
-	stringScore := strconv.Itoa(score)
-	_, err = f.WriteString(stringScore)
-	if os.IsPermission(err) {
-		// Setting oom_score_adj does not work in an
-		// unprivileged container. Ignore the error, but log
-		// it if we appear not to be in that situation.
-		if !rsystem.RunningInUserNS() {
-			logrus.Debugf("Permission denied writing %q to /proc/self/oom_score_adj", stringScore)
-		}
-		return nil
-	}
-
-	return err
+	return nil
 }
 
 func (daemon *Daemon) initCgroupsPath(path string) error {
@@ -1571,7 +1552,10 @@ func (daemon *Daemon) initCgroupsPath(path string) err
 	// for the period and runtime as this limits what the children can be set to.
 	daemon.initCgroupsPath(filepath.Dir(path))
 
-	mnt, root, err := cgroups.FindCgroupMountpointAndRoot("", "cpu")
+	mnt := ""
+	root := ""
+	var err error = nil
+	//mnt, root, err := cgroups.FindCgroupMountpointAndRoot("", "cpu")
 	if err != nil {
 		return err
 	}
