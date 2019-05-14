# $FreeBSD$

PORTNAME=       wazuh
PORTVERSION=    3.9.0
DISTVERSIONPREFIX=      v
CATEGORIES=     security
PKGNAMESUFFIX=  -agent

MAINTAINER=     m.muenz@gmail.com
COMMENT=        Security tool to monitor and check logs and intrusions

LICENSE=        GPLv2
LICENSE_FILE=   ${WRKSRC}/LICENSE

USES=           gmake perl5 readline

USE_GITHUB=     yes
GH_ACCOUNT=     wazuh
GH_TAGNAME=     v3.9.0

USE_OPENSSL=    yes
USE_GMAKE=	yes

USERS=          ossec ossecm ossecr
GROUPS=         ossec

.include <bsd.port.pre.mk>
do-build:
		@cd ${WRKSRC}/src && ${GMAKE} deps && ${GMAKE}
do-install:
		@cd ${WRKSRC} && ${SH} install.sh
.include <bsd.port.mk>
