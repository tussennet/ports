--- vendor/github.com/godbus/dbus/transport_freebsd.go.orig	2019-02-26 21:19:13 UTC
+++ vendor/github.com/godbus/dbus/transport_freebsd.go
@@ -0,0 +1,6 @@
+package dbus
+
+func (t *unixTransport) SendNullByte() error {
+	_, err := t.Write([]byte{0})
+	return err
+}
