--- README.md.orig	2020-10-02 12:15:24 UTC
+++ README.md
@@ -1,4 +1,4 @@
-# runc
+# FreeBSD runc
 
 [![Build Status](https://travis-ci.org/opencontainers/runc.svg?branch=master)](https://travis-ci.org/opencontainers/runc)
 [![Go Report Card](https://goreportcard.com/badge/github.com/opencontainers/runc)](https://goreportcard.com/report/github.com/opencontainers/runc)
@@ -8,6 +8,14 @@
 
 `runc` is a CLI tool for spawning and running containers according to the OCI specification.
 
+This is an experiment for implementing a FreeBSD runtime for `runc`.
+Some of runc commands have been implemented:
+create, delete, exec, kill, list, ps, run, spec, start, state
+
+The commands: init, checkpoint, restore, update, events, and pause/resume are not supported yet.
+
+FreeBSD Jail is not a child process, so 'init' does not make any sense on FreeBSD platform.
+
 ## Releases
 
 `runc` depends on and tracks the [runtime-spec](https://github.com/opencontainers/runtime-spec) repository.
@@ -29,15 +37,18 @@ The reporting process and disclosure communications ar
 `runc` currently supports the Linux platform with various architecture support.
 It must be built with Go version 1.6 or higher in order for some features to function properly.
 
-In order to enable seccomp support you will need to install `libseccomp` on your platform.
-> e.g. `libseccomp-devel` for CentOS, or `libseccomp-dev` for Ubuntu
+       (2) prepare config.json:
+       ./runc spec
+       sed -i 's;"sh";"/bin/sh";' config.json
 
-Otherwise, if you do not want to build `runc` with seccomp support you can add `BUILDTAGS=""` when running make.
+       (3) launch a shell:
+       ./runc run test
 
+
 ```bash
 # create a 'github.com/opencontainers' in your GOPATH/src
 cd github.com/opencontainers
-git clone https://github.com/opencontainers/runc
+git clone https://github.com/clovertrail/runc.git
 cd runc
 
 make
@@ -111,7 +122,7 @@ new dependencies.
 
 ## Using runc
 
-### Creating an OCI Bundle
+### Creating an FreeBSD OCI Bundle
 
 In order to use runc you must have your container in the format of an OCI bundle.
 If you have Docker installed you can use its `export` method to acquire a root filesystem from an existing Docker container.
@@ -124,10 +135,14 @@ cd /mycontainer
 # create the rootfs directory
 mkdir rootfs
 
-# export busybox via Docker into the rootfs directory
-docker export $(docker create busybox) | tar -C rootfs -xvf -
-```
+# Download FreeBSD 11.0 release image files (base.txz and lib32.txz) from official site:
+fetch ftp://ftp.freebsd.org/pub/FreeBSD/releases/amd64/amd64/11.0-RELEASE/base.txz
+fetch ftp://ftp.freebsd.org/pub/FreeBSD/releases/amd64/amd64/11.0-RELEASE/lib32.txz
 
+tar xvf -C base.txz rootfs
+tar xvf -C lib32.txz rootfs
+
+```
 After a root filesystem is populated you just generate a spec in the format of a `config.json` file inside your bundle.
 `runc` provides a `spec` command to generate a base template spec that you are then able to edit.
 To find features and documentation for fields in the spec please refer to the [specs](https://github.com/opencontainers/runtime-spec) repository.
@@ -152,19 +167,19 @@ If you used the unmodified `runc spec` template this s
 
 The second way to start a container is using the specs lifecycle operations.
 This gives you more power over how the container is created and managed while it is running.
-This will also launch the container in the background so you will have to edit the `config.json` to remove the `terminal` setting for the simple examples here.
-Your process field in the `config.json` should look like this below with `"terminal": false` and `"args": ["sleep", "5"]`.
+This will also launch the container and pop up a shell for you.
+Your process field in the `config.json` should look like this below with `"args": ["/bin/sh"]`.
 
 
 ```json
         "process": {
-                "terminal": false,
+                "terminal": true,
                 "user": {
                         "uid": 0,
                         "gid": 0
                 },
                 "args": [
-                        "sleep", "5"
+                        "/bin/sh"
                 ],
                 "env": [
                         "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
@@ -221,7 +236,7 @@ runc create mycontainerid
 runc list
 
 # start the process inside the container
-runc start mycontainerid
+runc run mycontainerid
 
 # after 5 seconds view that the container has exited and is now in the stopped state
 runc list
