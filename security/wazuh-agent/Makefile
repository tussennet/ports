# $FreeBSD$

PORTNAME=	wazuh
DISTVERSION=	3.9.0
DISTVERSIONPREFIX=	v
CATEGORIES=	security
PKGNAMESUFFIX=	-agent

MAINTAINER=	m.muenz@gmail.com
COMMENT=	Security tool to monitor and check logs and intrusions

LICENSE=	GPLv2
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		gmake perl5 readline

USE_GITHUB=	yes

USE_OPENSSL=	yes
USE_GMAKE=	yes

USERS=		ossec ossecm ossecr
GROUPS=		ossec

do-build:
		@cd ${WRKSRC}/src && ${GMAKE} deps && ${GMAKE}

do-install:
		@cd ${WRKSRC} && ${SH} install.sh

.include <bsd.port.mk>
