--- archutils/epoll_freebsd.go.orig	2020-10-02 13:03:30 UTC
+++ archutils/epoll_freebsd.go
@@ -0,0 +1,116 @@
+// +build freebsd
+
+// -----------------------------------------------------------------------------------------
+//
+//
+// (@kris-nova)
+//
+// Probably most of my god awful hacking is done in this file, most of this is in place to
+// hack around the Go standard library since we are using a Linux package in FreeBSD.. we
+// need to have certain constants defined that usually are only found based on $GOOS,.
+//
+// For more information or just to yell at me shoot me a line at kris@nivenly.com
+//
+//
+// -----------------------------------------------------------------------------------------
+
+package archutils
+
+// #cgo CFLAGS:  -I/usr/local/include/libepoll-shim
+// #cgo LDFLAGS: -L/usr/local/lib -Ild -lepoll-shim -lrt
+// #include	<sys/epoll.h>
+/*
+int EpollCreate1(int flag) {
+	return epoll_create1(flag);
+}
+
+int EpollCtl(int efd, int op,int sfd, int events, int fd) {
+	struct epoll_event event;
+	event.events = events;
+	event.data.fd = fd;
+	return epoll_ctl(efd, op, sfd, &event);
+}
+
+struct event_t {
+	uint32_t events;
+	epoll_data_t data;
+	int fd;
+};
+
+struct epoll_event events[128];
+
+int run_epoll_wait(int fd, struct event_t *event) {
+	int n, i;
+	n = epoll_wait(fd, events, 128, 0);
+	for (i = 0; i < n; i++) {
+		event[i].events = events[i].events;
+		event[i].fd = events[i].data.fd;
+	}
+	return n;
+}
+*/
+import "C"
+
+import (
+	"fmt"
+	//"unsafe"
+)
+
+// EpollCreate1 calls a C implementation
+func EpollCreate1(flag int) (int, error) {
+	fd := int(C.EpollCreate1(C.int(flag)))
+	if fd < 0 {
+		return fd, fmt.Errorf("failed to create epoll, errno is %d", fd)
+	}
+	return fd, nil
+}
+
+type FreeBSDEpollEventInterface interface {
+}
+
+type FreeBSDEpollEvent struct {
+	Events uint32
+	Fd     int32
+	Pad    int32
+}
+
+const (
+	// -----------------------------------------------------
+	// Hacking in control constants for FreeBSD Epoll port
+	// Note: these are not defined in the Go standard library
+	// so we define them here manually. Once the constants make it
+	// to the epoll.h file, the declaration in Go shouldn't matter.
+	//
+	// (@kris-nova)
+	//
+	FREEBSD_EPOLL_CTL_ADD = 1
+	FREEBSD_EPOLL_CTL_DEL = 2
+	FREEBSD_EPOLL_CLOEXEC = 0x00100000
+	FREEBSD_EPOLLHUP      = 0x010
+	FREEBSD_EPOLLIN       = 0x001
+	FREEBSD_SYS_EPOLL_CTL = 233
+)
+
+// EpollCtl calls a C implementation
+func EpollCtl(epfd int, op int, fd int, eventInterface FreeBSDEpollEventInterface) error {
+	event := eventInterface.(*FreeBSDEpollEvent)
+	errno := C.EpollCtl(C.int(epfd), C.int(FREEBSD_EPOLL_CTL_ADD), C.int(fd), C.int(event.Events), C.int(event.Fd))
+	if errno < 0 {
+		return fmt.Errorf("Failed to ctl epoll")
+	}
+	return nil
+}
+
+// EpollWait calls a C implementation
+func EpollWait(epfd int, events []FreeBSDEpollEvent, msec int) (int, error) {
+	var c_events [128]C.struct_event_t
+	//n := int(C.run_epoll_wait(C.int(epfd), (*C.struct_event_t)(unsafe.Pointer(&c_events))))
+	//if n < 0 {
+	//	return int(n), fmt.Errorf("Failed to wait epoll")
+	//}
+	for i := 0; i < epfd; i++ {
+		events[i].Fd = int32(c_events[i].fd)
+		events[i].Events = uint32(c_events[i].events)
+	}
+	return int(epfd), nil
+}
