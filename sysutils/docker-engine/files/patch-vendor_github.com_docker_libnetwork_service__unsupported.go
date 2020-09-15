--- vendor/github.com/docker/libnetwork/service_unsupported.go.orig	2020-09-04 14:54:57 UTC
+++ vendor/github.com/docker/libnetwork/service_unsupported.go
@@ -1,4 +1,4 @@
-// +build !linux,!windows
+// +build !linux,!windows,!freebsd
 
 package libnetwork
 
@@ -18,7 +18,7 @@ func (c *controller) rmServiceBinding(name, sid, nid, 
 	return fmt.Errorf("not supported")
 }
 
-func (sb *sandbox) populateLoadBalancers(ep *endpoint) {
+func (sb *sandbox) populateLoadbalancers(ep *endpoint) {
 }
 
 func arrangeIngressFilterRule() {
