# TODO for my ports

## KODI

### devel/kodi-platform:
- waiting for p8-platform update and kodi addon fix

### multimedia/kodi-addon-pvr-hts:
### multimedia/kodi-addon-pvr-iptvsimple:
- waiting for kodi addon fix
- IPTV runtime testing with https://wiki.ubuntuusers.de/Internet-TV/Stationen/


## DVB-EN50221

### multimedia/dtv-scan-tables:
- waiting for timeout of PR 229151 (20180703)

### multimedia/dvb-apps:
- SONAME missing for libraries
- See https://svnweb.freebsd.org/ports/head/science/cdf/files/patch-Makefile?r1=423146&r2=423145&pathrev=423146

### multimedia/tvheadend:
- Commit DVBEN50221 option after dvb-apps is committed
- Wait for feedback from hps@


## Other

### multimedia/minisatip:
- new port for https://github.com/catalinii/minisatip

### www/helma:
- runtime testing needed

