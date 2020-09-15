--- vendor/github.com/docker/libnetwork/service_freebsd.go.orig	2020-09-04 09:13:43 UTC
+++ vendor/github.com/docker/libnetwork/service_freebsd.go
@@ -0,0 +1,302 @@
+package libnetwork
+
+import (
+	"fmt"
+	"io"
+	"io/ioutil"
+	"net"
+	"os"
+	"os/exec"
+	"strings"
+	"sync"
+
+	"github.com/docker/docker/pkg/reexec"
+	"github.com/gogo/protobuf/proto"
+	"github.com/ishidawataru/sctp"
+	"github.com/sirupsen/logrus"
+)
+
+func init() {
+	reexec.Register("fwmarker", fwMarker)
+	reexec.Register("redirector", redirector)
+}
+
+// Populate all loadbalancers on the network that the passed endpoint
+// belongs to, into this sandbox.
+func (sb *sandbox) populateLoadBalancers(ep *endpoint) {
+	// This is an interface less endpoint. Nothing to do.
+	if ep.Iface() == nil {
+		return
+	}
+
+	n := ep.getNetwork()
+	eIP := ep.Iface().Address()
+
+	if n.ingress {
+		if err := addRedirectRules(sb.Key(), eIP, ep.ingressPorts); err != nil {
+			logrus.Errorf("Failed to add redirect rules for ep %s (%.7s): %v", ep.Name(), ep.ID(), err)
+		}
+	}
+}
+
+func (n *network) findLBEndpointSandbox() (*endpoint, *sandbox, error) {
+	// TODO: get endpoint from store?  See EndpointInfo()
+	var ep *endpoint
+	// Find this node's LB sandbox endpoint:  there should be exactly one
+	for _, e := range n.Endpoints() {
+		epi := e.Info()
+		if epi != nil && epi.LoadBalancer() {
+			ep = e.(*endpoint)
+			break
+		}
+	}
+	if ep == nil {
+		return nil, nil, fmt.Errorf("Unable to find load balancing endpoint for network %s", n.ID())
+	}
+	// Get the load balancer sandbox itself as well
+	sb, ok := ep.getSandbox()
+	if !ok {
+		return nil, nil, fmt.Errorf("Unable to get sandbox for %s(%s) in for %s", ep.Name(), ep.ID(), n.ID())
+	}
+	ep = sb.getEndpoint(ep.ID())
+	if ep == nil {
+		return nil, nil, fmt.Errorf("Load balancing endpoint %s(%s) removed from %s", ep.Name(), ep.ID(), n.ID())
+	}
+	return ep, sb, nil
+}
+
+// Searches the OS sandbox for the name of the endpoint interface
+// within the sandbox.   This is required for adding/removing IP
+// aliases to the interface.
+func findIfaceDstName(sb *sandbox, ep *endpoint) string {
+	srcName := ep.Iface().SrcName()
+	for _, i := range sb.osSbox.Info().Interfaces() {
+		if i.SrcName() == srcName {
+			return i.DstName()
+		}
+	}
+	return ""
+}
+
+// Add loadbalancer backend to the loadbalncer sandbox for the network.
+// If needed add the service as well.
+func (n *network) addLBBackend(ip net.IP, lb *loadBalancer) {
+	//return fmt.Errorf("not supported")
+}
+
+// Remove loadbalancer backend the load balancing endpoint for this
+// network. If 'rmService' is true, then remove the service entry as well.
+// If 'fullRemove' is true then completely remove the entry, otherwise
+// just deweight it for now.
+func (n *network) rmLBBackend(ip net.IP, lb *loadBalancer, rmService bool, fullRemove bool) {
+	//return fmt.Errorf("not supported")
+}
+
+const ingressChain = "DOCKER-INGRESS"
+
+var (
+	ingressOnce     sync.Once
+	ingressMu       sync.Mutex // lock for operations on ingress
+	ingressProxyTbl = make(map[string]io.Closer)
+	portConfigMu    sync.Mutex
+	portConfigTbl   = make(map[PortConfig]int)
+)
+
+func filterPortConfigs(ingressPorts []*PortConfig, isDelete bool) []*PortConfig {
+	portConfigMu.Lock()
+	iPorts := make([]*PortConfig, 0, len(ingressPorts))
+	for _, pc := range ingressPorts {
+		if isDelete {
+			if cnt, ok := portConfigTbl[*pc]; ok {
+				// This is the last reference to this
+				// port config. Delete the port config
+				// and add it to filtered list to be
+				// plumbed.
+				if cnt == 1 {
+					delete(portConfigTbl, *pc)
+					iPorts = append(iPorts, pc)
+					continue
+				}
+
+				portConfigTbl[*pc] = cnt - 1
+			}
+
+			continue
+		}
+
+		if cnt, ok := portConfigTbl[*pc]; ok {
+			portConfigTbl[*pc] = cnt + 1
+			continue
+		}
+
+		// We are adding it for the first time. Add it to the
+		// filter list to be plumbed.
+		portConfigTbl[*pc] = 1
+		iPorts = append(iPorts, pc)
+	}
+	portConfigMu.Unlock()
+
+	return iPorts
+}
+
+func programIngress(gwIP net.IP, ingressPorts []*PortConfig, isDelete bool) error {
+	return fmt.Errorf("not supported")
+}
+
+// In the filter table FORWARD chain the first rule should be to jump to
+// DOCKER-USER so the user is able to filter packet first.
+// The second rule should be jump to INGRESS-CHAIN.
+// This chain has the rules to allow access to the published ports for swarm tasks
+// from local bridge networks and docker_gwbridge (ie:taks on other swarm networks)
+func arrangeIngressFilterRule() {
+	//return fmt.Errorf("not supported")
+}
+
+func findOIFName(ip net.IP) (string, error) {
+	return "", fmt.Errorf("not supported")
+}
+
+func plumbProxy(iPort *PortConfig, isDelete bool) error {
+	var (
+		err error
+		l   io.Closer
+	)
+
+	portSpec := fmt.Sprintf("%d/%s", iPort.PublishedPort, strings.ToLower(PortConfig_Protocol_name[int32(iPort.Protocol)]))
+	if isDelete {
+		if listener, ok := ingressProxyTbl[portSpec]; ok {
+			if listener != nil {
+				listener.Close()
+			}
+		}
+
+		return nil
+	}
+
+	switch iPort.Protocol {
+	case ProtocolTCP:
+		l, err = net.ListenTCP("tcp", &net.TCPAddr{Port: int(iPort.PublishedPort)})
+	case ProtocolUDP:
+		l, err = net.ListenUDP("udp", &net.UDPAddr{Port: int(iPort.PublishedPort)})
+	case ProtocolSCTP:
+		l, err = sctp.ListenSCTP("sctp", &sctp.SCTPAddr{Port: int(iPort.PublishedPort)})
+	default:
+		err = fmt.Errorf("unknown protocol %v", iPort.Protocol)
+	}
+
+	if err != nil {
+		return err
+	}
+
+	ingressProxyTbl[portSpec] = l
+
+	return nil
+}
+
+func writePortsToFile(ports []*PortConfig) (string, error) {
+	f, err := ioutil.TempFile("", "port_configs")
+	if err != nil {
+		return "", err
+	}
+	defer f.Close()
+
+	buf, _ := proto.Marshal(&EndpointRecord{
+		IngressPorts: ports,
+	})
+
+	n, err := f.Write(buf)
+	if err != nil {
+		return "", err
+	}
+
+	if n < len(buf) {
+		return "", io.ErrShortWrite
+	}
+
+	return f.Name(), nil
+}
+
+func readPortsFromFile(fileName string) ([]*PortConfig, error) {
+	buf, err := ioutil.ReadFile(fileName)
+	if err != nil {
+		return nil, err
+	}
+
+	var epRec EndpointRecord
+	err = proto.Unmarshal(buf, &epRec)
+	if err != nil {
+		return nil, err
+	}
+
+	return epRec.IngressPorts, nil
+}
+
+// Invoke fwmarker reexec routine to mark vip destined packets with
+// the passed firewall mark.
+func invokeFWMarker(path string, vip net.IP, fwMark uint32, ingressPorts []*PortConfig, eIP *net.IPNet, isDelete bool, lbMode string) error {
+	var ingressPortsFile string
+
+	if len(ingressPorts) != 0 {
+		var err error
+		ingressPortsFile, err = writePortsToFile(ingressPorts)
+		if err != nil {
+			return err
+		}
+
+		defer os.Remove(ingressPortsFile)
+	}
+
+	addDelOpt := "-A"
+	if isDelete {
+		addDelOpt = "-D"
+	}
+
+	cmd := &exec.Cmd{
+		Path:   reexec.Self(),
+		Args:   append([]string{"fwmarker"}, path, vip.String(), fmt.Sprintf("%d", fwMark), addDelOpt, ingressPortsFile, eIP.String(), lbMode),
+		Stdout: os.Stdout,
+		Stderr: os.Stderr,
+	}
+
+	if err := cmd.Run(); err != nil {
+		return fmt.Errorf("reexec failed: %v", err)
+	}
+
+	return nil
+}
+
+// Firewall marker reexec function.
+func fwMarker() {
+	//return fmt.Errorf("not supported")
+}
+
+func addRedirectRules(path string, eIP *net.IPNet, ingressPorts []*PortConfig) error {
+	var ingressPortsFile string
+
+	if len(ingressPorts) != 0 {
+		var err error
+		ingressPortsFile, err = writePortsToFile(ingressPorts)
+		if err != nil {
+			return err
+		}
+		defer os.Remove(ingressPortsFile)
+	}
+
+	cmd := &exec.Cmd{
+		Path:   reexec.Self(),
+		Args:   append([]string{"redirector"}, path, eIP.String(), ingressPortsFile),
+		Stdout: os.Stdout,
+		Stderr: os.Stderr,
+	}
+
+	if err := cmd.Run(); err != nil {
+		return fmt.Errorf("reexec failed: %v", err)
+	}
+
+	return nil
+}
+
+// Redirector reexec function.
+func redirector() {
+	//return fmt.Errorf("not supported")
+}
