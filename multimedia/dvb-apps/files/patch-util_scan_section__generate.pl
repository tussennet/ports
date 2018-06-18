--- util/scan/section_generate.pl.orig	2018-06-18 19:55:13 UTC
+++ util/scan/section_generate.pl
@@ -4,7 +4,7 @@ use strict;
 
 die "no section perl file given" unless @ARGV;
 
-my $h = require($ARGV[0]);
+my $h = require("./".$ARGV[0]);
 
 our $basename;
 our $debug = $ARGV[1];
