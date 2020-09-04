--- vendor/github.com/godbus/dbus/transport_freebsd.go.orig	2020-09-04 09:13:43 UTC
+++ vendor/github.com/godbus/dbus/transport_freebsd.go
@@ -0,0 +1,6 @@
+package dbus
+
+func (t *unixTransport) SendNullByte() error {
+	_, err := t.Write([]byte{0})
+	return err
+}
