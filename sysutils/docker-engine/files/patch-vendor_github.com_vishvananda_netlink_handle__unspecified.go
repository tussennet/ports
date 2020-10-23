--- vendor/github.com/vishvananda/netlink/handle_unspecified.go.orig	2020-10-23 18:37:22 UTC
+++ vendor/github.com/vishvananda/netlink/handle_unspecified.go
@@ -185,18 +185,6 @@ func (h *Handle) ClassList(link Link, parent uint32) (
 	return nil, ErrNotImplemented
 }
 
-func (h *Handle) FilterDel(filter Filter) error {
-	return ErrNotImplemented
-}
-
-func (h *Handle) FilterAdd(filter Filter) error {
-	return ErrNotImplemented
-}
-
-func (h *Handle) FilterList(link Link, parent uint32) ([]Filter, error) {
-	return nil, ErrNotImplemented
-}
-
 func (h *Handle) NeighAdd(neigh *Neigh) error {
 	return ErrNotImplemented
 }
