--- daemon/daemon_unix.go.orig	2019-06-18 21:30:11 UTC
+++ daemon/daemon_unix.go
@@ -7,7 +7,6 @@
 	"context"
 	"fmt"
 	"io/ioutil"
-	"net"
 	"os"
 	"path/filepath"
 	"runtime"
@@ -29,7 +28,7 @@
 	"github.com/docker/docker/pkg/containerfs"
 	"github.com/docker/docker/pkg/idtools"
 	"github.com/docker/docker/pkg/ioutils"
-	"github.com/docker/docker/pkg/mount"
+	//"github.com/docker/docker/pkg/mount"
 	"github.com/docker/docker/pkg/parsers"
 	"github.com/docker/docker/pkg/parsers/kernel"
 	"github.com/docker/docker/pkg/sysinfo"
@@ -37,18 +36,14 @@
 	volumemounts "github.com/docker/docker/volume/mounts"
 	"github.com/docker/libnetwork"
 	nwconfig "github.com/docker/libnetwork/config"
-	"github.com/docker/libnetwork/drivers/bridge"
 	"github.com/docker/libnetwork/netlabel"
-	"github.com/docker/libnetwork/netutils"
 	"github.com/docker/libnetwork/options"
-	lntypes "github.com/docker/libnetwork/types"
-	"github.com/opencontainers/runc/libcontainer/cgroups"
+	// "github.com/opencontainers/runc/libcontainer/cgroups"
 	rsystem "github.com/opencontainers/runc/libcontainer/system"
 	"github.com/opencontainers/runtime-spec/specs-go"
 	"github.com/opencontainers/selinux/go-selinux/label"
 	"github.com/pkg/errors"
 	"github.com/sirupsen/logrus"
-	"github.com/vishvananda/netlink"
 	"golang.org/x/sys/unix"
 )
 
