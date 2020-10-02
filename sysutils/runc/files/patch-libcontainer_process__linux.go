--- libcontainer/process_linux.go.orig	2020-10-02 12:15:24 UTC
+++ libcontainer/process_linux.go
@@ -5,7 +5,6 @@ package libcontainer
 import (
 	"encoding/json"
 	"errors"
-	"fmt"
 	"io"
 	"os"
 	"os/exec"