@@ -914,143 +909,12 @@
 }
 
 func initBridgeDriver(controller libnetwork.NetworkController, config *config.Config) error {
-	bridgeName := bridge.DefaultBridgeName
-	if config.BridgeConfig.Iface != "" {
-		bridgeName = config.BridgeConfig.Iface
-	}
-	netOption := map[string]string{
-		bridge.BridgeName:         bridgeName,
-		bridge.DefaultBridge:      strconv.FormatBool(true),
-		netlabel.DriverMTU:        strconv.Itoa(config.Mtu),
-		bridge.EnableIPMasquerade: strconv.FormatBool(config.BridgeConfig.EnableIPMasq),
-		bridge.EnableICC:          strconv.FormatBool(config.BridgeConfig.InterContainerCommunication),
-	}
-
-	// --ip processing
-	if config.BridgeConfig.DefaultIP != nil {
-		netOption[bridge.DefaultBindingIP] = config.BridgeConfig.DefaultIP.String()
-	}
-
-	var (
-		ipamV4Conf *libnetwork.IpamConf
-		ipamV6Conf *libnetwork.IpamConf
-	)
-
-	ipamV4Conf = &libnetwork.IpamConf{AuxAddresses: make(map[string]string)}
-
-	nwList, nw6List, err := netutils.ElectInterfaceAddresses(bridgeName)
-	if err != nil {
-		return errors.Wrap(err, "list bridge addresses failed")
-	}
-
-	nw := nwList[0]
-	if len(nwList) > 1 && config.BridgeConfig.FixedCIDR != "" {
-		_, fCIDR, err := net.ParseCIDR(config.BridgeConfig.FixedCIDR)
-		if err != nil {
-			return errors.Wrap(err, "parse CIDR failed")
-		}
-		// Iterate through in case there are multiple addresses for the bridge
-		for _, entry := range nwList {
-			if fCIDR.Contains(entry.IP) {
-				nw = entry
-				break
-			}
-		}
-	}
-
-	ipamV4Conf.PreferredPool = lntypes.GetIPNetCanonical(nw).String()
-	hip, _ := lntypes.GetHostPartIP(nw.IP, nw.Mask)
-	if hip.IsGlobalUnicast() {
-		ipamV4Conf.Gateway = nw.IP.String()
-	}
-
-	if config.BridgeConfig.IP != "" {
-		ipamV4Conf.PreferredPool = config.BridgeConfig.IP
-		ip, _, err := net.ParseCIDR(config.BridgeConfig.IP)
-		if err != nil {
-			return err
-		}
-		ipamV4Conf.Gateway = ip.String()
-	} else if bridgeName == bridge.DefaultBridgeName && ipamV4Conf.PreferredPool != "" {
-		logrus.Infof("Default bridge (%s) is assigned with an IP address %s. Daemon option --bip can be used to set a preferred IP address", bridgeName, ipamV4Conf.PreferredPool)
-	}
-
-	if config.BridgeConfig.FixedCIDR != "" {
-		_, fCIDR, err := net.ParseCIDR(config.BridgeConfig.FixedCIDR)
-		if err != nil {
-			return err
-		}
-
-		ipamV4Conf.SubPool = fCIDR.String()
-	}
-
-	if config.BridgeConfig.DefaultGatewayIPv4 != nil {
-		ipamV4Conf.AuxAddresses["DefaultGatewayIPv4"] = config.BridgeConfig.DefaultGatewayIPv4.String()
-	}
-
-	var deferIPv6Alloc bool
-	if config.BridgeConfig.FixedCIDRv6 != "" {
-		_, fCIDRv6, err := net.ParseCIDR(config.BridgeConfig.FixedCIDRv6)
-		if err != nil {
-			return err
-		}
-
-		// In case user has specified the daemon flag --fixed-cidr-v6 and the passed network has
-		// at least 48 host bits, we need to guarantee the current behavior where the containers'
-		// IPv6 addresses will be constructed based on the containers' interface MAC address.
-		// We do so by telling libnetwork to defer the IPv6 address allocation for the endpoints
-		// on this network until after the driver has created the endpoint and returned the
-		// constructed address. Libnetwork will then reserve this address with the ipam driver.
-		ones, _ := fCIDRv6.Mask.Size()
-		deferIPv6Alloc = ones <= 80
-
-		if ipamV6Conf == nil {
-			ipamV6Conf = &libnetwork.IpamConf{AuxAddresses: make(map[string]string)}
-		}
-		ipamV6Conf.PreferredPool = fCIDRv6.String()
-
-		// In case the --fixed-cidr-v6 is specified and the current docker0 bridge IPv6
-		// address belongs to the same network, we need to inform libnetwork about it, so
-		// that it can be reserved with IPAM and it will not be given away to somebody else
-		for _, nw6 := range nw6List {
-			if fCIDRv6.Contains(nw6.IP) {
-				ipamV6Conf.Gateway = nw6.IP.String()
-				break
-			}
-		}
-	}
-
-	if config.BridgeConfig.DefaultGatewayIPv6 != nil {
-		if ipamV6Conf == nil {
-			ipamV6Conf = &libnetwork.IpamConf{AuxAddresses: make(map[string]string)}
-		}
-		ipamV6Conf.AuxAddresses["DefaultGatewayIPv6"] = config.BridgeConfig.DefaultGatewayIPv6.String()
-	}
-
-	v4Conf := []*libnetwork.IpamConf{ipamV4Conf}
-	v6Conf := []*libnetwork.IpamConf{}
-	if ipamV6Conf != nil {
-		v6Conf = append(v6Conf, ipamV6Conf)
-	}
-	// Initialize default network on "bridge" with the same name
-	_, err = controller.NewNetwork("bridge", "bridge", "",
-		libnetwork.NetworkOptionEnableIPv6(config.BridgeConfig.EnableIPv6),
-		libnetwork.NetworkOptionDriverOpts(netOption),
-		libnetwork.NetworkOptionIpam("default", "", v4Conf, v6Conf, nil),
-		libnetwork.NetworkOptionDeferIPv6Alloc(deferIPv6Alloc))
-	if err != nil {
-		return fmt.Errorf("Error creating default \"bridge\" network: %v", err)
-	}
-	return nil
+	return fmt.Errorf("Bridge network driver not supported on FreeBSD (yet)")
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
@@ -1237,45 +1101,45 @@
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
 
@@ -1364,7 +1228,7 @@
 	if !c.IsRunning() {
 		return nil, errNotRunning(c.ID)
 	}
-	cs, err := daemon.containerd.Stats(context.Background(), c.ID)
+	_, err := daemon.containerd.Stats(context.Background(), c.ID)
 	if err != nil {
 		if strings.Contains(err.Error(), "container not found") {
 			return nil, containerNotFound(c.ID)
@@ -1372,97 +1236,97 @@
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
@@ -1548,7 +1412,10 @@
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
